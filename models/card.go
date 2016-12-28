package models

import (
	"time"
)

type Card struct {
	ID           uint      `json:"id"`
	CardInfoID   uint      `json:"-"`
	CardInfo     CardInfo  `json:"card_info"`
	MultiverseID uint      `gorm:"default:nil" json:"multiverse_id,omitempty"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `gorm:"column:modified_at" json:"-"`
}
