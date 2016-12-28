package models

import (
	"time"
)

type CardInfo struct {
	ID        uint      `json:"id"`
	CardName  string    `json:"card_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `gorm:"column:modified_at" json:"-"`
}
