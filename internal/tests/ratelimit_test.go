package tests

import (
	"testing"

	"notification-agent/internal/middleware"
)

func TestRateLimiter(t *testing.T) {
	limiter := middleware.NewRateLimiter()

	if !limiter.Allow("user1") {
		t.Error("first request should be allowed")
	}

	if limiter.Allow("user1") {
		t.Error("second request should be blocked")
	}
}
