package response

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

func NewGameResponse(game *model.Game) *GameResponse {
	gameResponse := new(GameResponse)
	gameResponse.ID = int(game.ID)
	gameResponse.Host = int(game.Host)
	gameResponse.PlayerInTurn = int(game.PlayerInTurn)
	gameResponse.Status = game.Status

	players := game.GetPlayers()
	for _, player := range players {
		playerResponse := NewPlayerResponse(player)
		gameResponse.Players = append(gameResponse.Players, playerResponse)
	}

	return gameResponse
}

type GameResponse struct {
	ID           int                  `json:"id"`
	Host         int                  `json:"host"`
	PlayerInTurn int                  `json:"playerInTurn"`
	Status       datamodel.GameStatus `json:"status"`
	Players      []*PlayerResponse    `json:"players"`
}
