package middleware

import (
	"context"
	"net/http"

	customError "42tokyo-road-to-dojo-go/pkg/core/error"
	"42tokyo-road-to-dojo-go/pkg/core/logger"
)

type TToken string

var Token TToken

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// TODO: 存在するかチェック必要？
		token := request.Header.Get("X-Token")
		if token == "" {
			writer.WriteHeader(http.StatusBadRequest)
			logger.ErrorLogging("GET user/get: x-token not found error", customError.ErrTokenNotFound, request)
			return
		}
		ctx = context.WithValue(ctx, Token, token)

		next.ServeHTTP(writer, request.WithContext(ctx))
	}
}
