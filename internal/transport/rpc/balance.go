package rpc

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/service"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/balance"
	"github.com/fluxx1on/finance_transaction_system/internal/utils"
	"github.com/fluxx1on/finance_transaction_system/pkg/logger"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BalanceActionService struct {
	// Implements
	balance.UnimplementedBalanceActionServiceServer

	// Business logic handler
	fetcher *service.BalanceFetcher
}

func NewBalanceActionService(f *service.BalanceFetcher) *BalanceActionService {
	return &BalanceActionService{
		fetcher: f,
	}
}

func (s *BalanceActionService) Fill(ctx context.Context, req *balance.FillRequest) (
	*balance.FillResponse, error,
) {
	resp, err := s.fetcher.FetchFill(ctx, req)

	if err != nil {
		sttInfo := status.Convert(err)

		slog.Info(logger.GCodeSuite(utils.BalanceFill, sttInfo.Code()))
		return nil, status.Errorf(sttInfo.Code(), "%v", sttInfo.Err())
	}

	slog.Info(logger.GCodeSuite(utils.BalanceFill, codes.OK))
	return resp, nil
}

func (s *BalanceActionService) Get(ctx context.Context, req *balance.EmptyRequest) (
	*balance.AmountResponse, error,
) {
	resp, err := s.fetcher.FetchGet(ctx, req)

	if err != nil {
		sttInfo := status.Convert(err)

		slog.Info(logger.GCodeSuite(utils.BalanceGet, sttInfo.Code()))
		return nil, status.Errorf(sttInfo.Code(), "%v", sttInfo.Err())
	}

	slog.Info(logger.GCodeSuite(utils.BalanceGet, codes.OK))
	return resp, nil
}
