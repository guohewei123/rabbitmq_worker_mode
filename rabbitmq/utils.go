package rabbitmq

import "github.com/streadway/amqp"

// declareQueue 创建队列（队列名随机）
func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
    return ch.QueueDeclare(
        "",    // 随机生成队列名称 queue name
        false, // 是否持久化  durable -> 进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
        false, // 是否自动删除 autoDelete -> 是否为自动删除  意思是最后一个消费者断开链接以后是否将消息从队列当中删除  默认设置为false不自动删除
        true,  // 是否具有排他性 exclusive
        false, // 是否阻塞 noWait -> 发送消息以后是否要等待消费者的响应 消费了下一个才进来 就跟golang里面的无缓冲channle一个道理 默认为非阻塞即可设置为false
        nil,   // 额外属性 args -> 没有则直接诶传入空即可 nil
    )
}

// createConsume 为队列创建消费者
func createConsume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
    return ch.Consume(
        queueName, // 队列名称
        "",        // 用来区分多个消费者
        true,      // autoAck 是否自动应答 -> 默认是true, 如果为true则是只要消费端消费了那么服务器自动删除服务端队列里面的消息, 如果为false表示消费端消费了消息之后需要手动ack一下 告诉服务器我消费成功了你可以去删除服务端队列里面的值了
        false,     // exclusive 是否具有排他性
        false,     // noLocal -> 如果设置为true, 表示不能将同一个connection中发送的消息传递给这个connection中消费
        false,     // noWait -> 队列消费是否阻塞 false表示是阻塞 true表示是不阻塞
        nil)
}

// declareExchange 创建交换机(存在则不创建)
func declareExchange(ch *amqp.Channel, exchange string, kind string) error {
    return ch.ExchangeDeclare(
        exchange, // exchange Name 交换机名称
        kind,     // exchange mode -> "fanout"（订阅模式/广播类型） "direct"（Routing模式） "topic"（topic模式）
        true,     // 进入的消息是否持久化 durable -> 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
        false,    // 是否自动删除 autoDelete -> exchange自动删除的条件，有队列或者交换器绑定了本交换器，然后所有队列或交换器都与本交换器解除绑定，autoDelete=true时，此交换器就会被自动删除
        false,    // internal: true 表示这个Exchange不可以被client用来推送消息，仅用来Exchange 和 Exchange 之间绑定
        false,    // noWait
        nil,      // args
    )
}
