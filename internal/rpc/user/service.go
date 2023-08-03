package user

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/rpc/pb"
	"golang.org/x/exp/slog"
)

type UserService struct {
	// Implements
	pb.UnimplementedUserServiceServer

	// Business logic handler
	fetcher *UserFetcher
}

func NewUserService(f *UserFetcher) *UserService {
	return &UserService{
		fetcher: f,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.SignUpRequest) (
	*pb.SignUpResponse, error,
) {
	resp, err := s.fetcher.FetchCreate(ctx, req)

	if err != nil {
		// TODO:log
		return nil, err
	}

	slog.Info("") // TODO:log
	return resp, err
}

func (s *UserService) Login(ctx context.Context, req *pb.SignInRequest) (
	*pb.SignInResponse, error,
) {
	resp, err := s.fetcher.FetchLogin(ctx, req)

	if err != nil {
		// TODO:log
		return nil, err
	}

	slog.Info("") // TODO:log
	return resp, err
}
