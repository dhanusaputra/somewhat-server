package middleware

import (
	"net/http"

	"github.com/dhanusaputra/somewhat-server/util/auth"
	"github.com/dhanusaputra/somewhat-server/util/envutil"
)

var (
	defaultAuthMethodBlacklist = map[string]bool{
		"GET": true,
	}
	defaultAuthRequestURIBlacklist = map[string]bool{
		"/v1/login": true,
	}
)

const (
	defaultAuthEnable = true
)

// AddAuth ...
func AddAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !envutil.GetEnvAsBool("ENABLE_AUTH", defaultAuthEnable) || defaultAuthMethodBlacklist[r.Method] || defaultAuthRequestURIBlacklist[r.RequestURI] {
			h.ServeHTTP(w, r)
			return
		}
		authHeader, ok := r.Header["Authorization"]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("no authorization found in request"))
			return
		}
		_, err := auth.ValidateJWT(authHeader[0])
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}
		h.ServeHTTP(w, r)
	})
}
