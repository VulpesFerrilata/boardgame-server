package model

import (
	"github.com/VulpesFerrilata/catan/internal/domain/datamodel"
)

func NewGame() *Game {
	game := new(Game)
	game.Game = new(datamodel.Game)
	game.Status = datamodel.GS_WAITING
	return game
}

type Game struct {
	*datamodel.Game
	players          Players
	dices            Dices
	achievements     Achievements
	resourceCards    ResourceCards
	developmentCards DevelopmentCards
	robber           *Robber
	isRemoved        bool
}

func (g Game) GetPlayers() Players {
	return g.players
}

func (g *Game) GetDices() Dices {
	return g.dices
}

func (g *Game) GetAchievements() Achievements {
	return g.achievements
}

func (g *Game) GetRobber() *Robber {
	return g.robber
}

func (g Game) IsRemoved() bool {
	return g.isRemoved
}

func (g *Game) Init() {
	g.Status = datamodel.GS_STARTED

	NewDices(g)
	NewAchievements(g)
	NewResourceCards(g)
	NewDevelopmentCards(g)
}
