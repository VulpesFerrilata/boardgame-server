package repository

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/db"
)

type GameRepository interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
	Insert(ctx context.Context, game *model.Game) error
}

func NewGameRepository(db *db.DbContext) GameRepository {
	return &gameRepository{
		db: db,
	}
}

type gameRepository struct {
	db *db.DbContext
}

func (gr gameRepository) GetById(ctx context.Context, id uint) (*model.Game, error) {
	game := new(model.Game)
	return game, gr.db.GetDB(ctx).First(game, id).Error
}

func (gr gameRepository) Insert(ctx context.Context, game *model.Game) error {
	return gr.db.GetDB(ctx).Create(game).Error
}
