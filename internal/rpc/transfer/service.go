package transfer

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/rpc/pb"
	"golang.org/x/exp/slog"
)

type TransferService struct {
	// Implements
	pb.UnimplementedTransferServiceServer

	// Business logic handler
	fetcher *TransferFetcher
}

func NewTransferService(f *TransferFetcher) *TransferService {
	return &TransferService{
		fetcher: f,
	}
}

func (s *TransferService) Transfer(ctx context.Context, req *pb.TransferRequest) (
	*pb.TransferStatusResponse, error,
) {
	resp, err := s.fetcher.FetchTransfer(ctx, req)

	if err != nil { // requires respList isn't nil
		// TODO:log
		return nil, err
	}

	slog.Info("Transaction succesfully completed") // TODO:log
	return resp, err
}
