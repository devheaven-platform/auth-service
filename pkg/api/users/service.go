package users

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

// GetAllUsers is used to retrieve all the users from
// the database. It returns an slice of users and error
// if one occurred.
func (s *Service) GetAllUsers() ([]domain.User, error) {
	return s.platform.GetAllUsers()
}

// GetUserByID is used to retrieve one user from the
// the database by his/her id. It takes an id as parameter
// and returns an user and error if one occurred.
func (s *Service) GetUserByID(id uuid.UUID) (domain.User, error) {
	return s.platform.GetUserByID(id)
}

// CreateUser is used to create a new user in the database.
// It takes an slice of emails, slice of roles and password
// as parameters and returns an user and error if one occurred.
func (s *Service) CreateUser(emails []string, roles []string, password string) (domain.User, error) {
	e := []domain.Email{}
	for _, email := range emails {
		e = append(e, domain.Email{
			Email: email,
		})
	}

	r := []domain.Role{}
	for _, role := range roles {
		r = append(r, domain.Role{
			Role: role,
		})
	}

	return s.platform.CreateUser(domain.User{
		Emails:   e,
		Roles:    r,
		Password: password,
	})
}

// UpdateUser is used to update a user in the database.
// It takes an slice of emails, slice of roles and
// password as parameters and returns an user and error
// if one occurred.
func (s *Service) UpdateUser(id uuid.UUID, emails []string, roles []string, password string) (domain.User, error) {
	user, err := s.platform.GetUserByID(id)
	if err != nil {
		return domain.User{}, err
	}

	e := []domain.Email{}
	if len(emails) > 0 {
		for _, email := range emails {
			e = append(e, domain.Email{
				Email: email,
			})
		}
	} else {
		e = user.Emails
	}

	r := []domain.Role{}
	if len(roles) > 0 {
		for _, role := range roles {
			r = append(r, domain.Role{
				Role: role,
			})
		}
	} else {
		r = user.Roles
	}

	return s.platform.UpdateUser(user, domain.User{
		Emails:   e,
		Roles:    r,
		Password: password,
	})
}

// DeleteUser is used to delete a user from the database.
// It takes an user as parameter and returns an user and error
// if one occurred.
func (s *Service) DeleteUser(id uuid.UUID) (bool, error) {
	user, err := s.platform.GetUserByID(id)
	if err != nil {
		return false, err
	}

	return s.platform.DeleteUser(user)
}
