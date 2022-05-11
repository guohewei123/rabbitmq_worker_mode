package rabbitmq

import (
    "fmt"
    "github.com/streadway/amqp"
    "log"
)

// MQURL url 格式 "amqp://账号:密码@RabbitMQ服务器地址:端口号/Virtual Host"
//const MQURL = "amqp://username:password@192.168.10.129:5672/testmq"
//const MQURL = "amqp://guest:guest@192.168.10.53:15672//"
const MQURL = "amqp://ghw:123456@192.168.10.48:5672/test"

type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    // 队列名称
    QueueName string
    // 交换机
    Exchange string
    // 连接信息
    AmqpURL string
    // key
    Key string
}

// NewRabbitMQ 创建RabbitMQ 结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
    rabbitMQ := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, AmqpURL: MQURL}
    var err error
    // 创建rabbitMQ 连接
    rabbitMQ.conn, err = amqp.Dial(rabbitMQ.AmqpURL)
    rabbitMQ.failOnErr(err, "创建链接错误！")
    rabbitMQ.channel, err = rabbitMQ.conn.Channel()
    rabbitMQ.failOnErr(err, "获取channel失败！")
    return rabbitMQ
}

// Destroy 断开channel and connection链接
func (r *RabbitMQ) Destroy() {
    err := r.channel.Close()  //关闭信道资源
    if err != nil {
        fmt.Printf("cannot close chanel: %v", err)
    }
    err = r.conn.Close()   //关闭连接资源
    if err != nil {
        fmt.Printf("cannot close conn: %v", err)
    }
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
    if err != nil {
        log.Printf("%s:%s", message, err)
        panic(fmt.Sprintf("%s:%s", message, err))
    }

}
