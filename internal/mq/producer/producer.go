package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/fluxx1on/finance_transaction_system/internal/config"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/serial"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type Producer struct {
	channel *amqp.Channel

	exchangeName string

	routingKeys []string
}

func NewProducer(conn *amqp.Connection, cfg *config.RabbitConfig) (*Producer, error) {
	channel, err := conn.Channel()
	if err != nil {
		slog.Error("Failed to Open Channel")
		return nil, err
	}

	routingKeys := make([]string, 0, cfg.QueueAmount)

	for q := 0; q < cfg.QueueAmount; q++ {
		rk := fmt.Sprintf("%s:%d", cfg.RoutingKey, q)
		routingKeys = append(routingKeys, rk)
	}

	prod := &Producer{
		channel:      channel,
		exchangeName: cfg.ExchangeName,
		routingKeys:  routingKeys,
	}

	err = prod.DeclareExchange()
	if err != nil {
		slog.Error("Failed to declare exchange", err)
		return nil, err
	}

	return prod, nil
}

func (p *Producer) DeclareExchange() error {
	err := p.channel.ExchangeDeclare(
		p.exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) Publish(ctx context.Context, message serial.TransactionInfo) error {
	slog.Debug("Publish Method")

	messageJson, err := json.Marshal(message)
	if err != nil {
		slog.Warn("Message stuck in producer", err)
		return err
	}

	randWorkerID := rand.Intn(len(p.routingKeys)) // Queue choice

	err = p.channel.PublishWithContext(
		ctx,
		p.exchangeName,
		p.routingKeys[randWorkerID],
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
