package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fluxx1on/finance_transaction_system/internal/database"
	"golang.org/x/exp/slog"
)

const secret = "hardcoded-string"

func CreateToken(user *database.Person) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		slog.Error("Unsigned token", err)
		return "", err
	}

	return tokenString, nil
}
