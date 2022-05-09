package repository

import (
	"context"

	"42tokyo-road-to-dojo-go/pkg/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, name string) (*entity.User, error)
	Get(ctx context.Context, token string) (*entity.User, error)
	Update(ctx context.Context, name, token string) (*entity.User, error)
}
