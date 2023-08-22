package middleware

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/auth"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var _ grpc.UnaryServerInterceptor = AuthInterceptor

type Key string

const UserKey Key = "user"

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, _ := metadata.FromIncomingContext(ctx)
	tokenString := md["Authorization"][0]

	person, err := auth.CheckAccessToken(tokenString)
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, UserKey, person)

	return handler(newCtx, req)
}

func GetUserFromContext(ctx context.Context) (*repo.Person, error) {

	if user := ctx.Value(UserKey); user != nil {
		if user, is := user.(*repo.Person); is {
			return user, nil
		}
	}

	return nil, status.Error(codes.Unauthenticated, "")
}
