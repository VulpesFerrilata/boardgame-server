package dto

import "github.com/VulpesFerrilata/boardgame-server/catan/internal/domain/model"

type FieldDTO struct {
	ID        int
	Q         int
	R         int
	Number    int
	FieldType model.FieldType
}
