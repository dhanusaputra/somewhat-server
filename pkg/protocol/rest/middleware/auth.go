package middleware

import (
	"net/http"

	"github.com/dhanusaputra/somewhat-server/util/auth"
	"google.golang.org/grpc/metadata"
)

// AddAuth ...
func AddAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("no metadata found in request"))
		}
		_, err := auth.ValidateJWT(md.Get("authorization")[0])
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		}
	})
}
