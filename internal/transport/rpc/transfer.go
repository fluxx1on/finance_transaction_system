package rpc

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/service"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/transfer"
	"golang.org/x/exp/slog"
)

type TransferService struct {
	// Implements
	transfer.UnimplementedTransferServiceServer

	// Business logic handler
	fetcher *service.TransferFetcher
}

func NewTransferService(f *service.TransferFetcher) *TransferService {
	return &TransferService{
		fetcher: f,
	}
}

func (s *TransferService) Transfer(ctx context.Context, req *transfer.TransferRequest) (
	*transfer.TransferStatusResponse, error,
) {
	resp, err := s.fetcher.FetchTransfer(ctx, req)

	if err != nil { // requires respList isn't nil
		// TODO:log
		return nil, err
	}

	slog.Info("Transaction successfully completed") // TODO:log
	return resp, err
}
