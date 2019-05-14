package transport

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/chi"
)

// healthTransport represents a health transport object.
type healthTransport struct {
	transport.BaseTransport
}

// CreateTransport is used to create the health
// controller. It returns an instance of a chi.Mux router.
func CreateTransport() *chi.Mux {
	transport := healthTransport{}

	router := chi.NewRouter()
	router.Get("/", transport.getHealth)

	return router
}

// getHealth handles the /health/ route. It takes an
// http.ResponseWriter and http.Request as parameters.
func (t healthTransport) getHealth(res http.ResponseWriter, req *http.Request) {
	t.RespondJSON(res, http.StatusOK, domain.Health{
		Message: "Service is running",
	})
}
