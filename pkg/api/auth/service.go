package auth

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/google/uuid"
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

// Me is used to retrieve the current user from
// the database. It takes an user id as parameter
// and returns an user object and error if one
// occurred.
func (s *Service) Me(id uuid.UUID) ([]*domain.User, error) {
	return nil, nil
}

// Login is used to log user into the system.
// It takes an email and password as input
// if the credentials are valid an token object
// will be returned.
func (s *Service) Login(email string, password string) (*domain.Token, error) {
	return nil, nil
}
