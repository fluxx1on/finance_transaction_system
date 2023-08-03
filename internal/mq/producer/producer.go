package producer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fluxx1on/finance_transaction_system/internal/rpc"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type Producer struct {
	channel *amqp.Channel

	exchangeName string

	routingKey string
}

func NewProducer(conn *amqp.Connection, exchngName, routingKey string) *Producer {
	channel, err := conn.Channel()
	if err != nil {
		slog.Error("Failed to Open Channel")
		return nil
	}

	prod := &Producer{
		channel:      channel,
		exchangeName: exchngName,
		routingKey:   routingKey,
	}

	err = prod.DeclareExchange()
	if err != nil {
		slog.Error("Failed to declare exchange", err)
		return nil
	}

	return prod
}

func (p *Producer) DeclareExchange() error {
	err := p.channel.ExchangeDeclare(
		p.exchangeName, // Название вашего Exchange
		"direct",       // Type Exchange - direct
		true,           // Durable - true
		false,          // AutoDelete - false
		false,          // Internal - false
		false,          // NoWait - false
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) Publish(ctx context.Context, message rpc.TransactionInfo) error {
	slog.Info("Publish Method")

	messageJson, err := json.Marshal(message)
	if err != nil {
		slog.Warn("Message stuck in producer", err)
		return err
	}

	err = p.channel.PublishWithContext(
		ctx,
		p.exchangeName,
		p.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageJson,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
