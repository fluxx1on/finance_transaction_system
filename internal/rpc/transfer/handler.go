package transfer

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/database"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/producer"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/pb"
)

type TransferFetcher struct {
	db *database.CreditDB

	producer *producer.Producer
}

func NewTransferFetcher(db *database.CreditDB, p *producer.Producer) *TransferFetcher {

	return &TransferFetcher{
		db:       db,
		producer: p,
	}
}

func (s *TransferFetcher) FetchTransfer(ctx context.Context, req *pb.TransferRequest) (
	*pb.TransferStatusResponse, error,
) {
	var (
		status bool = true
		errMsg string
	)

	senderID, recipientID, err := s.db.PreTransfer(req.UserToken, req.RecipientName, int(req.TransferSum))
	if err != nil {
		status = false
		errMsg = err.Error()
	}

	message := rpc.TransactionInfo{
		SenderID:         senderID,
		RecipientID:      recipientID,
		AmountToTransfer: int(req.TransferSum),
	}
	s.producer.Publish(ctx, message)

	resp := &pb.TransferStatusResponse{
		Status:       status,
		ErrorMessage: errMsg,
	}

	return resp, err
}
