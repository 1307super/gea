package rabbitmq

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/streadway/amqp"
	"os"
)

// Producer holds all infromation
// about the RabbitMQ connection
// This setup does limit a consumer
// to one exchange. This should not be
// an issue. Having to connect to multiple
// exchanges means something else is
// structured improperly.
type Producer struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	uri          string // uri of the rabbitmq server
	exchange     string // exchange that we will bind to
	exchangeType string // topic, direct, etc...
	queueName    string // exchange that we will bind to
	routingKey   string // topic, direct, etc...
	table        amqp.Table
	header       amqp.Table
}

func newProducer(uri, exchange, exchangeType, routingKey, queueName string) *Producer {
	producer := &Producer{
		uri:          uri,
		exchange:     exchange,
		exchangeType: exchangeType,
		queueName:    queueName,
		routingKey:   routingKey,
	}
	return producer
}

func NewProducer(exchange, exchangeType, routingKey, queueName string, retry ...bool) *Producer {
	rabbitUri := g.Cfg().GetString("rabbitmq.url")
	producer := newProducer(rabbitUri, exchange, exchangeType, routingKey, queueName)

	if err := producer.Connect(); err != nil {
		fmt.Printf("%+v\n%s", err, fmt.Sprintf("connect error"))
		os.Exit(0)
	}

	if len(retry) > 0 && retry[0]{
		return producer
	}
	err := producer.AnnounceQueue()
	if err == nil {
		return producer
	}
	fmt.Println()
	fmt.Printf("%+v\n%s", err, fmt.Sprintf("Error when calling AnnounceQueue()"))
	return nil
}

func (c *Producer) Publish(body string) error {
	if err := c.channel.Publish(
		c.exchange,
		c.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Headers:     c.header,
		},
	); err != nil {
		return fmt.Errorf(fmt.Sprintf("Publish err :%s \n", err))
	}
	return nil
}

func (c *Producer) Retry(retryNums int32, args ...string) *Producer {
	//原始路由key
	oldRoutingKey := args[0]
	//原始交换机名
	oldExchangeName := args[1]

	table := make(map[string]interface{}, 3)
	table["x-dead-letter-routing-key"] = oldRoutingKey
	if oldExchangeName != "" {
		table["x-dead-letter-exchange"] = oldExchangeName
	} else {
		table["x-dead-letter-exchange"] = ""
	}

	//table["x-message-ttl"] = int64(20000)
	table["x-message-ttl"] = int64(5000)
	c.table = table

	header := make(map[string]interface{}, 1)
	header["retry_nums"] = retryNums + int32(1)
	c.header = header
	c.AnnounceQueue()
	return c
}

// Connect to RabbitMQ server
func (c *Producer) Connect() error {
	var err error
	fmt.Printf("Try Connect with times: %s", c.uri)
	fmt.Println()
	c.conn, err = amqp.Dial(c.uri)

	if err != nil {
		return fmt.Errorf("Dial: %+v", err)
	}

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
func (c *Producer) AnnounceQueue() error {

	if err := c.channel.ExchangeDeclare(c.exchange, "direct", true, false, false, false, nil); err != nil {
		return fmt.Errorf("ExchangeDeclare Declare: %+v", err)
	}

	queue, err := c.channel.QueueDeclare(
		c.queueName, // name of the queue
		true,        // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // noWait
		c.table,     // arguments
	)

	if err != nil {
		return fmt.Errorf("Queue Declare: %s", err)
	}

	if err = c.channel.QueueBind(
		queue.Name,   // name of the queue
		c.routingKey, // routingKey
		c.exchange,   // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("Queue Bind: %+v", err)
	}
	return nil
}

func (c *Producer) Close() {
	if c.channel != nil {
		c.channel.Close()
		c.channel = nil
	}
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
}
