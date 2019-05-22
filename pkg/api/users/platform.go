package users

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
)

// Platform represents the platform in the api
// resource. This interface is used by the service
// layer to interact with the database.
type Platform interface {
	GetAllUsers() ([]*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	CreateUser(domain.User) (*domain.User, error)
	UpdateUser(domain.User) (*domain.User, error)
	DeleteUser(domain.User) (bool, error)
}
