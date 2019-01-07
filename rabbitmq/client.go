package main

import (
	"log"

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

}

func main() {
	logs, err := log_ch.Consume("log_queue", "pager", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("basic.consume: %v", err)
	}

	for my_log := range logs {

		log.Printf("asd: %+v", my_log)
		log.Printf("asdcxc: %+v", string(my_log.Body))
		my_log.Ack(true)
	}
}
