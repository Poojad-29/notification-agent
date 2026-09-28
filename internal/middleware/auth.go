package middleware

import (
	"errors"
	"os"
)

func ValidateToken(token string) error {
	expectedToken := os.Getenv("GRPC_TOKEN")

	if token == "" {
		return errors.New("token is missing")
	}

	if token != expectedToken {
		return errors.New("invalid token")
	}

	return nil
}
