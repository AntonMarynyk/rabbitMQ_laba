package utils

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func consume(ch *amqp.Channel, queue_name string) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		queue_name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, fmt.Sprintf("Failed to register a consumer for %s", queue_name))

	return msgs
}

func init_exchange(ch *amqp.Channel, exchange_name string, exchange_type string) {
	err := ch.ExchangeDeclare(
		exchange_name, // name
		exchange_type, // type
		true,          // durable
		true,          // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, fmt.Sprintf("Failed to declare an exchange %s", exchange_name))
}

func init_queue(ch *amqp.Channel, queue_name string, args amqp.Table) amqp.Queue {
	q, err := ch.QueueDeclare(
		queue_name, // name
		false,      // durable
		true,       // delete when unused
		false,      // exclusive
		false,      // no-wait
		args,       // arguments
	)
	failOnError(err, fmt.Sprintf("Failed to declare a %s", queue_name))

	return q
}

func bind_queue(ch *amqp.Channel, queue_name string, routing_key string, exchange_name string) {
	err := ch.QueueBind(
		queue_name,    // queue name
		routing_key,   // routing key
		exchange_name, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
}

func main_queue(ch *amqp.Channel) amqp.Queue {
	// exchange for main queue
	init_exchange(ch, "exchange", "direct")

	// main queue
	q := init_queue(ch, "user_notifications", nil)

	// bind queue for exchange with routing key user1
	bind_queue(ch, q.Name, "user1", "exchange")

	return q
}

func secondary_queue(ch *amqp.Channel) amqp.Queue {
	// exchange for main queue
	init_exchange(ch, "exchange", "direct")

	// main queue
	q := init_queue(ch, "queue_for_testing_dlq", amqp.Table{
		"x-message-ttl":          int32(1000),
		"x-dead-letter-exchange": "user_dlx",
	})

	// bind queue for exchange with routing key user1
	bind_queue(ch, q.Name, "user2", "exchange")

	return q
}

func dead_letter_queue(ch *amqp.Channel) amqp.Queue {
	// dead letter exchange
	init_exchange(ch, "user_dlx", "fanout")

	// dead letter queue
	dlq := init_queue(ch, "user_create_dlx", nil)

	// bind dlq for dlx without routing_key
	bind_queue(ch, dlq.Name, "", "user_dlx")

	return dlq
}
