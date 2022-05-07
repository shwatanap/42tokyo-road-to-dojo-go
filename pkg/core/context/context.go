package context

import (
	"context"
	"net/http"

	ua "github.com/mileusna/useragent"
)

type OS struct{}

var OsKey OS

func NewContext(r *http.Request) context.Context {
	ctx := r.Context()
	userAgent := r.UserAgent()
	ua := ua.Parse(userAgent)

	return context.WithValue(ctx, OsKey, ua.OS)
}
