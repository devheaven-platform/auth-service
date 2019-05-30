package platform

import (
	"github.com/devheaven-platform/auth-service/pkg/api/users"
	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/google/uuid"
	"github.com/imdario/mergo"
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
func CreatePlatform(db *gorm.DB) users.Platform {
	return &platform{
		db: db,
	}
}

// GetAllUsers is used to retrieve all the users from
// the database. It returns an slice of users and error
// if one occurred.
func (p *platform) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := p.db.Set("gorm:auto_preload", true).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// GetUserByID is used to retrieve one user from the
// the database by his/her id. It takes an id as parameter
// and returns an user and error if one occurred.
func (p *platform) GetUserByID(id uuid.UUID) (domain.User, error) {
	var user domain.User
	if err := p.db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// CreateUser is used to create a new user in the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (p *platform) CreateUser(user domain.User) (domain.User, error) {
	if err := p.db.Create(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// UpdateUser is used to update a user in the database.
// It takes an user and an user with updates as parameters
// and returns an user and error if one occurred.
func (p *platform) UpdateUser(user domain.User, update domain.User) (domain.User, error) {
	if err := mergo.Merge(&update, user); err != nil {
		return domain.User{}, err
	}
	update.ID = user.ID

	if err := p.db.Save(&update).Error; err != nil {
		return domain.User{}, err
	}
	return update, nil
}

// DeleteUser is used to delete a user from the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (p *platform) DeleteUser(user domain.User) (bool, error) {
	if err := p.db.Delete(user).Error; err != nil {
		return false, err
	}
	return true, nil
}
