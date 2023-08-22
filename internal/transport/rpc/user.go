package rpc

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/service"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/user"
	"golang.org/x/exp/slog"
)

type UserService struct {
	// Implements
	user.UnimplementedUserServiceServer

	// Business logic handler
	fetcher *service.UserFetcher
}

func NewUserService(f *service.UserFetcher) *UserService {
	return &UserService{
		fetcher: f,
	}
}

func (s *UserService) Register(ctx context.Context, req *user.SignUpRequest) (
	*user.SignUpResponse, error,
) {
	resp, err := s.fetcher.FetchRegister(ctx, req)

	if err != nil {
		// TODO:log
		return nil, err
	}

	slog.Info("") // TODO:log
	return resp, err
}

func (s *UserService) Login(ctx context.Context, req *user.SignInRequest) (
	*user.SignInResponse, error,
) {
	resp, err := s.fetcher.FetchLogin(ctx, req)

	if err != nil {
		// TODO:log
		return nil, err
	}

	slog.Info("") // TODO:log
	return resp, err
}
