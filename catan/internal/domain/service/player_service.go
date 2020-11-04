package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/catan/internal/domain/repository"
)

type PlayerService interface {
	GetPlayerRepository() repository.SafePlayerRepository
	Save(ctx context.Context, player *model.Player) error
}

type playerService struct {
	playerRepository repository.PlayerRepository
}

func (ps playerService) GetPlayerRepository() repository.SafePlayerRepository {
	return ps.playerRepository
}

func (ps playerService) validate(ctx context.Context, player *model.Player) error {
	//TODO: validate player
	return nil
}

func (ps playerService) Save(ctx context.Context, player *model.Player) error {
	if player.IsRemoved() {
		return ps.playerRepository.Delete(ctx, player)
	}

	if err := ps.validate(ctx, player); err != nil {
		return err
	}

	return ps.playerRepository.InsertOrUpdate(ctx, player)
}
