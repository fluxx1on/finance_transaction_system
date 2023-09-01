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

	if err != nil {
		sttInfo := status.Convert(err)

		slog.Info(logger.GCodeSuite(utils.TransferTransfer, sttInfo.Code()))
		return nil, status.Errorf(sttInfo.Code(), "%v", sttInfo.Err())
	}

	slog.Info(logger.GCodeSuite(utils.TransferTransfer, codes.OK))
	return resp, nil
}
