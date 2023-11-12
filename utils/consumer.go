package utils

import (
	"log"
)

func Consumer() {
	conn, ch := init_connection_channel()
	defer conn.Close()
	defer ch.Close()

	q := main_queue(ch)
	secondary_queue(ch)
	dlq := dead_letter_queue(ch)

	msgs := consume(ch, q.Name)
	dlq_msgs := consume(ch, dlq.Name)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message from user_notifications queue: %s", d.Body)
		}
	}()

	go func() {
		for d := range dlq_msgs {
			log.Printf("Received a message from dlq: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
