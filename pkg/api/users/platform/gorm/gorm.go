package gorm

import (
	"github.com/devheaven-platform/auth-service/pkg/api/users/platform"
	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/jinzhu/gorm"
)

type userGormPlatform struct {
	db *gorm.DB
}

func CreateUserGormPlatform(db *gorm.DB) platform.UserPlatform {
	return &userGormPlatform{
		db: db,
	}
}

func (p *userGormPlatform) GetUserById(id int) (*domain.User, error) {
	var user = new(domain.User)
	if err := p.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *userGormPlatform) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := p.db.Find(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
