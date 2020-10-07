package dto

import "github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"

type ResourceCardDTO struct {
	ID           int
	ResourceType model.ResourceType
}
