package utils

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func init_connection_channel() (*amqp.Connection, *amqp.Channel) {
	// create connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	// open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return conn, ch
}
