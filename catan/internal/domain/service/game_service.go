package service

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/repository"
)

type GameService interface {
	IsExists(ctx context.Context, game *model.Game) (bool, error)
	GetGameRepository() repository.ReadOnlyGameRepository
	New(ctx context.Context) (*model.Game, error)
}

func NewGameService(gameRepository repository.GameRepository) GameService {
	return &gameService{
		gameRepository: gameRepository,
	}
}

type gameService struct {
	gameRepository repository.GameRepository
}

func (gs gameService) GetGameRepository() repository.ReadOnlyGameRepository {
	return gs.gameRepository
}

func (gs gameService) New(ctx context.Context) (*model.Game, error) {
	game := new(model.Game)
	game.Status = model.GS_WAITED
	return game, gs.gameRepository.Insert(ctx, game)
}
