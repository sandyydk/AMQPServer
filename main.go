package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://admin:admin@192.168.99.100:5672"
	}

	log.Println(amqpURL)
	connection, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal("Error connecting to amqp -", err.Error())
	}

	// Open a channel to publish messages
	channel, err := connection.Channel()
	if err != nil {
		log.Println("Error creating a amqp channel - ", err.Error())
		return
	}

	defer channel.Close()
	// Declare exchange to which the publisher(you) intend to publish messages. Possible Types :- "topic", "direct", "fanout"
	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	if err != nil {
		log.Println("Error declaring an exchange - ", err.Error())
		panic(err)
	}

	message := amqp.Publishing{
		Body: []byte("Sample test message"),
	}

	err = channel.Publish("events", "some-routing-key", false, false, message)
	if err != nil {
		log.Println("Error publishing a message to the events exchange - ", err.Error())
	}

	log.Println("Published message to AMQP")

}
