package request

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type RoomRequest struct {
	ID     int                  `json:"id"`
	Status datamodel.GameStatus `json:"status"`
}
