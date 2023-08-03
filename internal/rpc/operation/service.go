package operation

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/rpc/pb"
	"golang.org/x/exp/slog"
)

type OperationService struct {
	// Implements
	pb.UnimplementedOperationServiceServer

	// Business logic handler
	fetcher *OperationFetcher
}

func NewOperationService(f *OperationFetcher) *OperationService {
	return &OperationService{
		fetcher: f,
	}
}

func (s *OperationService) OperationList(ctx context.Context, req *pb.OperationRequestList) (
	*pb.OperationResponseList, error,
) {
	resp, err := s.fetcher.FetchOperationList(ctx, req)

	if err != nil { // requires resp isn't nil
		// TODO:log
		return nil, err
	}

	slog.Info("Response sent succesfully") // TODO:log
	return resp, err
}
