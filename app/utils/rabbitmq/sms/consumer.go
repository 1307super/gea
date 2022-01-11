package sms

import (
	"gea/app/utils/rabbitmq"
)

type Consumer struct {
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

// Consumer 实现消费者 消费消息失败 自动进入延时尝试  尝试3次之后入库db
/*
返回值 error 为nil  则表示该消息消费成功
否则消息会进入ttl延时队列  重复尝试消费3次
3次后消息如果还是失败 消息就执行失败  进入告警 FailAction
*/
func (t *Consumer) Consumer(dataByte []byte) error {
	// 处理消费数据
	return nil
}

// FailAction 消息已经消费3次 失败了 请进行处理
/*
如果消息 消费3次后 仍然失败  此处可以根据情况 对消息进行告警提醒 或者 补偿  入库db  钉钉告警等等
*/
func (t *Consumer) FailAction(err error, dataByte []byte) error {
	// 处理消费失败的逻辑
	return nil
}

func SmsRecv(topic string, exchangeName string) {
	/*
		runNums: 表示任务并发处理数量  一般建议 普通任务1-3    就可以了
	*/
	consume := NewConsumer()
	rabbitmq.RunConsumer("sms.recv", exchangeName, "direct", topic, consume)
}
