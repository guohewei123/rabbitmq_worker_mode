package rabbitmq

import (
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "time"
)

// NewRabbitMQSimple 简单模式step1，创建简单RabbitMQ
func NewRabbitMQSimple(queueName string) *RabbitMQ {
    //simple模式下交换机为空，因为会默认使用rabbitmq默认的default交换机，而不是真的没有bindkey，绑定key也是为空的
    //特别注意：simple模式是最简单的rabbitmq的一种模式 他只需要传递queue队列名称过去即可  像exchange交换机会默认使用default交换机  绑定建key的会不必要传
    return NewRabbitMQ(queueName, "", "")
}

// PublishSimple 简单模式step2，生成者代码
func (r *RabbitMQ) PublishSimple(message string) {
    // 1. 申请队列，如果队列不存在会自动创建，如果存在则跳过创建直接使用
    // 保证队列存在，消息能发送到队列中
    _, err := r.channel.QueueDeclare(
        r.QueueName,
        false, // 是否持久化 -> 进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
        false, // 是否自动删除 -> 是否为自动删除  意思是最后一个消费者断开链接以后是否将消息从队列当中删除  默认设置为false不自动删除
        false, // 是否具有排他性 ->
        false, // 是否阻塞 -> 发送消息以后是否要等待消费者的响应 消费了下一个才进来 就跟golang里面的无缓冲channle一个道理 默认为非阻塞即可设置为false
        nil,   // 额外属性 -> 没有则直接诶传入空即可 nil
    )
    if err != nil {
        fmt.Println(err)
    }
    // 2. 发送消息到队列中
    err = r.channel.Publish(
        r.Exchange, // 交换机 simple模式下默认为空 我们在上边已经赋值为空了  虽然为空 但其实也是在用的rabbitmq当中的default交换机运行
        r.QueueName,
        false, // 如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
        false, // 如果为true，当Exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
        amqp.Publishing{ //要发送的消息
            ContentType: "text/plain",
            Body:        []byte(message),
        })
    if err != nil {
        fmt.Println(err)
    }
}

// ConsumeSimple 简单模式step3，消费代码
func (r *RabbitMQ) ConsumeSimple(duration time.Duration, autoAck bool) {
    // 1. 申请队列，如果队列不存在会自动创建，如果存在则跳过创建
    // 保证队列存在，消息能发送到队列中
    _, err := r.channel.QueueDeclare(
        r.QueueName,
        false, // 是否持久化
        false, // 是否自动删除
        false, // 是否具有排他性
        false, // 是否阻塞
        nil,   // 额外属性
    )
    if err != nil {
        fmt.Println(err)
    }

    // 2.接收消息 queue string, consumer string, autoAck bool, exclusive bool, noLocal bool,
    msgs, err := r.channel.Consume(
        r.QueueName, // queue 队列名称
        "",          // consumer 用来区分多个消费者
        autoAck,     // autoAck 是否自动应答
        false,       // exclusive 是否具有排他性
        false,       // noLocal 如果设置为true, 表示不能将同一个connection中生产者发送的消息传递给这个connection中消费
        false,       // noWait 队列消费是否阻塞 false 表示阻塞
        nil)
    if err != nil {
        fmt.Println(err)
    }

    forever := make(chan bool)
    // 启用协程处理消息
    go func() {
        /*for true {
            d := <-msgs
            log.Printf("Received a message: %s", d.Body)
            time.Sleep(duration)
            if !autoAck {
                r.channel.Ack(d.DeliveryTag, false)
            }
        }*/

        for d := range msgs {
            // 实现我们要处理的逻辑函数
            log.Printf("Received a message: %s", d.Body)
            time.Sleep(duration)
            if !autoAck {
                //如果为true表示确认所有未确认的消息
                //如果为false表示确认当前消息
                //执行完业务逻辑成功之后我们再手动ack告诉服务器你可以删除这个消息啦！ 这样就保障了数据的绝对的安全不丢失！
                //err := r.channel.Ack(d.DeliveryTag, false)
                err := d.Ack(false)
                if err != nil {
                    println(err)
                }
            }
        }
    }()
    log.Printf("[*] waiting for messages, [退出请按]To exit press CTRL+C")
    <-forever
}

// ConsumeSimpleQos 简单模式step3，消费代码
func (r *RabbitMQ) ConsumeSimpleQos(duration time.Duration) {
    // 1. 申请队列，如果队列不存在会自动创建，如果存在则跳过创建
    // 保证队列存在，消息能发送到队列中
    _, err := r.channel.QueueDeclare(
        r.QueueName,
        false, // 是否持久化
        false, // 是否自动删除
        false, // 是否具有排他性
        false, // 是否阻塞
        nil,   // 额外属性
    )
    if err != nil {
        fmt.Println(err)
    }

    // 关于Qos参数含义如何设置: https://www.cnblogs.com/throwable/p/13834465.html
    // prefetchCount参数的设置: https://blog.rabbitmq.com/posts/2014/04/finding-bottlenecks-with-rabbitmq-3-3
    // 官方是建议把prefetchCount设置为30
    r.channel.Qos(
        3,     //prefetchCount 当前消费者一次能接受的最大消息数量
        0,     //prefetchSize 用于限制分发内容的大小上限，默认值0代表无限制
        // global -> 控制 prefetchCount 和prefetchSize 的生效范围：
        // 如果为true (prefetchCount值在当前Channel的所有消费者共享--所有消费者加起来等于prefetchCount)
        // 如果为false (prefetchCount对于基于当前Channel创建的消费者生效--每个消费者都是prefetchCount)
        true,
    )

    // 2.接收消息 queue string, consumer string, autoAck bool, exclusive bool, noLocal bool,
    msgs, err := r.channel.Consume(
        r.QueueName, // queue 队列名称
        "",          // consumer 用来区分多个消费者
        false,       // autoAck 是否自动应答
        false,       // exclusive 是否具有排他性
        false,       // noLocal 如果设置为true, 表示不能将同一个connection中生产者发送的消息传递给这个connection中消费
        false,       // noWait 队列消费是否阻塞 false 表示阻塞
        nil)
    if err != nil {
        fmt.Println(err)
    }

    forever := make(chan bool)
    // 启用协程处理消息
    go func() {
        for d := range msgs {
            // 实现我们要处理的逻辑函数
            log.Printf("Received a message: %s", d.Body)
            time.Sleep(duration)
            //如果为true表示确认所有未确认的消息
            //如果为false表示确认当前消息
            //执行完业务逻辑成功之后我们再手动ack告诉服务器你可以删除这个消息啦！ 这样就保障了数据的绝对的安全不丢失！
            err := d.Ack(false)
            if err != nil {
                println(err)
            }
        }
    }()
    log.Printf("[*] waiting for messages, [退出请按]To exit press CTRL+C")
    <-forever
}
