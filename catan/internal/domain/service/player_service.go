package service

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/repository"
)

type PlayerService interface {
	IsExists(ctx context.Context, player *model.Player) (bool, error)
	GetPlayerRepository() repository.ReadOnlyPlayerRepository
	Create(ctx context.Context, player *model.Player) error
}

func NewPlayerService(playerRepository repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepository: playerRepository,
	}
}

type playerService struct {
	playerRepository repository.PlayerRepository
}
