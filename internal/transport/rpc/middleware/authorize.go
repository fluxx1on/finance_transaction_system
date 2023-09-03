package middleware

import (
	"context"

	"github.com/fluxx1on/finance_transaction_system/internal/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var _ grpc.UnaryServerInterceptor = AuthInterceptor

// TODO : Divide services (authorized/unauthorized)
func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, _ := metadata.FromIncomingContext(ctx)

	token := md["Authorization"]
	if token != nil {
		tokenString := token[0]

		person, err := auth.CheckAccessToken(tokenString)
		if err != nil {
			tokenString = token[1]
			person, err = auth.CheckRefreshToken(tokenString)
			if err != nil {
				return nil, err
			}

			user, err := auth.GenerateAccessToken(person)
			if err != nil {
				return nil, err
			}

			md := metadata.Pairs("Authorization", user.Token)

			metadata.NewOutgoingContext(ctx, md)
		}

		newCtx := auth.SetUserIntoContext(ctx, person)

		return handler(newCtx, req)
	}

	return handler(ctx, req)
}
