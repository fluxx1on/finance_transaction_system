package auth

import (
	"strings"

	pass "github.com/fluxx1on/finance_transaction_system/internal/auth/password"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateByUsername(username string) bool {

	// TODO

	return true
}

func validateByPassword(password string) bool {

	// TODO

	return true
}

func Register(username, password, password2 string) bool {
	var (
		UVal bool = validateByUsername(username)
		PVal bool = validateByPassword(password)
	)

	return UVal && PVal && strings.Compare(password, password2) == 0
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
