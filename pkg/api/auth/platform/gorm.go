package platform

import (
	"github.com/devheaven-platform/auth-service/pkg/api/auth"
	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// This object is used to group all the platform
// functions together and implements the Platform
// interface.
type platform struct {
	db *gorm.DB
}

// CreatePlatform is used to intialize the platform
// layer. It takes an gorm db as parameter and returns
// an instance of the Platform interface.
func CreatePlatform(db *gorm.DB) auth.Platform {
	return &platform{
		db: db,
	}
}

// Me is used to retrieve one user from the
// the database by his/her id. It takes an id as parameter
// and returns an user and error if one occurred.
func (p *platform) Me(id uuid.UUID) (domain.User, error) {
	var user domain.User
	if err := p.db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// GetByEmail is used to retrieve one user from the
// the database by his/her email. It takes an email
// as parameter  and returns an user and error if
// one occurred.
func (p *platform) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := p.db.Set("gorm:auto_preload", true).Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
