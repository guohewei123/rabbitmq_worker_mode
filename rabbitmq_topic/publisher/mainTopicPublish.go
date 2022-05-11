package main

import (
	"fmt"
	"learnrabbitmq/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbitMQOne := rabbitmq.NewRabbitMQTopic( "TopicExchange", "topic.key.one")
	rabbitMQTwo :=  rabbitmq.NewRabbitMQTopic( "TopicExchange", "topic.key.two")
	for i := 0; i <= 10; i++ {
		rabbitMQOne.PublishTopic("RoutingOne模式生成第" + strconv.Itoa(i) + "条数据")
		rabbitMQTwo.PublishTopic("RoutingTwo模式生成第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
		fmt.Println("Topic模式生成第" + strconv.Itoa(i) + "条数据")
	}
}
