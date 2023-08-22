package service

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/auth"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserFetcher struct {
	db *repo.CreditDB
}

func NewUserFetcher(db *repo.CreditDB) *UserFetcher {

	return &UserFetcher{
		db: db,
	}
}

func (f *UserFetcher) FetchRegister(ctx context.Context, req *user.SignUpRequest) (
	*user.SignUpResponse, error,
) {

	// TODO
	return nil, nil
}

func (f *UserFetcher) FetchLogin(ctx context.Context, req *user.SignInRequest) (
	*user.SignInResponse, error,
) {
	var errMsg string

	rtoken, atoken, err := auth.Authentication(f.db, req.Username, req.Password)
	if err != nil {
		errMsg = err.Error()
	}

	tokens := &user.Tokens{
		RefreshToken:    rtoken.Token,
		RefreshExpiryAt: timestamppb.New(rtoken.ExpiredAt),
		AccessToken:     atoken.Token,
		AccessExpiryAt:  timestamppb.New(atoken.ExpiredAt),
	}
	resp := &user.SignInResponse{
		Tokens:       tokens,
		ErrorMessage: errMsg,
	}
	return resp, err
}
