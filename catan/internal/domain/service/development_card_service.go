package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type DevelopmentCardService interface {
	GetDevelopmentCardRepository() repository.SafeDevelopmentCardRepository
	Save(ctx context.Context, developmentCard *model.DevelopmentCard) error
}

type developmentCardService struct {
	developmentCardRepository repository.DevelopmentCardRepository
}

func (dcs developmentCardService) GetDevelopmentCardRepository() repository.SafeDevelopmentCardRepository {
	return dcs.developmentCardRepository
}

func (dcs developmentCardService) validate(ctx context.Context, developmentCard *model.DevelopmentCard) error {
	//TODO: validate dice
	return nil
}

func (dcs developmentCardService) Save(ctx context.Context, developmentCard *model.DevelopmentCard) error {
	if err := dcs.validate(ctx, developmentCard); err != nil {
		return err
	}

	return dcs.developmentCardRepository.InsertOrUpdate(ctx, developmentCard)
}
