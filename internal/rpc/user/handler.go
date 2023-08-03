package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/fluxx1on/finance_transaction_system/internal/database"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/pb"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
)

const cost = 10

type UserFetcher struct {
	db *database.CreditDB
}

func NewUserFetcher(db *database.CreditDB) *UserFetcher {

	return &UserFetcher{
		db: db,
	}
}

func (f *UserFetcher) FetchCreate(ctx context.Context, req *pb.SignUpRequest) (
	*pb.SignUpResponse, error,
) {

	// TODO
	return nil, nil
}

func (f *UserFetcher) authentication(username, password string) (*database.Token, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		slog.Debug("Can't generate hashed_password", err)
		return nil, fmt.Errorf("invalid data")
	}

	person, err := f.db.UserByAuth(username)
	if err != nil || person.Password != string(hashed_password) {
		return nil, fmt.Errorf("wrong username or password")
	}

	token, err := f.db.Token(person.ID)
	if err != nil {
		return nil, fmt.Errorf("unsigned token")
	}

	return token, nil
}

func (f *UserFetcher) FetchLogin(ctx context.Context, req *pb.SignInRequest) (
	*pb.SignInResponse, error,
) {
	var errMsg string

	token, err := f.authentication(req.Username, req.Password)
	if err != nil {
		errMsg = strings.Title(err.Error())
	}

	resp := &pb.SignInResponse{
		UserToken:    token.Token,
		ErrorMessage: errMsg,
	}
	return resp, err
}
