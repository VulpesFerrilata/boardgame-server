package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeRobberRepository interface {
	GetByGameId(ctx context.Context, gameId uint) (*model.Robber, error)
}

type RobberRepository interface {
	SafeRobberRepository
	InsertOrUpdate(ctx context.Context, robber *model.Robber) error
}
