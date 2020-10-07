package service

import "github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/repository"

type PlayerService interface {
}

func NewPlayerService(playerRepository repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepository: playerRepository,
	}
}

type playerService struct {
	playerRepository repository.PlayerRepository
}
