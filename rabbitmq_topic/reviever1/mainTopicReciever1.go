package main

import "learnrabbitmq/rabbitmq"

func main() {
	rabbitMQOne := rabbitmq.NewRabbitMQTopic( "TopicExchange", "#")
	rabbitMQOne.ConsumeTopic()
}