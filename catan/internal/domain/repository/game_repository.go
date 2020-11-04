package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeGameRepository interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
}

type GameRepository interface {
	SafeGameRepository
	InsertOrUpdate(ctx context.Context, game *model.Game) error
	Delete(ctx context.Context, game *model.Game) error
}
