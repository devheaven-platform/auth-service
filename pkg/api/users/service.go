package users

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
)

// Service represents the service object in the api
// resource. This object is used by the transport
// layer to interact with the platform layer.
type Service struct {
	platform Platform
}

// CreateService is used to intialize the service.
// It takes an platform as parameter and returns
// an service object.
func CreateService(platform Platform) Service {
	return Service{
		platform: platform,
	}
}

// GetAllUsers is used to retrieve all the users from
// the database. It returns an list of users and error
// if one occurred.
func (s *Service) GetAllUsers() ([]*domain.User, error) {
	return nil, nil
}

// GetUserByID is used to retrieve one user from the
// the database by his/her id. It takes an id as parameter
// and returns an user and error if one occurred.
func (s *Service) GetUserByID(id int) (*domain.User, error) {
	return nil, nil
}

// CreateUser is used to create a new user in the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (s *Service) CreateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}

// UpdateUser is used to update a user in the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (s *Service) UpdateUser(user domain.User) (*domain.User, error) {
	return nil, nil
}

// DeleteUser is used to delete a user from the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (s *Service) DeleteUser(user domain.User) (bool, error) {
	return false, nil
}
