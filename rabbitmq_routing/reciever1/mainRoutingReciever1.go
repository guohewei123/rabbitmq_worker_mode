package main

import "learnrabbitmq/rabbitmq"

func main() {
	rabbitMQOne := rabbitmq.NewRabbitMQRouting( "routingExchange", "routingKeyOne")
	rabbitMQOne.ConsumeRouting()
}