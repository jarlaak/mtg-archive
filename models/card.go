package models

import (
	"time"
)

type Card struct {
	ID uint
	CardName string
	MultiverseID uint
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"column:modified_at"`
}
