package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewPlayer(game *Game) *Player {
	player := new(Player)
	player.SetGame(game)
	return player
}

type Player struct {
	*datamodel.Player
	game             *Game
	user             *User
	achievements     Achievements
	resourceCards    ResourceCards
	developmentCards DevelopmentCards
	isRemoved        bool
}

func (p *Player) SetGame(game *Game) {
	p.GameID = game.ID
	p.game = game
	game.players.append(p)
	p.achievements.SetGame(game)
	p.resourceCards.SetGame(game)
	p.developmentCards.SetGame(game)
}

func (p *Player) SetUser(user *User) {
	p.UserID = user.ID
	p.user = user
}

func (p Player) GetUser() *User {
	return p.user
}

func (p Player) IsRemoved() bool {
	return p.isRemoved
}

func (p *Player) Remove() {
	p.isRemoved = true
	if len(p.game.players) == 0 {
		p.game.isRemoved = true
	}
}

func (p Player) IsHost() bool {
	minPlayerId := p.ID
	for _, player := range p.game.players {
		if player.ID < minPlayerId {
			minPlayerId = player.ID
		}
	}
	return p.ID == minPlayerId
}
