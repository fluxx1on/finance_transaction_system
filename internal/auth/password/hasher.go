package password

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
)

const (
	cost = 10

	// Global

	// It's important that password size limited by 18 letters
	// Cause bcrypt.GenerateFromPassword gets []byte with max len 72
	// So to provide utf-8 letters to password we need to
	// store max 18 letters by usually
	MaxPasswordSize = 18
)

func GetHash(unhashedPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(unhashedPassword), cost)
	if err != nil {
		slog.Debug("Can't generate hashed_password", err)
		return "", err
	}

	return string(hashedPassword), nil
}

func Compare(unhashed, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(unhashed))

	return err == nil
}
