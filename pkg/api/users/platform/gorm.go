package platform

import (
	"github.com/devheaven-platform/auth-service/pkg/api/users"
	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/jinzhu/gorm"
)

// TODO: go doc
type platform struct {
	db *gorm.DB
}

// TODO: go doc
func CreatePlatform(db *gorm.DB) users.Platform {
	return &platform{
		db: db,
	}
}

// TODO: go doc
func (p *platform) GetAllUsers() ([]*domain.User, error) {
	return nil, nil
}

// TODO: go doc
func (p *platform) GetUserByID(id int) (*domain.User, error) {
	return nil, nil
}

// TODO: go doc
func (p *platform) CreateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}

// TODO: go doc
func (p *platform) UpdateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}

// TODO: go doc
func (p *platform) DeleteUser(user domain.User) (bool, error) {
	return false, nil
}
