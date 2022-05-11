package main

import "learnrabbitmq/rabbitmq"

func main() {
	RabbitMQTwo :=  rabbitmq.NewRabbitMQTopic( "TopicExchange", "topic.*.two")
	RabbitMQTwo.ConsumeTopic()
}