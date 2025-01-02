package main

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}

func main() {

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	id, err := uuid.NewV7()
	handleError(err, "Failed to generate UUID")

	channel, err := conn.Channel()
	handleError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"person",
		false,
		false,
		false,
		false,
		nil)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		publishPerson(id, err, channel, ctx, queue)
	}
}

func handleError(err error, msg string) {
	if err != nil {
		log.Panicf("%s : %s", msg, err)
	}
}

func publishPerson(id uuid.UUID, err error, channel *amqp091.Channel, ctx context.Context, queue amqp091.Queue) {
	person := Person{
		ID:   id,
		Name: "",
		Age:  0,
	}

	body, err := json.Marshal(person)
	handleError(err, "Failed to marshal JSON")

	err = channel.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	handleError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
