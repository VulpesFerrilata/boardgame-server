package repository

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type SafeResourceCardRepository interface {
	FindByGameIdByPlayerId(ctx context.Context, gameId uint, playerId *uint) (model.ResourceCards, error)
}

type ResourceCardRepository interface {
	SafeResourceCardRepository
	InsertOrUpdate(ctx context.Context, resourceCard *model.ResourceCard) error
}
