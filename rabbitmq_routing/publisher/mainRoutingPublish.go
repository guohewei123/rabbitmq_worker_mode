package main

import (
	"fmt"
	"learnrabbitmq/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbitMQOne := rabbitmq.NewRabbitMQRouting( "routingExchange", "routingKeyOne")
	RabbitMQTwo :=  rabbitmq.NewRabbitMQRouting( "routingExchange", "routingKeyTwo")
	for i := 0; i <= 10; i++ {
		rabbitMQOne.PublishRouting("RoutingOne模式生成第" + strconv.Itoa(i) + "条数据")
		RabbitMQTwo.PublishRouting("RoutingTwo模式生成第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
		fmt.Println("Routing模式生成第" + strconv.Itoa(i) + "条数据")
	}
}
