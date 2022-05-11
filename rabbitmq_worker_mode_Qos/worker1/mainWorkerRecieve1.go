package main

import (
	"learnrabbitmq/rabbitmq"
	"time"
)

func main() {
	rabbit := rabbitmq.NewRabbitMQSimple("workMode")
	rabbit.ConsumeSimpleQos(time.Millisecond * 500)
}
