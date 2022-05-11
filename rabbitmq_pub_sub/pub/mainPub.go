package main

import (
	"fmt"
	"learnrabbitmq/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbit := rabbitmq.NewRabbitMQPubSub("newProduct")

	for i := 0; i <= 10; i++ {
		rabbit.PublishPub("订阅模式生成第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
		fmt.Println("订阅模式生成第" + strconv.Itoa(i) + "条数据")
	}
}
