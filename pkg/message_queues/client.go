package rabbitmq

import (
	"fmt"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// create a publisher
// create a consumer

type RabbitMQClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// create a rabbit mq connection
func NewRabbitMQ(uri string) (*RabbitMQClient, error) {
	conn, err := connect(uri)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQClient{Conn: conn, Channel: ch}, nil
}

// declare an exchange
func (c *RabbitMQClient) DeclareExchange(name string, kind string, durable bool) error {
	ch := c.Channel
	// defer ch.Close()
	return ch.ExchangeDeclare(
		name,    // name of the exchange
		kind,    // kind of exchange
		durable, // durable exchange
		false,   // autoDelete
		false,   // internal exchange
		false,   // noWait
		nil,     // arguments
	)
}

// close connection and channel
func (c *RabbitMQClient) Close() {

	if c.Channel != nil {
		c.Channel.Close()
	}
	if c.Conn != nil {
		c.Conn.Close()
	}
}

// Bind QueueToExchange binds a queue to an exchange with a specified key
func (c *RabbitMQClient) QueueBind(queue amqp.Queue, exchange, key string) error {
	return c.Channel.QueueBind(queue.Name, key, exchange, false, nil)
}

// CreateQueue creates a new queue on the RabbitMQ server
func (c *RabbitMQClient) CreateQueue(queue string) (amqp.Queue, error) {
	q, err := c.Channel.QueueDeclare(
		queue, // name of queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	return q, err
}

type ConsumerChannel struct {
	msgs <-chan amqp.Delivery
}

// create a consumer for message in queue
func (c *RabbitMQClient) Consume(queue amqp.Queue, autoAck bool) (*ConsumerChannel, error) {
	msgs, err := c.Channel.Consume(
		queue.Name, // name
		"",         // consumerTag
		autoAck,    // autoAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // args
	)

	return &ConsumerChannel{msgs: msgs}, err
}

// read from the consumer and process  each message
func (consumer *ConsumerChannel) ProcessEachMessage(f func(delivery amqp.Delivery)) {
	for d := range consumer.msgs {
		go f(d)
	}
}

func connect(uri string) (*amqp.Connection, error) {
	var connection *amqp.Connection
	var counts int64
	var backOff = 1 * time.Second

	for {
		c, err := amqp.Dial(uri)
		if err != nil {
			fmt.Println("Rabbitmq not ready")
			counts++
		} else {
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err.Error() + " after retrying 5 times.")
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		time.Sleep(backOff)
	}
	// otelamqpp.WrapConnection(connection)

	return connection, nil

}
