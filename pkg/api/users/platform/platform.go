package platform

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
)

type UserPlatform interface {
	GetUserById(id int) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}
