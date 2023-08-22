package rpc

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/service"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/transfer"
	"golang.org/x/exp/slog"
)

type OperationService struct {
	// Implements
	transfer.UnimplementedOperationServiceServer

	// Business logic handler
	fetcher *service.OperationFetcher
}

func NewOperationService(f *service.OperationFetcher) *OperationService {
	return &OperationService{
		fetcher: f,
	}
}

func (s *OperationService) OperationList(ctx context.Context, req *transfer.EmptyRequest) (
	*transfer.OperationResponseList, error,
) {
	resp, err := s.fetcher.FetchOperationList(ctx, req)

	if err != nil { // requires resp isn't nil
		// TODO:log
		return nil, err
	}

	slog.Info("Response sent successfully") // TODO:log
	return resp, err
}
