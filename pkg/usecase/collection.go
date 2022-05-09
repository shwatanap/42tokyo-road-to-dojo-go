package usecase

import (
	"context"

	"42tokyo-road-to-dojo-go/pkg/domain/entity"
	"42tokyo-road-to-dojo-go/pkg/domain/repository"
)

type CollectionUsecase interface {
	List(ctx context.Context, token string) ([]*entity.UserChara, error)
}

type collectionUsecase struct {
	userRepo      repository.UserRepository
	userCharaRepo repository.UserCharaRepository
}

func NewCollectionUsecase(ur repository.UserRepository, ucr repository.UserCharaRepository) CollectionUsecase {
	collectionUsecase := collectionUsecase{userRepo: ur, userCharaRepo: ucr}
	return &collectionUsecase
}

func (cu *collectionUsecase) List(ctx context.Context, token string) ([]*entity.UserChara, error) {
	user, err := cu.userRepo.Get(ctx, token)
	if err != nil {
		return nil, err
	}

	userChara, err := cu.userCharaRepo.List(ctx, *user)
	if err != nil {
		return nil, err
	}
	return userChara, nil
}
