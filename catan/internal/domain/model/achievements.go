package model

func NewAchievements(game *Game) Achievements {
	var achievements Achievements

	longestRoadAchievement := NewLongestRoadAchievement(game)
	achievements.append(longestRoadAchievement)

	largestArmyAchievement := NewLargestArmyAchievement(game)
	achievements.append(largestArmyAchievement)

	return achievements
}

type Achievements []*Achievement

func (a Achievements) append(achievement *Achievement) {
	a = append(a, achievement)
}

func (a Achievements) remove(achievement *Achievement) {
	for idx := range a {
		if a[idx] == achievement {
			a = append(a[:idx], a[idx+1:]...)
			return
		}
	}
}

func (a Achievements) SetGame(game *Game) {
	for _, achievement := range a {
		achievement.SetGame(game)
	}
}

func (a Achievements) SetPlayer(player *Player) {
	for _, achievement := range a {
		achievement.SetPlayer(player)
	}
}
