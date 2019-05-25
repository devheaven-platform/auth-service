package auth

import (
	"os"
	"strconv"
	"time"

	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service represents the service object in the api
// resource. This object is used by the transport
// layer to interact with the platform layer.
type Service struct {
	platform Platform
	auth     *jwtauth.JWTAuth
}

// CreateService is used to intialize the service.
// It takes an platform as parameter and returns
// an service object.
func CreateService(platform Platform, auth *jwtauth.JWTAuth) Service {
	return Service{
		platform: platform,
		auth:     auth,
	}
}

// Me is used to retrieve the current user from
// the database. It takes an user id as parameter
// and returns an user object and error if one
// occurred.
func (s *Service) Me(id uuid.UUID) (domain.User, error) {
	return s.platform.Me(id)
}

// Login is used to log user into the system.
// It takes an email and password as input
// if the credentials are valid an token object
// will be returned.
func (s *Service) Login(email string, password string) (domain.Token, error) {
	user, err := s.platform.GetByEmail(email)
	if err != nil {
		return domain.Token{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.Token{}, err
	}

	r := []string{}
	for _, role := range user.Roles {
		r = append(r, role.Role)
	}
	iss := os.Getenv("JWT_ISSUER")
	v, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES"))
	exp := time.Now().Add(time.Hour * time.Duration(v)).Unix()
	iat := time.Now().Unix()

	_, token, err := s.auth.Encode(jwtauth.Claims{"sub": user.ID.String(), "roles": r, "iss": iss, "exp": exp, "iat": iat})
	if err != nil {
		return domain.Token{}, err
	}

	return domain.Token{Token: token, Expires: exp}, nil
}
