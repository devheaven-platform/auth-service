package transport

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/chi"
)

// This object is used to group all the transport
// functions together.
type httpTransport struct {
	transport.BaseHTTPTransport
}

// CreateHTTPTransport is used to create the health
// controller. It returns an instance of a chi.Mux router.
func CreateHTTPTransport() *chi.Mux {
	transport := httpTransport{}

	router := chi.NewRouter()
	router.Get("/", transport.getHealth)

	return router
}

// getHealth handles the /health/ route. It takes an
// http.ResponseWriter and http.Request as parameters.
func (t httpTransport) getHealth(res http.ResponseWriter, req *http.Request) {
	t.RespondJSON(res, http.StatusOK, domain.Health{
		Message: "Service is running",
	})
}
