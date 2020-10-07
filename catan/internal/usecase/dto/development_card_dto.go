package dto

import "github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"

type DevelopmentCardDTO struct {
	ID              int
	DevelopmentType model.DevelopmentType
}
