package repository

import (
	"context"

	"42tokyo-road-to-dojo-go/pkg/domain/entity"
)

type UserCharaRepository interface {
	List(ctx context.Context, ue entity.User) ([]*entity.UserChara, error)
	Store(ctx context.Context, ue entity.User, ces []*entity.Chara) error
}
