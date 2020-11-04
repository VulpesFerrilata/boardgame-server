package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type ResourceCardService interface {
	GetResourceCardRepository() repository.SafeResourceCardRepository
	Save(ctx context.Context, resourceCard *model.ResourceCard) error
}

type resourceCardService struct {
	resourceCardRepository repository.ResourceCardRepository
}

func (rcs resourceCardService) GetResourceCardRepository() repository.SafeResourceCardRepository {
	return rcs.resourceCardRepository
}

func (rcs resourceCardService) validate(ctx context.Context, resourceCard *model.ResourceCard) error {
	//TODO: validate resource card
	return nil
}

func (rcs resourceCardService) Save(ctx context.Context, resourceCard *model.ResourceCard) error {
	if err := rcs.validate(ctx, resourceCard); err != nil {
		return err
	}

	return rcs.resourceCardRepository.InsertOrUpdate(ctx, resourceCard)
}
