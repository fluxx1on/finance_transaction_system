package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	secret        = "hardcoded-string"
	accessExpiry  = time.Hour
	refreshExpiry = 30 * 24 * time.Hour
)

func Secret() []byte { return []byte(secret) }

func CheckRefreshToken(tokenString string) (*repo.Person, error) {
	return checkToken(tokenString)
}

func CheckAccessToken(tokenString string) (*repo.Person, error) {
	return checkToken(tokenString)
}

func checkToken(tokenString string) (*repo.Person, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret(), nil
	})

	if err != nil || !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "Invalid or expired token")
	}

	var person repo.Person
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		person = repo.Person{
			ID:       claims["id"].(uint64),
			Username: claims["username"].(string),
		}
	}

	return &person, nil
}

func GenerateRefreshToken(user *repo.Person) (*Token, error) {

	expiredAt := time.Now().Add(refreshExpiry)

	return genToken(user, expiredAt)
}

func GenerateAccessToken(user *repo.Person) (*Token, error) {

	expiredAt := time.Now().Add(accessExpiry)

	return genToken(user, expiredAt)
}

func genToken(user *repo.Person, expiredAt time.Time) (*Token, error) {

	expire := expiredAt.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"expiryAt": expire,
	})

	tokenString, err := token.SignedString(Secret())
	if err != nil {
		slog.Error("Unsigned token", err)
		return nil, status.Errorf(codes.Unavailable, "Service unavailable")
	}

	tokenized := &Token{
		UserID:    user.ID,
		Token:     tokenString,
		ExpiredAt: expiredAt,
	}

	return tokenized, nil
}

type Token struct {
	UserID    uint64    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}
