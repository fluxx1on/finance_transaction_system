package mq

import (
	"context"
	"fmt"
	"time"

	"github.com/fluxx1on/finance_transaction_system/internal/database"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/consumer"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type Scheduler struct {
	conn *amqp.Connection

	channels   []*amqp.Channel
	queueNames []string

	workers []*consumer.Worker

	routingKey string

	// Shutdown context manager
	shutdownContext context.Context

	shutdownCall context.CancelFunc
}

func NewScheduler(conn *amqp.Connection, lenChannel, lenWorkerByChannel int,
	exchngName, routingKey string, db *database.CreditDB) *Scheduler {
	var (
		channels   = make([]*amqp.Channel, 0, lenChannel)
		queueNames = make([]string, 0, lenChannel)
		workers    = make([]*consumer.Worker, 0, lenWorkerByChannel)
	)
	ctx, cancel := context.WithCancel(context.Background())

	for ch := 0; ch < lenChannel; ch++ {
		channel, err := conn.Channel()
		if err != nil {
			slog.Error("Failed to Open Channel", err)
			return nil
		}

		queueName := fmt.Sprintf("transaction%d", ch+1)
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
			return nil
		}

		err = channel.QueueBind(
			queueName,
			routingKey,
			exchngName,
			false,
			nil,
		)
		if err != nil {
			slog.Error("Failed to bind queue to exchange", err)
			return nil
		}

		channels = append(channels, channel)
		queueNames = append(queueNames, queueName)

		for w := 0; w < lenWorkerByChannel; w++ {
			id := (ch)*lenWorkerByChannel + w + 1
			workers = append(workers, consumer.NewWorker(
				id, db, channel, queueName, ctx,
			))
		}
	}

	return &Scheduler{
		conn:            conn,
		channels:        channels,
		queueNames:      queueNames,
		workers:         workers,
		shutdownContext: ctx,
		shutdownCall:    cancel,
	}
}

func (s *Scheduler) StartWorkers() {
	for _, worker := range s.workers {
		go worker.Consume()
	}
}

func (s *Scheduler) ShutdownJob() {
	// before shutdown need to save MQs
	s.shutdownContext.Done()

	slog.Info("Calling to stop Consumers...")
	<-time.After(2 * time.Second)

	s.shutdownCall()
}
