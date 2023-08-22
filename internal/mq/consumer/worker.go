package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fluxx1on/finance_transaction_system/internal/mq/serial"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

const contextCanceled = "context canceled"

type Worker struct {
	id int

	db *repo.CreditDB

	ch *amqp.Channel

	queueName string

	ctx context.Context
}

func NewWorker(id int, db *repo.CreditDB, channel *amqp.Channel, queueName string, ctx context.Context) *Worker {
	return &Worker{
		id:        id,
		db:        db,
		ch:        channel,
		queueName: queueName,
		ctx:       ctx,
	}
}

func (w *Worker) runTransaction(task amqp.Delivery) error {

	if w.ctx.Err() != nil {
		return fmt.Errorf("") // TODO
	}
	var info serial.TransactionInfo
	err := json.Unmarshal(task.Body, &info)
	if err != nil {
		slog.Error("Unmarshaling Error in Worker")
	}

	if err = w.db.TransferTransaction(context.Background(), info); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err = task.Ack(true); err != nil {
		return fmt.Errorf("ack failed: %w", err)
	}

	slog.Info("Task Acknowlegment")

	return nil
}

func (w *Worker) Consume() {
	listener, err := w.ch.Consume(
		w.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		slog.Error("RabbitMQ Channel didn't consume", err)
	}

	for task := range listener {
		if err := w.runTransaction(task); err != nil {
			if err.Error() == contextCanceled {
				slog.Info(fmt.Sprintf("Worker#%d shutdown by context cancellation", w.id), err)
				return
			}

			slog.Error("Task completion failed", err)
		}
	}
}
