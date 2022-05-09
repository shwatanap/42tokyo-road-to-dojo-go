package middleware

import (
	"net/http"

	"42tokyo-road-to-dojo-go/pkg/core/context"
)

func Context(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewContext(r)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
