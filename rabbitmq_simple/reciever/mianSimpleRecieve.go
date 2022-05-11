package main

import (
	"learnrabbitmq/rabbitmq"
	"time"
)

func main() {
	rabbit := rabbitmq.NewRabbitMQSimple("SimpleMode")
	rabbit.ConsumeSimple(time.Millisecond, true)
}
