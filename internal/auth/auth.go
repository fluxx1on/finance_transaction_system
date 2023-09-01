package auth

import (
	"context"
	"strings"

	pass "github.com/fluxx1on/finance_transaction_system/internal/auth/password"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Key string

const UserKey Key = "user"

func validateByUsername(username string) bool {

	if len(username) > 24 {
		return false
	}

	return true
}

func validateByPassword(password string) bool {

	if len(password) > 18 || len(password) < 8 {
		return false
	}

	return true
}

func Register(db *repo.CreditDB, username, password, password2 string) (*Token, *Token, error) {
	var (
		UVal bool = validateByUsername(username)
		PVal bool = validateByPassword(password)
	)

	if UVal && PVal && strings.Compare(password, password2) == 0 {
		user, err := db.CreateUser(username, password)
		if err != nil {
			return nil, nil, status.Error(codes.AlreadyExists, err.Error())
		}

		rtoken, err := GenerateRefreshToken(user)
		if err != nil {
			return nil, nil, err
		}

		atoken, err := GenerateAccessToken(user)
		if err != nil {
			return nil, nil, err
		}

		return rtoken, atoken, nil
	}

	return nil, nil, status.Error(codes.InvalidArgument, "Unexpected intervention")
}

// First is a refresh token
// Second is an access token
func Authentication(db *repo.CreditDB, username, password string) (*Token, *Token, error) {
	person, err := db.GetUser(username)
	if err != nil && pass.Compare(password, person.Password) {
		return nil, nil, status.Errorf(codes.Unauthenticated, "Wrong username or password")
	}

	rtoken, err := GenerateRefreshToken(person)
	if err != nil {
		return nil, nil, err
	}

	atoken, err := GenerateAccessToken(person)
	if err != nil {
		return nil, nil, err
	}

	return rtoken, atoken, nil
}

func SetUserIntoContext(ctx context.Context, user *repo.Person) context.Context {
	return context.WithValue(ctx, UserKey, user)
}

func GetUserFromContext(ctx context.Context) (*repo.Person, error) {

	if user := ctx.Value(UserKey); user != nil {
		if user, is := user.(*repo.Person); is {
			return user, nil
		}
	}

	return nil, status.Error(codes.Unauthenticated, "")
}
