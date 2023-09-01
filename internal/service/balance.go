package service

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/balance"
)

type BalanceFetcher struct {
	db *repo.CreditDB
}

func NewBalanceFetcher(db *repo.CreditDB) *BalanceFetcher {

	return &BalanceFetcher{
		db: db,
	}
}

func (s *BalanceFetcher) FetchFill(ctx context.Context, req *balance.FillRequest) (
	*balance.FillResponse, error,
) {

	// TODO
	return nil, nil
}

func (s *BalanceFetcher) FetchGet(ctx context.Context, req *balance.EmptyRequest) (
	*balance.AmountResponse, error,
) {

	// TODO
	return nil, nil
}
