package main

import "learnrabbitmq/rabbitmq"

func main() {
	RabbitMQTwo :=  rabbitmq.NewRabbitMQRouting( "routingExchange", "routingKeyTwo")
	RabbitMQTwo.ConsumeRouting()
}