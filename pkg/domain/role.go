package domain

import (
	"time"
)

// TODO: go doc
type Role struct {
	Role      string    `gorm:"primary_key" json:"role"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}
