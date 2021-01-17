package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	// 是否开启文件异步传输
	AsyncTransferEnable = true
	// 服务入口
	RabbitURL = "amqp://guest:guset@127.0.0.1:5762"
	//用于文件的transer交换机
	TransExchangeName = "uploadserver.oss"
	//oss 转移队列名
	TransOssQueryName = "uploadserver.trans.oss"
	//oss转移失败后移入另一个队列名
	TransOssErrQueueName = "uploadserver.trans.oss.err"
	// routing key
	TransOssRoutingKey = "oss"
)

// 消息格式
type TransferData struct {
	FileHash      string
	CurLocation   string
	DestLocation  string
	DestStoreType string
}

var conn *amqp.Connection
var channel *amqp.Channel

func NewChannel() bool {
	// 判断channel 是否已经创建过
	if channel != nil {
		return true
	}
	//获得一个rabbitmq的链接
	conn, err := amqp.Dial(RabbitURL)
	if err != nil {
		return false
	}
	channel, err = conn.Channel()
	if err != nil {
		return false
	}
	return true
}

//生产
func Publish(exchange, routingKey string, msg []byte) bool {

	//1。检查channel是否正常
	if !NewChannel() {
		return false
	}
	//2。执行发布消息
	err := channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		Headers:     nil,
		ContentType: "text/plain",
		Body:        msg,
	})
	if err != nil {
		return false
	}
	return true
}

// 消费
var done chan bool

func Consume(qName, CName string, callback func(msg []byte) bool) {
	//1。 通过channel 获得消息信道
	msgs, err := channel.Consume(qName, CName, true, false, false, false, nil)
	if err != nil {
		return
	}
	//2。循环获取队列消息
	done = make(chan bool)
	go func() {
		for msg := range msgs {
			processSuc := callback(msg.Body)
			if !processSuc {
				//TODO 写入另一个队列
			}
		}
	}()
	//3。调用callback方法来处理新的消息
	<-done
}
