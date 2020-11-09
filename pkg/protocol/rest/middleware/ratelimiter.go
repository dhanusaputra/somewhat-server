package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/httprate"
)

// AddRateLimiter ...
func AddRateLimiter(next http.Handler) http.Handler {
	return httprate.Limit(10, 1*time.Minute, httprate.KeyByIP)(next)
}
