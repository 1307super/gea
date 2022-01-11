package sms

import (
	"gea/app/utils/rabbitmq"
)

// Message 队列消息
type Message struct {
	Id  string      `json:"id" from:"id"`
	Msg interface{} `json:"msg" from:"msg"`
}

type Producer struct {
}

func NewProducer() *Producer {
	return &Producer{}
}

// Publish 生产
func (mq *Producer) Publish(topic, exchangeName, body string) {
	p := rabbitmq.NewProducer(exchangeName, "direct", exchangeName, topic)
	defer p.Close()
	p.Publish(body)
}
