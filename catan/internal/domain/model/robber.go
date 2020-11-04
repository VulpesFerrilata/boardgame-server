package model

type Robber struct {
	*Robber
	game *Game
}

func (r *Robber) SetGame(game *Game) {
	r.game.ID = game.ID
	r.game = game
	r.game.robber = r
}
