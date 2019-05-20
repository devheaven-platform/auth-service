package domain

import (
	"time"
)

// TODO: go doc
type Email struct {
	Email     string    `gorm:"primary_key" json:"email"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}
