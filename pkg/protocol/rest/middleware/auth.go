package middleware

import (
	"net/http"

	"github.com/dhanusaputra/somewhat-server/util/authutil"
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
		if !envutil.GetEnvAsBool("ENABLE_AUTH", defaultAuthEnable) || envutil.GetEnvAsMapBool("AUTH_METHOD_BLACKLIST", defaultAuthMethodBlacklist, ",")[r.Method] || envutil.GetEnvAsMapBool("AUTH_REQUESTURI_BLACKLIST", defaultAuthRequestURIBlacklist, ",")[r.RequestURI] {
			h.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("no authorization found in request"))
			return
		}
		_, _, err := authutil.ValidateJWT(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}
		h.ServeHTTP(w, r)
	})
}
