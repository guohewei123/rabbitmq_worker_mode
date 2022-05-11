package main

import (
	"learnrabbitmq/rabbitmq"
	"time"
)

func main() {
	rabbit := rabbitmq.NewRabbitMQSimple("workMode")
	rabbit.ConsumeSimple(time.Millisecond * 500, false)
}
