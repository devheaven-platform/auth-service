package platform

import (
	"github.com/devheaven-platform/auth-service/pkg/api/users"
	"github.com/devheaven-platform/auth-service/pkg/domain"
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
// the database. It returns an list of users and error
// if one occurred.
func (p *platform) GetAllUsers() ([]*domain.User, error) {
	return nil, nil
}

// GetUserByID is used to retrieve one user from the
// the database by his/her id. It takes an id as parameter
// and returns an user and error if one occurred.
func (p *platform) GetUserByID(id int) (*domain.User, error) {
	return nil, nil
}

// CreateUser is used to create a new user in the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (p *platform) CreateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}

// UpdateUser is used to update a user in the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (p *platform) UpdateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}

// DeleteUser is used to delete a user from the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (p *platform) DeleteUser(user domain.User) (bool, error) {
	return false, nil
}
