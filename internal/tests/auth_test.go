package tests

import (
	"os"
	"testing"

	"notification-agent/internal/middleware"
)

func TestValidateToken(t *testing.T) {
	os.Setenv("GRPC_TOKEN", "my-secret-token")

	err := middleware.ValidateToken("my-secret-token")

	if err != nil {
		t.Errorf("expected valid token, got error: %v", err)
	}
}
