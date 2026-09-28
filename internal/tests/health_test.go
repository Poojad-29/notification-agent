package tests

import (
	"testing"

	"notification-agent/internal/health"
)

func TestHealthCheck(t *testing.T) {
	status := health.Check()

	if status != "UP" {
		t.Error("health check failed")
	}
}
