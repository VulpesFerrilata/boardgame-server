package service

import (
	"context"

	"github.com/VulpesFerrilata/catan/internal/domain/model"
	"github.com/VulpesFerrilata/grpc/protoc/user"
)

type PlayerAggregateService interface {
	FindByGameId(ctx context.Context, gameId uint) (model.Players, error)
	Save(ctx context.Context, player *model.Player) error
}

type playerAggregateService struct {
	playerService          PlayerService
	userService            user.UserService
	achievementService     AchievementService
	resourceCardService    ResourceCardService
	developmentCardService DevelopmentCardService
}

func (pas playerAggregateService) FindByGameId(ctx context.Context, gameId uint) (model.Players, error) {
	players, err := pas.playerService.GetPlayerRepository().FindByGameId(ctx, gameId)
	if err != nil {
		return nil, err
	}

	for _, player := range players {
		userRequest := new(user.UserRequest)
		userRequest.ID = int64(player.UserID)
		userPb, err := pas.userService.GetUserById(ctx, userRequest)
		if err != nil {
			return nil, err
		}
		model.NewUser(player, userPb)

		achievements, err := pas.achievementService.GetAchievementRepository().FindByGameIdByPlayerId(ctx, gameId, &player.ID)
		if err != nil {
			return nil, err
		}
		achievements.SetPlayer(player)

		resourceCards, err := pas.resourceCardService.GetResourceCardRepository().FindByGameIdByPlayerId(ctx, gameId, &player.ID)
		if err != nil {
			return nil, err
		}
		resourceCards.SetPlayer(player)

		developmentCards, err := pas.developmentCardService.GetDevelopmentCardRepository().FindByGameIdByPlayerId(ctx, gameId, &player.ID)
		if err != nil {
			return nil, err
		}
		developmentCards.SetPlayer(player)
	}

	return players, nil
}

func (pas playerAggregateService) Save(ctx context.Context, player *model.Player) error {
	return pas.playerService.Save(ctx, player)
}
