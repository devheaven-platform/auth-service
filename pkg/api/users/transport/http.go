package transport

import (
	"log"
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/api/users/platform"
	"github.com/go-chi/chi"
)

type userTransport struct {
	userPlatform platform.UserPlatform
}

func CreateTransport(platform platform.UserPlatform) *chi.Mux {
	// Create the controller
	transport := &userTransport{
		userPlatform: platform,
	}

	// Create routes
	router := chi.NewRouter()
	router.Get("/", transport.getAllUsers)
	router.Get("/{id}", transport.getUserByID)

	// Return the router
	return router
}

func (platform *userTransport) getAllUsers(res http.ResponseWriter, req *http.Request) {
	log.Println("Func: GetAllUsers")
	res.Header().Set("Content-Type", "application/json")
}

func (platform *userTransport) getUserByID(res http.ResponseWriter, req *http.Request) {
	log.Println("Func: GetUserByID")
}

func (platform *userTransport) updateUser(res http.ResponseWriter, req *http.Request) {
	log.Println("Func: UpdateUser")
}
