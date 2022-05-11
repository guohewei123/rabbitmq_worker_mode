package main

import "learnrabbitmq/rabbitmq"

func main() {
	rabbit := rabbitmq.NewRabbitMQPubSub("newProduct")
	rabbit.ConsumeSub()
}