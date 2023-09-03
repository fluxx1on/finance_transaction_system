package service

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/auth"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/balance"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	var (
		sttMsg bool = true
		errMsg string
	)

	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		sttMsg = false
		errMsg = err.Error()
	}

	if sttMsg {
		err = s.db.FillBalance(user.ID, req.Amount)
		if err != nil {
			sttMsg = false
			errMsg = status.Error(codes.Canceled, "Something went wrong. Try later").Error()
		}
	}

	resp := &balance.FillResponse{
		Status:       sttMsg,
		ErrorMessage: errMsg,
	}

	return resp, err
}

func (s *BalanceFetcher) FetchGet(ctx context.Context, req *balance.EmptyRequest) (
	*balance.AmountResponse, error,
) {
	var (
		sttMsg bool = true
		errMsg string
	)

	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		sttMsg = false
		errMsg = err.Error()
	}

	var balanceValue int32
	if sttMsg {
		balanceValue, err = s.db.GetBalance(user.ID)
		if err != nil {
			sttMsg = false
			errMsg = status.Error(codes.NotFound, err.Error()).Error()
		}
	}

	resp := &balance.AmountResponse{
		Status:       sttMsg,
		Value:        balanceValue,
		ErrorMessage: errMsg,
	}

	return resp, err
}
