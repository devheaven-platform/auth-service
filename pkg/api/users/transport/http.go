package transport

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/api/users"
	base "github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/chi"
)

// TODO: go doc
type transport struct {
	base.BaseTransport
	service users.Service
}

// TODO: go doc
func CreateTransport(service users.Service) *chi.Mux {
	transport := &transport{
		service: service,
	}

	router := chi.NewRouter()
	router.Get("/", transport.getAllUsers)
	router.Get("/{id}", transport.getUserByID)
	router.Post("/", transport.createUser)
	router.Patch("/{id}", transport.updateUser)
	router.Delete("/{id}", transport.deleteUser)

	return router
}

// TODO: go doc
func (t *transport) getAllUsers(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}

// TODO: go doc
func (t *transport) getUserByID(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}

// TODO: go doc
func (t *transport) createUser(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}

// TODO: go doc
func (t *transport) updateUser(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}

// TODO: go doc
func (t *transport) deleteUser(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}
