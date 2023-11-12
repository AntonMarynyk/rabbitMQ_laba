package utils

import (
	"context"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func validateProducerInput() {
	if len(os.Args) < 4 {
		log.Panicf("Usage: go run main.go producer [routing_key] [message]")
		return
	}
}

func parseProducerRoutingKey(args []string) string {
	return args[2]
}

func parseProducerBody(args []string) string {
	return args[3]
}

func declareContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return ctx, cancel
}

func publishMessage(ctx context.Context, ch *amqp.Channel) {
	routing_key := parseProducerRoutingKey(os.Args)
	message := parseProducerBody(os.Args)

	err := ch.PublishWithContext(ctx,
		"exchange",  // exchange
		routing_key, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", message)
}
