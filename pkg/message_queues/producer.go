package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Emitter struct {
	Client *RabbitMQClient
}

func SetUpEmitter(client *RabbitMQClient, e *Exchange) (*Emitter, error) {

	err := client.DeclareExchange(e.ExchangeName, e.Kind, e.Durable)

	if err != nil {
		client.Close()
		return nil, err
	}

	return &Emitter{Client: client}, nil

}

func (c *Emitter) PublishMessage(ctx context.Context, exchangeName string, key string, msg []byte) error {
	tp := otel.GetTracerProvider()
	ctx, span := tp.Tracer("rabbitmq").Start(ctx, "Publishing message", trace.WithAttributes(
		attribute.String("exchange_name", exchangeName),
		attribute.String("routing_key", key),
	),
		trace.WithSpanKind(trace.SpanKindProducer),
	)
	defer span.End()
	ch := c.Client.Channel
	err := ch.PublishWithContext(
		ctx,
		exchangeName, // publish to an exchange
		key,          // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	return err
}
