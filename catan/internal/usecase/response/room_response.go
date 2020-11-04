package response

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewRoomResponse(room *model.Room) *RoomResponse {
	roomResponse := new(RoomResponse)
	roomResponse.ID = int(room.ID)
	roomResponse.Status = room.Status
	roomResponse.PlayerCount = room.PlayerCount
	return roomResponse
}

type RoomResponse struct {
	ID          int                  `json:"id"`
	Status      datamodel.GameStatus `json:"status"`
	PlayerCount int                  `json:"playerCount"`
}
