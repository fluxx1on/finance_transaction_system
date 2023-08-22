package mq

import (
	"context"
	"fmt"
	"time"

	"github.com/fluxx1on/finance_transaction_system/internal/config"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/consumer"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type ManagerMQ struct {
	conn *amqp.Connection

	channel    *amqp.Channel
	queueNames []string

	workers []*consumer.Worker

	// Shutdown context manager
	shutdownContext context.Context

	shutdownCall context.CancelFunc
}

func NewManagerMQ(conn *amqp.Connection, cfg *config.RabbitClient, db *repo.CreditDB) (*ManagerMQ, error) {
	var (
		queueNames = make([]string, 0, cfg.QueueAmount)
		workers    = make([]*consumer.Worker, 0, cfg.WorkerByChannelAmount)
	)
	ctx, cancel := context.WithCancel(context.Background())

	channel, err := conn.Channel()
	if err != nil {
		slog.Error("Failed to Open Channel", err)
		cancel()
		return nil, err
	}

	for q := 0; q < cfg.QueueAmount; q++ {

		queueName := fmt.Sprintf("transaction%d", q+1)
		_, err = channel.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			slog.Error("Failed to declare queue", err)
			cancel()
			return nil, err
		}

		rk := fmt.Sprintf("%s:%d", cfg.RoutingKey, q)

		err = channel.QueueBind(
			queueName,
			rk,
			cfg.ExchangeName,
			false,
			nil,
		)
		if err != nil {
			slog.Error("Failed to bind queue to exchange", err)
			cancel()
			return nil, err
		}
		queueNames = append(queueNames, queueName)

		for w := 0; w < cfg.WorkerByChannelAmount; w++ {
			id := (q)*cfg.WorkerByChannelAmount + w + 1
			workers = append(workers, consumer.NewWorker(
				id, db, channel, queueName, ctx,
			))
		}
	}

	return &ManagerMQ{
		conn:            conn,
		channel:         channel,
		queueNames:      queueNames,
		workers:         workers,
		shutdownContext: ctx,
		shutdownCall:    cancel,
	}, nil
}

func (s *ManagerMQ) StartWorkers() {
	for _, worker := range s.workers {
		go worker.Consume()
	}
}

func (s *ManagerMQ) ShutdownJob() {
	// before shutdown need to save MQs
	slog.Info("Calling to stop Consumers...")
	s.shutdownCall()

	<-time.After(2 * time.Second)

	if err := s.channel.Close(); err != nil {
		slog.Error("Problems with closing rabbitmq channel", err)
	}

	if err := s.conn.Close(); err != nil {
		slog.Error("Problems with closing rabbitmq connection", err)
	}
}
