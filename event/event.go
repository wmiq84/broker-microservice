package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	// return exchange with these vars
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable?
		false,        // auto-deleted
		false,        // internal?
		false,        // no wait
		nil,          // arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	// return queue with these vars
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete unused
		true,  // exclusive
		false, // no-wait?
		nil,   // arguments
	)
}
