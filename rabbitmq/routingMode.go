package rabbitmq

import (
    "github.com/streadway/amqp"
    "log"
)

// NewRabbitMQRouting 订阅模式step1，创建简单RabbitMQ
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
    return NewRabbitMQ("", exchangeName, routingKey)
}

// PublishRouting 订阅模式step2，生成者代码
func (r *RabbitMQ) PublishRouting(message string) {
    // 1. 尝试创建交换机
    err := declareExchange(r.channel, r.Exchange, "direct")
    r.failOnErr(err, "Failed to declare an exchange")

    // 2. 发送消息到队列中
    err = r.channel.Publish(
        r.Exchange,
        r.Key,
        false, // 如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
        false, // 如果为true，当Exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
}

// ConsumeRouting 简单模式step3，消费代码
func (r *RabbitMQ) ConsumeRouting() {
    // 1. 尝试创建交换机
    err := declareExchange(r.channel, r.Exchange, "direct")
    r.failOnErr(err, "Failed to declare an exchange")

    // 2. 试探申请队列，这里注意队列名称不要写
    q, err := declareQueue(r.channel)
    r.failOnErr(err, "Failed to declare a queue")

    // 绑定队列到exchange 中
    err = r.channel.QueueBind(
        q.Name,     // 队列名称
        r.Key,      // routing Key
        r.Exchange, // 交换机名称
        false,
        nil,
    )

    // 2.接收消息
    msgs, err := createConsume(r.channel, q.Name)
    r.failOnErr(err, "Create consume Failed")

    forever := make(chan bool)
    // 启用协程处理消息
    go func() {
        for d := range msgs {
            // 实现我们要处理的逻辑函数
            log.Printf("Received a message: %s", d.Body)
        }
    }()
    log.Printf("[*] waiting for messages, [退出请按]To exit press CTRL+C")
    <-forever
}
