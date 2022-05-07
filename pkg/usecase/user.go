package usecase

import (
	"context"

	"42tokyo-road-to-dojo-go/pkg/domain/entity"
	"42tokyo-road-to-dojo-go/pkg/domain/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, name string) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo: ur}
	return &userUsecase
}

func (uu *userUsecase) Create(ctx context.Context, name string) (user *entity.User, err error) {
	user, err = uu.userRepo.Create(ctx, name)
	return
}
