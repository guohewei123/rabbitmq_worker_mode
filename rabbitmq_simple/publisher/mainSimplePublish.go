package main

import (
	"fmt"
	"learnrabbitmq/rabbitmq"
)

func main() {
	rabbit := rabbitmq.NewRabbitMQSimple("SimpleMode")
	rabbit.PublishSimple("Hello SimpleMode!")
	fmt.Println("发送成功")
}
