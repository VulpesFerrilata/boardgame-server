package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewDevelopmentCard(game *Game, developmentType datamodel.DevelopmentType) *DevelopmentCard {
	developmentCard := new(DevelopmentCard)
	developmentCard.Type = developmentType
	developmentCard.SetGame(game)
	return developmentCard
}

type DevelopmentCard struct {
	*datamodel.DevelopmentCard
	game   *Game
	player *Player
}

func (dc *DevelopmentCard) SetGame(game *Game) {
	dc.GameID = game.ID
	dc.game = game

	dc.game.developmentCards.append(dc)
}

func (dc *DevelopmentCard) SetPlayer(player *Player) {
	if dc.player != nil {
		dc.PlayerID = nil
		dc.player.developmentCards.remove(dc)
	}

	dc.player = player
	if player != nil {
		dc.PlayerID = &player.ID
		player.developmentCards.append(dc)
	}
}
