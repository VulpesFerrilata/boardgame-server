package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

func NewLongestRoadAchievement(game *Game) *Achievement {
	achievement := new(Achievement)
	achievement.Type = datamodel.AT_LONGEST_ROAD
	achievement.BonusPoints = 2
	achievement.SetGame(game)
	return achievement
}

func NewLargestArmyAchievement(game *Game) *Achievement {
	achievement := new(Achievement)
	achievement.Type = datamodel.AT_LARGEST_ARMY
	achievement.BonusPoints = 2
	achievement.SetGame(game)
	return achievement
}

type Achievement struct {
	*datamodel.Achievement
	game   *Game
	player *Player
}

func (a *Achievement) SetGame(game *Game) {
	a.GameID = game.ID
	a.game = game
	a.game.achievements.append(a)
}

func (a *Achievement) SetPlayer(player *Player) {
	if a.player != nil {
		a.PlayerID = nil
		a.player.achievements.remove(a)
	}

	a.player = player
	if player != nil {
		a.PlayerID = &player.ID
		player.achievements.append(a)
	}
}
