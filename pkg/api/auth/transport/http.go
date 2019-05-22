package transport

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/api/auth"
	base "github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/chi"
)

// This object is used to group all the transport
// functions together.
type transport struct {
	base.BaseTransport
	service auth.Service
}

// CreateTransport is used to intialize the transport
// layer. It takes an service as parameter and returns
// an router object.
func CreateTransport(service auth.Service) *chi.Mux {
	transport := &transport{
		service: service,
	}

	router := chi.NewRouter()
	router.Get("/me", transport.me)
	router.Get("/login", transport.login)

	return router
}

// me is used to retrieve the current user. This function
// listens on the /auth/me endpoint. It takes an ReponseWriter
// and Request as parameters.
func (t *transport) me(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}

// login is used to log a user into the system. This function
// listens on the /auth/login endpoint. It takes an ResponseWriter
// and Request as parameters.
func (t *transport) login(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}
