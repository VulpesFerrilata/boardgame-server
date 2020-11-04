package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type GameService interface {
	GetGameRepository() repository.SafeGameRepository
	Save(ctx context.Context, game *model.Game) error
}

type gameService struct {
	gameRepository repository.GameRepository
}

func (gs gameService) GetGameRepository() repository.SafeGameRepository {
	return gs.gameRepository
}

func (gs gameService) validate(ctx context.Context, game *model.Game) error {
	//TODO: validate game
	return nil
}

func (gs gameService) Save(ctx context.Context, game *model.Game) error {
	if game.IsRemoved() {
		return gs.gameRepository.Delete(ctx, game)
	}

	if err := gs.validate(ctx, game); err != nil {
		return err
	}

	return gs.gameRepository.InsertOrUpdate(ctx, game)
}
