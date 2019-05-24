package auth

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/google/uuid"
)

// Platform represents the platform in the api
// resource. This interface is used by the service
// layer to interact with the database.
type Platform interface {
	Me(id uuid.UUID) (domain.User, error)
}
