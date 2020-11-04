package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewResourceCard(game *Game, resourceType datamodel.ResourceType) *ResourceCard {
	resourceCard := new(ResourceCard)
	resourceCard.Type = resourceType
	resourceCard.SetGame(game)
	return resourceCard
}

type ResourceCard struct {
	*datamodel.ResourceCard
	game   *Game
	player *Player
}

func (rc *ResourceCard) SetGame(game *Game) {
	rc.GameID = game.ID
	rc.game = game

	rc.game.resourceCards.append(rc)
}

func (rc *ResourceCard) SetPlayer(player *Player) {
	if rc.player != nil {
		rc.PlayerID = nil
		rc.player.resourceCards.remove(rc)
	}

	rc.player = player
	if player != nil {
		rc.PlayerID = &player.ID
		player.resourceCards.append(rc)
	}
}
