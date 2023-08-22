package service

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/transfer"
)

type OperationFetcher struct {
	db *repo.CreditDB
}

func NewOperationFetcher(db *repo.CreditDB) *OperationFetcher {

	return &OperationFetcher{
		db: db,
	}
}

func (s *OperationFetcher) FetchOperationList(ctx context.Context, req *transfer.EmptyRequest) (
	*transfer.OperationResponseList, error,
) {

	// TODO
	return nil, nil
}
