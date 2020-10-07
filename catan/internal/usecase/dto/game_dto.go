package dto

import "github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"

type GameDTO struct {
	ID               int
	PlayerInTurn     int
	Turn             int
	Status           model.GameStatus
	Dices            []*DiceDTO
	Players          []*PlayerDTO
	ResourceCards    []*ResourceCardDTO
	DevelopmentCards []*DevelopmentCardDTO
	Fields           []*FieldDTO
	Constructions    []*ConstructionDTO
}
