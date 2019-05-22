package users

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
)

// TODO: go doc
type Platform interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	CreateUser(domain.User) (*domain.User, error)
	UpdateUser(domain.User) (*domain.User, error)
	DeleteUser(domain.User) (bool, error)
}
