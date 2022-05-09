package middleware

import (
	"net/http"
)

func Layers(handler http.HandlerFunc) http.HandlerFunc {
	return Context(
		Logger(
			handler,
		),
	)
}

func AuthLayers(handler http.HandlerFunc) http.HandlerFunc {
	return Context(
		Logger(
			Authenticate(
				handler,
			),
		),
	)
}
