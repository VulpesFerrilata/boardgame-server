package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type AchievementService interface {
	GetAchievementRepository() repository.SafeAchievementRepository
	Save(ctx context.Context, achievement *model.Achievement) error
}

type achievementService struct {
	achievementRepository repository.AchievementRepository
}

func (as achievementService) GetAchievementRepository() repository.SafeAchievementRepository {
	return as.achievementRepository
}

func (as achievementService) validate(ctx context.Context, achievement *model.Achievement) error {
	//TODO: validate dice
	return nil
}

func (as achievementService) Save(ctx context.Context, achievement *model.Achievement) error {
	if err := as.validate(ctx, achievement); err != nil {
		return err
	}

	return as.achievementRepository.InsertOrUpdate(ctx, achievement)
}
