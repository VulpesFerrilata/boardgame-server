package response

import "github.com/VulpesFerrilata/catan/internal/domain/model"

func NewPlayerResponse(player *model.Player) *PlayerResponse {
	playerResponse := new(PlayerResponse)
	playerResponse.ID = int(player.ID)
	playerResponse.GameID = int(player.GameID)
	playerResponse.Color = player.Color
	playerResponse.TurnOrder = player.TurnOrder
	return playerResponse
}

type PlayerResponse struct {
	ID        int           `json:"id"`
	GameID    int           `json:"gameId"`
	UserID    int           `json:"userId"`
	User      *UserResponse `json:"user"`
	Color     string        `json:"color"`
	TurnOrder int           `json:"turnOrder"`
}
