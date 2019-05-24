package transport

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/api/auth"
	"github.com/devheaven-platform/auth-service/pkg/utils/middleware"
	base "github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
	router.Group(func(router chi.Router) {
		router.Use(middleware.Authenticator)
		router.Get("/me", transport.me)
	})
	router.Group(func(router chi.Router) {
		router.Post("/login", transport.login)
	})

	return router
}

// me is used to retrieve the current user. This function
// listens on the /auth/me endpoint. It takes an ReponseWriter
// and Request as parameters.
func (t *transport) me(res http.ResponseWriter, req *http.Request) {
	_, claims, _ := jwtauth.FromContext(req.Context())
	id, err := uuid.Parse(chi.URLParam(req, "id"))
	if err != nil {
		t.RespondError(res, "Your not authorized to access this resource", http.StatusUnauthorized)
		return
	}

	result, err := t.service.Me(id)
	if err != nil {
		log.WithError(err).Warn("An error occurred while retrieving your account")
		t.RespondError(res, "An internal server error occurred", http.StatusInternalServerError)
	}

	t.RespondJSON(res, http.StatusOK, result)
}

// login is used to log a user into the system. This function
// listens on the /auth/login endpoint. It takes an ResponseWriter
// and Request as parameters.
func (t *transport) login(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}
