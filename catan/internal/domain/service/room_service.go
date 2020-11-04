package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type RoomService interface {
	GetRoomById(ctx context.Context, id uint) (*model.Room, error)
	FindRoomsByStatus(ctx context.Context, status datamodel.GameStatus) ([]*model.Room, error)
}
