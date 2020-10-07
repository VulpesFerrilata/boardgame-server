package service

import (
	"context"
	"math/rand"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/repository"
)

type DiceService interface {
	GetDiceRepository() repository.DiceRepository
	Init(ctx context.Context, gameId uint) ([]*model.Dice, error)
	Roll(ctx context.Context, gameId uint) ([]*model.Dice, error)
	ResetState(ctx context.Context, gameId uint) ([]*model.Dice, error)
}

func NewDiceService(diceRepository repository.DiceRepository) DiceService {
	return &diceService{
		diceRepository: diceRepository,
	}
}

type diceService struct {
	diceRepository repository.DiceRepository
}

func (ds diceService) GetDiceRepository() repository.DiceRepository {
	return ds.diceRepository
}

func (ds diceService) Init(ctx context.Context, gameId uint) ([]*model.Dice, error) {
	dices := make([]*model.Dice, 2)

	for _, dice := range dices {
		dice.GameID = gameId
		dice.Number = 1
		dice.IsRolled = false
	}

	return dices, ds.diceRepository.Insert(ctx, dices...)
}

func (ds diceService) Roll(ctx context.Context, gameId uint) ([]*model.Dice, error) {
	dices, err := ds.diceRepository.FindByGameId(ctx, gameId)
	if err != nil {
		return nil, err
	}

	for _, dice := range dices {
		dice.Number = rand.Intn(5) + 1
		dice.IsRolled = true
	}
	if err := ds.diceRepository.Save(ctx, dices...); err != nil {
		return nil, err
	}
	return dices, nil
}

func (ds diceService) ResetState(ctx context.Context, gameId uint) ([]*model.Dice, error) {
	dices, err := ds.diceRepository.FindByGameId(ctx, gameId)
	if err != nil {
		return nil, err
	}

	for _, dice := range dices {
		dice.IsRolled = false
	}
	if err := ds.diceRepository.Save(ctx, dices...); err != nil {
		return nil, err
	}
	return dices, nil
}
