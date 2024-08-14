package server

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	rate      int
	burst     int
	tokens    int
	lastCheck time.Time
	mu        sync.Mutex
}

func NewRateLimiter(rate int, burst int) *RateLimiter {
	return &RateLimiter{
		rate:      rate,
		burst:     burst,
		tokens:    burst,
		lastCheck: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastCheck).Seconds()
	rl.tokens += int(elapsed * float64(rl.rate))
	if rl.tokens > rl.burst {
		rl.tokens = rl.burst
	}
	rl.lastCheck = now

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

func (rl *RateLimiter) Middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !rl.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}
