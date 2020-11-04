package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewResourceCards(game *Game) ResourceCards {
	var resourceCards ResourceCards

	resourceTypes := []datamodel.ResourceType{
		datamodel.RT_BRICK,
		datamodel.RT_GRAIN,
		datamodel.RT_SHEEP,
		datamodel.RT_STONE,
		datamodel.RT_WOOD,
	}
	for _, resourceType := range resourceTypes {
		for i := 1; i <= 19; i++ {
			resourceCard := NewResourceCard(game, resourceType)
			resourceCards.append(resourceCard)
		}
	}

	return resourceCards
}

type ResourceCards []*ResourceCard

func (rc ResourceCards) append(resourceCard *ResourceCard) {
	rc = append(rc, resourceCard)
}

func (rc ResourceCards) remove(resourceCard *ResourceCard) {
	for idx := range rc {
		if rc[idx] == resourceCard {
			rc = append(rc[:idx], rc[idx+1:]...)
			return
		}
	}
}

func (rc ResourceCards) SetGame(game *Game) {
	for _, resourceCard := range rc {
		resourceCard.SetGame(game)
	}
}

func (rc ResourceCards) SetPlayer(player *Player) {
	for _, resourceCard := range rc {
		resourceCard.SetPlayer(player)
	}
}
