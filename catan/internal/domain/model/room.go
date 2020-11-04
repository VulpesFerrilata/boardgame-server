package model

import "github.com/VulpesFerrilata/catan/internal/domain/datamodel"

type Room struct {
	ID          uint
	Status      datamodel.GameStatus
	PlayerCount int
}
