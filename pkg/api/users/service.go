package users

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
)

// TODO: go doc
type Service struct {
	platform Platform
}

// TODO: go doc
func CreateService(platform Platform) Service {
	return Service{
		platform: platform,
	}
}

// TODO: go doc
func (s *Service) GetAllUsers() ([]*domain.User, error) {
	return nil, nil
}

// TODO: go doc
func (s *Service) GetUserByID(id int) (*domain.User, error) {
	return nil, nil
}

// TODO: go doc
func (s *Service) CreateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}

// TODO: go doc
func (s *Service) UpdateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}

// TODO: go doc
func (s *Service) DeleteUser(user domain.User) (bool, error) {
	return false, nil
}
