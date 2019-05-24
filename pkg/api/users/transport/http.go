package transport

import (
	"encoding/json"
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/api/users"
	base "github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/devheaven-platform/auth-service/pkg/utils/validation"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
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
		t.RespondError(res, "An internal server error occurred", http.StatusInternalServerError)
		return
	}

	t.RespondJSON(res, http.StatusOK, result)
}

// getUserByID is used to retrieve one user. This function
// listens on the /users/{id} endpoint. It takes an ReponseWriter
// and Request as parameters.
func (t *transport) getUserByID(res http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(chi.URLParam(req, "id"))
	if err != nil {
		t.RespondError(res, "Id is invalid", http.StatusBadRequest)
		return
	}

	result, err := t.service.GetUserByID(id)
	if err == gorm.ErrRecordNotFound {
		t.RespondError(res, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.WithError(err).Warn("An error occurred while retrieving the user")
		t.RespondError(res, "An internal server error occurred", http.StatusInternalServerError)
	}

	t.RespondJSON(res, http.StatusOK, result)
}

// createUser is used to create a new user This function listens on
// the /users/ endpoint. It takes an ReponseWriter and Request as
// parameters.
func (t *transport) createUser(res http.ResponseWriter, req *http.Request) {
	type request struct {
		Firstname string   `json:"firstname" validate:"required,min=2,max=20"`
		Lastname  string   `json:"lastname" validate:"required,min=2,max=20"`
		Emails    []string `json:"emails" validate:"required,gte=1,dive,email"`
		Roles     []string `json:"roles" validate:"required,gte=1,dive,oneof=ROLE_USER ROLE_DEVELOPER ROLE_HR ROLE_MANAGER"`
		Password  string   `json:"password" validate:"required"`
	}

	data := request{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		t.RespondError(res, "An error occurred while converting the request body", http.StatusUnsupportedMediaType)
		return
	}

	errs := validation.Validate(data)
	if errs != nil {
		t.RespondValidationError(res, "One or more values are invalid", http.StatusBadRequest, errs)
		return
	}

	result, err := t.service.CreateUser(data.Firstname, data.Lastname, data.Emails, data.Roles, data.Password)
	if err != nil {
		log.WithError(err).Warn("An error occurred while creating the user")
		t.RespondError(res, "An internal server error occurred", http.StatusInternalServerError)
		return
	}

	t.RespondJSON(res, http.StatusCreated, result)
}

// updateUser is used to update a user. This function listens on
// the /users/{id} endpoint. It takes an ReponseWriter and Request
// as parameters.
func (t *transport) updateUser(res http.ResponseWriter, req *http.Request) {
	type request struct {
		Firstname string   `json:"firstname" validate:"omitempty,min=2,max=20"`
		Lastname  string   `json:"lastname" validate:"omitempty,min=2,max=20"`
		Emails    []string `json:"emails" validate:"omitempty,gte=1,dive,email"`
		Roles     []string `json:"roles" validate:"omitempty,gte=1,dive,oneof=ROLE_USER ROLE_DEVELOPER ROLE_HR ROLE_MANAGER"`
		Password  string   `json:"password" validate:"omitempty"`
	}

	data := request{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		t.RespondError(res, "An error occurred while converting the request body", http.StatusUnsupportedMediaType)
		return
	}

	id, err := uuid.Parse(chi.URLParam(req, "id"))
	if err != nil {
		t.RespondError(res, "Id is invalid", http.StatusBadRequest)
		return
	}

	errs := validation.Validate(data)
	if errs != nil {
		t.RespondValidationError(res, "One or more values are invalid", http.StatusBadRequest, errs)
		return
	}

	result, err := t.service.UpdateUser(id, data.Firstname, data.Lastname, data.Emails, data.Roles, data.Password)
	if err == gorm.ErrRecordNotFound {
		t.RespondError(res, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.WithError(err).Warn("An error occurred while updating the user")
		t.RespondError(res, "An internal server error occurred", http.StatusInternalServerError)
		return
	}

	t.RespondJSON(res, http.StatusOK, result)
}

// deleteUser is used to delete a user. This function listens on
// the /users/{id} endpoint. It takes an ReponseWriter and Request
// as parameters.
func (t *transport) deleteUser(res http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(chi.URLParam(req, "id"))
	if err != nil {
		t.RespondError(res, "Id is invalid", http.StatusBadRequest)
		return
	}

	result, err := t.service.DeleteUser(id)
	if err == gorm.ErrRecordNotFound {
		t.RespondError(res, "User not found", http.StatusNotFound)
		return
	}
	if err != nil || result != true {
		log.WithError(err).Warn("An error occurred while deleting the user")
		t.RespondError(res, "An internal server error occurred", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
