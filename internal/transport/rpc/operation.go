package rpc

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/service"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/transfer"
	"github.com/fluxx1on/finance_transaction_system/internal/utils"
	"github.com/fluxx1on/finance_transaction_system/pkg/logger"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	if err != nil {
		sttInfo := status.Convert(err)

		slog.Info(logger.GCodeSuite(utils.TransferTransfer, sttInfo.Code()))
		return nil, status.Errorf(sttInfo.Code(), "%v", sttInfo.Err())
	}

	slog.Info(logger.GCodeSuite(utils.OperationOperationList, codes.OK))
	return resp, nil
}
