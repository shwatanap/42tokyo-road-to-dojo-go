//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"

	"42tokyo-road-to-dojo-go/pkg/infra/repository"
	"42tokyo-road-to-dojo-go/pkg/presen/handler"
	"42tokyo-road-to-dojo-go/pkg/usecase"
)

func initUserHandler(driver *sql.DB) handler.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	)
	return nil
}
