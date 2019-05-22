package domain

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User represents an user in the system. The object
// contains some basic properties such as firstname,
// lastname, password and enabled. It also contains
// a list of emails that user has and an list of roles
// the user has.
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Emails    []Email   `gorm:"many2many:user_emails" json:"emails"`
	Roles     []Role    `gorm:"many2many:user_roles" json:"roles"`
	Password  string    `json:"password,omitempty" `
	Enabled   bool      `gorm:"default:true" json:"enabled"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}

// BeforeCreate is invoked by Gorm before an user is
// inserted in the database. This function is used to
// hash the passowrd of the user. It takes an gorm
// Scope as parameter and returns an error if one occurred.
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
