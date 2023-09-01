package rpc

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/service"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/user"
	"github.com/fluxx1on/finance_transaction_system/internal/utils"
	"github.com/fluxx1on/finance_transaction_system/pkg/logger"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		sttInfo := status.Convert(err)

		slog.Info(logger.GCodeSuite(utils.UserRegister, sttInfo.Code()))
		return nil, status.Errorf(sttInfo.Code(), "%v", sttInfo.Err())
	}

	slog.Info(logger.GCodeSuite(utils.UserRegister, codes.OK))
	return resp, nil
}

func (s *UserService) Login(ctx context.Context, req *user.SignInRequest) (
	*user.SignInResponse, error,
) {
	resp, err := s.fetcher.FetchLogin(ctx, req)

	if err != nil {
		sttInfo := status.Convert(err)

		slog.Info(logger.GCodeSuite(utils.UserLogin, sttInfo.Code()))
		return nil, status.Errorf(sttInfo.Code(), "%v", sttInfo.Err())
	}

	slog.Info(logger.GCodeSuite(utils.UserLogin, codes.OK))
	return resp, nil
}
