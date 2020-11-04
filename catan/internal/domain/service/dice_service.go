package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type DiceService interface {
	GetDiceRepository() repository.SafeDiceRepository
	Save(ctx context.Context, dice *model.Dice) error
}

func NewDiceService(diceRepository repository.DiceRepository) DiceService {
	return &diceService{
		diceRepository: diceRepository,
	}
}

type diceService struct {
	diceRepository repository.DiceRepository
}

func (ds diceService) GetDiceRepository() repository.SafeDiceRepository {
	return ds.diceRepository
}

func (ds diceService) validate(ctx context.Context, dice *model.Dice) error {
	//TODO: validate dice
	return nil
}

func (ds diceService) Save(ctx context.Context, dice *model.Dice) error {
	if err := ds.validate(ctx, dice); err != nil {
		return err
	}

	return ds.diceRepository.InsertOrUpdate(ctx, dice)
}
