package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewDevelopmentCards(game *Game) DevelopmentCards {
	var developmentCards DevelopmentCards

	for i := 1; i <= 14; i++ {
		knightCard := NewDevelopmentCard(game, datamodel.DT_KNIGHT)
		developmentCards.append(knightCard)
	}

	developmentTypes := []datamodel.DevelopmentType{
		datamodel.DT_ROAD_BUILDING,
		datamodel.DT_YEAR_OF_PLENTY,
		datamodel.DT_MONOPOLY,
	}
	for _, developmentType := range developmentTypes {
		for i := 1; i <= 2; i++ {
			developmentCard := NewDevelopmentCard(game, developmentType)
			developmentCards.append(developmentCard)
		}
	}

	for i := 1; i <= 5; i++ {
		victoryPointCard := NewDevelopmentCard(game, datamodel.DT_VICTORY_POINTS)
		developmentCards.append(victoryPointCard)
	}

	return developmentCards
}

type DevelopmentCards []*DevelopmentCard

func (dc DevelopmentCards) append(developmentCard *DevelopmentCard) {
	dc = append(dc, developmentCard)
}

func (dc DevelopmentCards) remove(developmentCard *DevelopmentCard) {
	for idx := range dc {
		if dc[idx] == developmentCard {
			dc = append(dc[:idx], dc[idx+1:]...)
			return
		}
	}
}

func (dc DevelopmentCards) SetGame(game *Game) {
	for _, developmentCard := range dc {
		developmentCard.SetGame(game)
	}
}

func (dc DevelopmentCards) SetPlayer(player *Player) {
	for _, developmentCard := range dc {
		developmentCard.SetPlayer(player)
	}
}
