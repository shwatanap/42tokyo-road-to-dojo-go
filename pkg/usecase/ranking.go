package usecase

import (
	"context"

	"42tokyo-road-to-dojo-go/pkg/domain/entity"
	"42tokyo-road-to-dojo-go/pkg/domain/repository"
)

type RankingUsecase interface {
	List(ctx context.Context, start int) ([]*entity.User, error)
}

type rankingUsecase struct {
	userRepo repository.UserRepository
}

func NewRankingUsecase(ur repository.UserRepository) RankingUsecase {
	rankingUsecase := rankingUsecase{userRepo: ur}
	return &rankingUsecase
}

func (ru *rankingUsecase) List(ctx context.Context, start int) ([]*entity.User, error) {
	return ru.userRepo.LimitGet(ctx, start)
}
