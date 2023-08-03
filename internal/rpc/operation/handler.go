package operation

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/database"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/pb"
)

type OperationFetcher struct {
	db *database.CreditDB
}

func NewOperationFetcher(db *database.CreditDB) *OperationFetcher {

	return &OperationFetcher{
		db: db,
	}
}

func (s *OperationFetcher) FetchOperationList(ctx context.Context, req *pb.OperationRequestList) (
	*pb.OperationResponseList, error,
) {

	// TODO
	return nil, nil
}
