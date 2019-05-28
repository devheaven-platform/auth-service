package domain

import (
	"time"

	"github.com/google/uuid"
)

// Email represents an email in the system. The object
// contains the email value which also is the primary
// key.
type Email struct {
	Email     string    `gorm:"primary_key" json:"email"`
	UserID    uuid.UUID `json:"-"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}
