package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type RobberService interface {
	GetRobberRepository() repository.SafeRobberRepository
	Save(ctx context.Context, robber *model.Robber) error
}

type robberService struct {
	robberRepository repository.RobberRepository
}

func (rs robberService) GetRobberRepository() repository.SafeRobberRepository {
	return rs.robberRepository
}

func (rs robberService) validate(ctx context.Context, robber *model.Robber) error {
	//TODO: validate dice
	return nil
}

func (rs robberService) Save(ctx context.Context, robber *model.Robber) error {
	if err := rs.validate(ctx, robber); err != nil {
		return err
	}

	return rs.robberRepository.InsertOrUpdate(ctx, robber)
}
