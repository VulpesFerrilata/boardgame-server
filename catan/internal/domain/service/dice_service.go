package service

import (
	"context"
	"math/rand"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/repository"
)

type DiceService interface {
	Roll(ctx context.Context, gameId uint) ([]*model.Dice, error)
}

func NewDiceService(diceRepository repository.DiceRepository) DiceService {
	return &diceService{
		diceRepository: diceRepository,
	}
}

type diceService struct {
	diceRepository repository.DiceRepository
}

func (ds diceService) Roll(ctx context.Context, gameId uint) ([]*model.Dice, error) {
	dices, err := ds.diceRepository.FindByGameId(ctx, gameId)
	if err != nil {
		return nil, err
	}

	for _, dice := range dices {
		dice.Number = rand.Intn(5) + 1
	}
	if err := ds.diceRepository.Save(ctx, dices...); err != nil {
		return nil, err
	}
	return dices, nil
}
