package middleware

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var _ grpc.UnaryServerInterceptor = AuthInterceptor

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

	newCtx := auth.SetUserIntoContext(ctx, person)

	return handler(newCtx, req)
}
