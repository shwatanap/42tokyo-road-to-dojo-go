package repository

import (
	"context"

	"42tokyo-road-to-dojo-go/pkg/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, name string) (*entity.User, error)
}
