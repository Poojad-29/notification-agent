package middleware

import (
	"sync"
	"time"
)

type RateLimiter struct {
	requests map[string]time.Time
	mu       sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]time.Time),
	}
}

func (r *RateLimiter) Allow(clientID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	lastRequest, exists := r.requests[clientID]

	if exists && time.Since(lastRequest) < time.Second {
		return false
	}

	r.requests[clientID] = time.Now()
	return true
}
