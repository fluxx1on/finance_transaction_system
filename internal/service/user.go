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

	var (
		sttMsg bool
		errMsg string
	)

	// TODO : Add Password2 to auth.Register arguments (also in .proto)
	rtoken, atoken, err := auth.Register(f.db, req.Username, req.Password, req.Password)
	if err != nil {
		sttMsg = false
		errMsg = err.Error()
	}

	var tokens *user.Tokens
	if sttMsg {
		tokens = &user.Tokens{
			RefreshToken:    rtoken.Token,
			RefreshExpiryAt: timestamppb.New(rtoken.ExpiredAt),
			AccessToken:     atoken.Token,
			AccessExpiryAt:  timestamppb.New(atoken.ExpiredAt),
		}
	}

	resp := &user.SignUpResponse{
		Status:       sttMsg,
		Tokens:       tokens,
		ErrorMessage: errMsg,
	}

	return resp, err
}

func (f *UserFetcher) FetchLogin(ctx context.Context, req *user.SignInRequest) (
	*user.SignInResponse, error,
) {
	var (
		sttMsg bool
		errMsg string
	)

	rtoken, atoken, err := auth.Authentication(f.db, req.Username, req.Password)
	if err != nil {
		sttMsg = false
		errMsg = err.Error()
	}

	var tokens *user.Tokens
	if sttMsg {
		tokens = &user.Tokens{
			RefreshToken:    rtoken.Token,
			RefreshExpiryAt: timestamppb.New(rtoken.ExpiredAt),
			AccessToken:     atoken.Token,
			AccessExpiryAt:  timestamppb.New(atoken.ExpiredAt),
		}
	}

	resp := &user.SignInResponse{
		Status:       sttMsg,
		Tokens:       tokens,
		ErrorMessage: errMsg,
	}

	return resp, err
}
