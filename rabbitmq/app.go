package main

/**
* 当前代码ack确认发送到exchange，并不确认是否发送到消息队列
**/
import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var (
	log_ch       *amqp.Channel
	log_confirms chan amqp.Confirmation
)

func init() {
	source, err := amqp.Dial("amqp://127.0.0.1:5672")

	if err != nil {
		log.Fatalf("connecting open souce err, %s", err)
	}

	log_ch, err = source.Channel()
	if err != nil {
		log.Fatalf("channel.open source: %s", err)
	}

	if err := log_ch.ExchangeDeclare("log", "topic", true, false, false, false, nil); err != nil {
		log.Fatalf("exchange.declare destination: %s", err)
	}

	if _, err := log_ch.QueueDeclare("log_queue", true, false, false, false, nil); err != nil {
		log.Fatalf("queue.declare source: %s", err)
	}

	if err := log_ch.QueueBind("log_queue", "info", "log", false, nil); err != nil {
		log.Fatalf("queue.bind source: %s", err)
	}

	// 消息队列进入确认模式
	log_confirms = log_ch.NotifyPublish(make(chan amqp.Confirmation, 1))

	if err := log_ch.Confirm(false); err != nil {
		log.Fatalf("confirm.select destination: %s", err)
	}

}

func main() {
	// 发送订阅
	err := log_ch.Publish("log", "info", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		Body:         []byte("go go amqp"),
	})

	if err != nil {
		// Since publish is asynchronous this can happen if the network connection
		// is reset or if the server has run out of resources.
		log.Fatalf("basic.publish: %v", err)
	}
	// ack机制
	if log_confirmed := <-log_confirms; log_confirmed.Ack {
		//    msg.Ack(false)
		fmt.Println("rabbitmq成功接受消息成功")
	} else {
		//   msg.Nack(false, false)
		log.Fatal("rabbitmq接受消息失败")
	}
}
