package model

import "gorm.io/gorm"

type Harbor struct {
	*gorm.Model
	Q          int
	R          int
	HarborType HarborType
	FieldQ     int
	FieldR     int
}

type HarborType string

const (
	HT_BRICK   HarborType = "BRICK"
	HT_GRAIN   HarborType = "GRAIN"
	HT_SHEEP   HarborType = "SHEEP"
	HT_STONE   HarborType = "STONE"
	HT_WOOD    HarborType = "WOOD"
	HT_GENERAL HarborType = "GENERAL"
)
