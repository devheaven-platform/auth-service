package domain

import (
	"time"
)

// Role represents an role in the system. The object
// contains the role value which also is the primary
// key.
type Role struct {
	Role      string    `gorm:"primary_key" json:"role"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}
