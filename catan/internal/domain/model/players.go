package model

type Players []*Player

func (p Players) append(player *Player) {
	p = append(p, player)
}

func (p Players) SetGame(game *Game) {
	for _, player := range p {
		player.SetGame(game)
	}
}

func (p Players) Filter(f func(p *Player) bool) Players {
	var players Players
	for _, player := range p {
		if f(player) {
			players.append(player)
		}
	}
	return players
}

func (p Players) First() *Player {
	if len(p) > 0 {
		return p[0]
	}
	return nil
}
