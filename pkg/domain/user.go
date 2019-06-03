package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User represents an user in the system. The object
// contains a list of emails that user has and an list
// of roles the user has and the password of the user.
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Emails    []Email   `json:"emails"`
	Roles     []Role    `gorm:"many2many:user_roles" json:"roles"`
	Password  string    `json:"-"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}

// BeforeCreate is invoked by Gorm before an user is
// inserted in the database. This function is used to
// hash the password of the user. It takes an gorm
// Scope as parameter and returns an error if one occurred.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}

	return scope.SetColumn("Password", string(hash[:]))
}

// BeforeUpdate is invoked by Gorm before an user is
// updated in the database. This function is used to
// re-hash the password of the user. It takes an gorm
// Scope as parameter and returns an error if one occurred.
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		return err
	}

	return scope.SetColumn("Password", string(hash[:]))
}
