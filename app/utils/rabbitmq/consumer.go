package rabbitmq

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/streadway/amqp"
	"math/rand"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

// Consumer holds all infromation
// about the RabbitMQ connection
// This setup does limit a consumer
// to one exchange. This should not be
// an issue. Having to connect to multiple
// exchanges means something else is
// structured improperly.
type Consumer struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	done         chan error
	consumerTag  string // Name that consumer identifies itself to the server with
	uri          string // uri of the rabbitmq server
	exchange     string // exchange that we will bind to
	exchangeType string // topic, direct, etc...
	queueName    string // queueName that we will bind to

	lastRecoverTime int64
	//track service current status
	currentStatus atomic.Value
	receiver  Receiver
}

// 定义接收者接口
type Receiver interface {
	Consumer([]byte) error
	FailAction(error, []byte) error
}

// NewConsumer returns a Consumer struct that has been initialized properly
// essentially don't touch conn, channel, or done and you can create Consumer manually
func newConsumer(consumerTag, uri, exchange, exchangeType, queueName string) *Consumer {
	name, err := os.Hostname()
	if err != nil {
		name = "_sim"
	}
	consumer := &Consumer{
		consumerTag:     fmt.Sprintf("%s%s", consumerTag, name),
		uri:             uri,
		exchange:        exchange,
		exchangeType:    exchangeType,
		queueName:       queueName,
		done:            make(chan error),
		lastRecoverTime: time.Now().Unix(),
	}
	consumer.currentStatus.Store(true)
	return consumer
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func RunConsumer(consumerTag, exchange, exchangeType, queueName string, receiver Receiver) {
	rabbitUri := g.Cfg().GetString("rabbitmq.url")
	consumer := newConsumer(consumerTag, rabbitUri, exchange, exchangeType, queueName)
	if err := consumer.Connect(); err != nil {
		fmt.Printf("%+v\n%s", err, fmt.Sprintf("[%s]connect error", consumerTag))
		os.Exit(0)
	}
	consumer.receiver = receiver
	deliveries, err := consumer.AnnounceQueue()
	if err != nil {
		fmt.Printf("%+v\n%s", err, fmt.Sprintf("[%s]Error when calling AnnounceQueue()", consumerTag))
		fmt.Println()
	}
	consumer.Handle(deliveries, maxParallelism())
}

// ReConnect is called in places where NotifyClose() channel is called
// wait 30 seconds before trying to reconnect. Any shorter amount of time
// will  likely destroy the error log while waiting for servers to come
// back online. This requires two parameters which is just to satisfy
// the AccounceQueue call and allows greater flexability
func (c *Consumer) ReConnect(retryTime int) (<-chan amqp.Delivery, error) {
	c.Close()
	time.Sleep(time.Duration(15+rand.Intn(60)+2*retryTime) * time.Second)
	fmt.Printf("Try ReConnect with times: %d", retryTime)
	fmt.Println()

	if err := c.Connect(); err != nil {
		return nil, err
	}

	deliveries, err := c.AnnounceQueue()
	if err != nil {
		return deliveries, errors.New("Couldn't connect")
	}
	return deliveries, nil
}

// Connect to RabbitMQ server
func (c *Consumer) Connect() error {

	var err error
	//fmt.Printf("Try Connect with times: %s", c.uri)
	//fmt.Println()
	c.conn, err = amqp.Dial(c.uri)

	if err != nil {
		return fmt.Errorf("Dial: %+v", err)
	}

	go func() {
		// Waits here for the channel to be closed
		fmt.Printf("closing: %+v", <-c.conn.NotifyClose(make(chan *amqp.Error)))
		fmt.Println()
		// Let Handle know it's not time to reconnect
		c.done <- errors.New("Channel Closed")
	}()

	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %+v", err)
	}

	if err = c.channel.ExchangeDeclare(
		c.exchange,     // name of the exchange
		c.exchangeType, // type
		true,           // durable
		false,          // delete when complete
		false,          // internal
		false,          // noWait
		nil,            // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %+v", err)
	}

	return nil
}

// AnnounceQueue sets the queue that will be listened to for this
// connection...
func (c *Consumer) AnnounceQueue() (<-chan amqp.Delivery, error) {

	queue, err := c.channel.QueueDeclare(
		c.queueName, // name of the queue
		true,        // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // noWait
		nil,         // arguments
	)

	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	// Qos determines the amount of messages that the queue will pass to you before
	// it waits for you to ack them. This will slow down queue consumption but
	// give you more certainty that all messages are being processed. As load increases
	// I would reccomend upping the about of Threads and Processors the go process
	// uses before changing this although you will eventually need to reach some
	// balance between threads, procs, and Qos.
	err = c.channel.Qos(5, 0, false)
	if err != nil {
		return nil, fmt.Errorf("Error setting qos: %s", err)
	}

	if err = c.channel.QueueBind(
		queue.Name, // name of the queue
		c.exchange, // routingKey
		c.exchange, // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	deliveries, err := c.channel.Consume(
		queue.Name,    // name
		c.consumerTag, // consumerTag,
		false,         // noAck
		false,         // exclusive
		false,         // noLocal
		false,         // noWait
		nil,           // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}
	return deliveries, nil
}

func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
		c.channel = nil
	}
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
}

func (c *Consumer) Handle(
	deliveries <-chan amqp.Delivery,
	threads int) {
	var err error
	for {
		for i := 0; i < threads; i++ {
			go func() {
				for msg := range deliveries {
					retry_nums, ok := msg.Headers["retry_nums"].(int32)
					if !ok {
						retry_nums = int32(0)
					}
					body := msg.Body[:]
					err = c.receiver.Consumer(body)
					if err != nil {
						//消息处理失败 进入延时尝试机制
						if retry_nums < 3 {
							fmt.Println(string(msg.Body))
							fmt.Printf("消息处理失败 消息开始进入尝试  ttl延时队列 \n")
							c.retry_msg(msg.Body, retry_nums)
						} else {
							//消息失败 入库db
							fmt.Printf("消息处理3次后还是失败了 进入消费失败逻辑 \n")
							c.receiver.FailAction(err, msg.Body)
						}
					}
					msg.Ack(false)
				}
			}()
		}

		// Go into reconnect loop when
		// c.done is passed non nil values
		if <-c.done != nil {
			c.currentStatus.Store(false)
			retryTime := 1
			for {
				deliveries, err = c.ReConnect(retryTime)
				if err != nil {
					fmt.Printf("Reconnecting Error: %+v", err)
					fmt.Println()
					retryTime += 1
				} else {
					break
				}
			}
		}
		fmt.Print("Reconnected!!!")
	}
}

//消息处理失败之后 延时尝试
func (c *Consumer) retry_msg(msg []byte, retryNums int32) {
	fmt.Println(time.Now().Second())
	//原始队列名称 交换机名称
	newQName := c.queueName + "_retry_3"
	newExchangeName := c.exchange
	newRoutingKey := c.exchange + "_retry_3"

	mq := NewProducer(newExchangeName, c.exchangeType, newRoutingKey, newQName, true)
	mq.Retry(retryNums,c.exchange,c.exchange).Publish(string(msg))
	defer mq.Close()
}
