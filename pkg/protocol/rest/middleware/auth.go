package middleware

import (
	"net/http"

	"github.com/dhanusaputra/somewhat-server/util/auth"
	"github.com/dhanusaputra/somewhat-server/util/envutil"
)

var (
	defaultAuthMethodBlacklists = map[string]bool{
		"GET": true,
	}
	defaultAuthRequestURIBlacklists = map[string]bool{
		"/v1/login": true,
	}
)

const (
	defaultAuthEnable = true
)

// AddAuth ...
func AddAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !envutil.GetEnvAsBool("ENABLE_AUTH", defaultAuthEnable) || envutil.GetEnvAsMapBool("AUTH_METHOD_BLACKLISTS", defaultAuthMethodBlacklists, ",")[r.Method] || envutil.GetEnvAsMapBool("AUTH_REQUESTURI_BLACKLISTS", defaultAuthRequestURIBlacklists, ",")[r.RequestURI] {
			h.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("no authorization found in request"))
			return
		}
		_, err := auth.ValidateJWT(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}
		h.ServeHTTP(w, r)
	})
}
