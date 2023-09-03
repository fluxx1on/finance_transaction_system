package service

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/auth"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/producer"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/serial"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/transfer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		sttMsg bool = true
		errMsg string
	)

	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		sttMsg = false
		errMsg = err.Error()
	}

	var senderID, recipientID uint64
	if sttMsg {
		senderID, recipientID, err = s.db.PreTransfer(user.ID, req.RecipientName, req.TransferSum)
		if err != nil {
			sttMsg = false
			errMsg = err.Error()
		}
	}

	if sttMsg {
		message := serial.TransactionInfo{
			SenderID:         senderID,
			RecipientID:      recipientID,
			AmountToTransfer: req.TransferSum,
		}
		err = s.producer.Publish(ctx, message)
		if err != nil {
			sttMsg = false
			errMsg = status.Error(codes.Internal, "Something went wrong").Error()
		}
	}

	resp := &transfer.TransferStatusResponse{
		Status:       sttMsg,
		ErrorMessage: errMsg,
	}

	return resp, err
}
