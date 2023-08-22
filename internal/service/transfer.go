package service

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/mq/producer"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/serial"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/middleware"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/transfer"
)

type TransferFetcher struct {
	db *repo.CreditDB

	producer *producer.Producer
}

func NewTransferFetcher(db *repo.CreditDB, p *producer.Producer) *TransferFetcher {

	return &TransferFetcher{
		db:       db,
		producer: p,
	}
}

func (s *TransferFetcher) FetchTransfer(ctx context.Context, req *transfer.TransferRequest) (
	*transfer.TransferStatusResponse, error,
) {
	var (
		status bool = true
		errMsg string
	)

	user, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		status = false
		errMsg = err.Error()
	}

	var senderID, recipientID int
	if status {
		senderID, recipientID, err = s.db.PreTransfer(user.Username, req.RecipientName, int(req.TransferSum))
		if err != nil {
			status = false
			errMsg = err.Error()
		}
	}

	if status {
		message := serial.TransactionInfo{
			SenderID:         senderID,
			RecipientID:      recipientID,
			AmountToTransfer: int(req.TransferSum),
		}
		s.producer.Publish(ctx, message)
	}

	resp := &transfer.TransferStatusResponse{
		Status:       status,
		ErrorMessage: errMsg,
	}

	return resp, err
}
