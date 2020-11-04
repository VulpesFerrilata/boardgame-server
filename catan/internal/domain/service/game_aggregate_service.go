package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
)

type GameAggregateService interface {
	GetById(ctx context.Context, id uint) (*model.Game, error)
	Save(ctx context.Context, game *model.Game) error
}

type gameAggregateService struct {
	gameService            GameService
	playerAggregateService PlayerAggregateService
	diceService            DiceService
	achievementService     AchievementService
	resourceCardService    ResourceCardService
	developmentCardService DevelopmentCardService
	robberService          RobberService
}

func (gas gameAggregateService) GetById(ctx context.Context, id uint) (*model.Game, error) {
	game, err := gas.gameService.GetGameRepository().GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := gas.playerAggregateService.FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	players.SetGame(game)

	dices, err := gas.diceService.GetDiceRepository().FindByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	dices.SetGame(game)

	achievements, err := gas.achievementService.GetAchievementRepository().FindByGameIdByPlayerId(ctx, game.ID, nil)
	if err != nil {
		return nil, err
	}
	achievements.SetGame(game)

	resourceCards, err := gas.resourceCardService.GetResourceCardRepository().FindByGameIdByPlayerId(ctx, game.ID, nil)
	if err != nil {
		return nil, err
	}
	resourceCards.SetGame(game)

	developmentCards, err := gas.developmentCardService.GetDevelopmentCardRepository().FindByGameIdByPlayerId(ctx, game.ID, nil)
	if err != nil {
		return nil, err
	}
	developmentCards.SetGame(game)

	robber, err := gas.robberService.GetRobberRepository().GetByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	robber.SetGame(game)

	return game, nil
}

func (gas gameAggregateService) Save(ctx context.Context, game *model.Game) error {
	if err := gas.gameService.Save(ctx, game); err != nil {
		return err
	}

	players := game.GetPlayers()
	for _, player := range players {
		if err := gas.playerAggregateService.Save(ctx, player); err != nil {
			return err
		}
	}

	dices := game.GetDices()
	for _, dice := range dices {
		if err := gas.diceService.Save(ctx, dice); err != nil {
			return err
		}
	}

	achievements := game.GetAchievements()
	for _, achievement := range achievements {
		if err := gas.achievementService.Save(ctx, achievement); err != nil {
			return err
		}
	}

	return nil
}
