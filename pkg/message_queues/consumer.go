package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type Exchange struct {
	ExchangeName string
	Kind         string
	Durable      bool
}

type Queue struct {
	QueueName  string
	QueBindKey string
}

type ProcessFunction func(delivery amqp.Delivery)

type Consumer struct {
	Client *RabbitMQClient
}

func NewConsumer(client *RabbitMQClient) (*Consumer) {
	return &Consumer{client}
}

func (c *Consumer) ConsumerListen(e *Exchange, q *Queue, processFunction ProcessFunction) error {
	client := c.Client
	err := c.Client.DeclareExchange(e.ExchangeName, e.Kind, e.Durable)
	if err != nil {
		return err
	}

	qu, err := client.CreateQueue(q.QueueName)
	if err != nil {
		return err
	}

	err = client.QueueBind(qu, e.ExchangeName, q.QueBindKey)
	if err != nil {
		return err
	}

	consumerChannel, err := client.Consume(qu, true)
	if err != nil {
		return err
	}
	forever := make(chan bool)
	go consumerChannel.ProcessEachMessage(processFunction)

	<-forever
	return nil
}
