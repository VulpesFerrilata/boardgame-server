package repository

import (
	"context"

	"github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/db"
)

type PlayerRepository interface {
}

func NewPlayerRepository(db *db.DbContext) PlayerRepository {
	return &playerRepository{
		db: db,
	}
}

type playerRepository struct {
	db *db.DbContext
}

func (pr playerRepository) GetByUserId(ctx context.Context, userId uint) (*model.Player, error) {
	player := new(model.Player)
	return player, pr.db.GetDB(ctx).First(player).Error
}

func (pr playerRepository) Insert(ctx context.Context, players ...*model.Player) error {
	return pr.db.GetDB(ctx).Create(players).Error
}

func (pr playerRepository) Save(ctx context.Context, players ...*model.Player) error {
	return pr.db.GetDB(ctx).Save(players).Error
}
