package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// TODO: go doc
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Emails    []Email   `gorm:"many2many:user_emails" json:"emails"`
	Roles     []Role    `gorm:"many2many:user_roles" json:"roles"`
	Password  string    `json:"password,omitempty" `
	Enabled   bool      `json:"enabled" gorm:"default:true"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}

// TODO: go doc
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}

	scope.SetColumn("Password", string(hash[:]))
	return scope.SetColumn("ID", uuid)
}

// TODO: go doc
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		return err
	}

	return scope.SetColumn("Password", string(hash[:]))
}
