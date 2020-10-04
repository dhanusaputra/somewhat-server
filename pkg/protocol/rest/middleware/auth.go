package middleware

import (
	"net/http"

	"github.com/dhanusaputra/somewhat-server/pkg/env"
	"github.com/dhanusaputra/somewhat-server/util/authutil"
)

// AddAuth ...
func AddAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := env.AuthMiddlewareConf
		if !conf.AuthEnable || conf.AuthMethodBlacklist[r.Method] || conf.AuthRequestURIBlacklist[r.RequestURI] {
			next.ServeHTTP(w, r)
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
		next.ServeHTTP(w, r)
	})
}
