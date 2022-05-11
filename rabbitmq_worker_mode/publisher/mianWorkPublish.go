package main

import (
	"fmt"
	"learnrabbitmq/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbit := rabbitmq.NewRabbitMQSimple("workMode")

	for i := 0; i <= 100; i++ {
		rabbit.PublishSimple("Hello workMode!" + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 1)
		fmt.Println(i)
	}
}
