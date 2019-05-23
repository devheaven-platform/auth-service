package transport

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/api/users"
	base "github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// This object is used to group all the transport
// functions together.
type transport struct {
	base.BaseTransport
	service users.Service
}

// CreateTransport is used to intialize the transport
// layer. It takes an service as parameter and returns
// an router object.
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

// getAllUsers is used to retrieve all the users. This function
// listens on the /users/ endpoint. It takes an ReponseWriter
// and Request as parameters.
func (t *transport) getAllUsers(res http.ResponseWriter, req *http.Request) {
	result, err := t.service.GetAllUsers()

	if err != nil {
		log.WithError(err).Warn("An error occurred while retrieving the users")
		t.RespondError(res, "InternalServerError", http.StatusInternalServerError)
		return
	}

	t.RespondJSON(res, http.StatusOK, result)
}

// getUserByID is used to retrieve one user. This function
// listens on the /users/{id} endpoint. It takes an ReponseWriter
// and Request as parameters.
func (t *transport) getUserByID(res http.ResponseWriter, req *http.Request) {
	_, err := uuid.Parse(chi.URLParam(req, "id"))

	if err != nil {
		t.RespondError(res, "Id is invalid", http.StatusBadRequest)
		return
	}

	t.RespondError(res, "Not Implemented", 501)
}

// createUser is used to create a new user This function listens on
// the /users/ endpoint. It takes an ReponseWriter and Request as
// parameters.
func (t *transport) createUser(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}

// updateUser is used to update a user. This function listens on
// the /users/{id} endpoint. It takes an ReponseWriter and Request
// as parameters.
func (t *transport) updateUser(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}

// deleteUser is used to delete a user. This function listens on
// the /users/{id} endpoint. It takes an ReponseWriter and Request
// as parameters.
func (t *transport) deleteUser(res http.ResponseWriter, req *http.Request) {
	t.RespondError(res, "Not Implemented", 501)
}
