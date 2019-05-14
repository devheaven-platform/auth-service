package api

import (
	"net/http"

	healthTransport "github.com/devheaven-platform/auth-service/pkg/api/health/transport"
	metricsTransport "github.com/devheaven-platform/auth-service/pkg/api/metrics/transport"
	swaggerTransport "github.com/devheaven-platform/auth-service/pkg/api/swagger/transport"
	"github.com/devheaven-platform/auth-service/pkg/utils/db"
	"github.com/devheaven-platform/auth-service/pkg/utils/logging"
	"github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

// CreateRouter is invoked by the main.go binary.
// This function creates the router and connects
// to the database.
func CreateRouter() chi.Router {
	db, err := db.OpenConnection()
	db.LogMode(false)

	if err != nil {
		log.WithError(err).Fatal("An error occurred while connecting to the database")
	}
	defer db.Close()

	// Create main router
	router := chi.NewRouter()

	// Add middleware
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RealIP,
		middleware.Recoverer,
		logging.NewStructuredLogger(log.StandardLogger()),
	)

	// Add routes
	transport := transport.BaseTransport{}
	router.Route("/", func(r chi.Router) {
		// General
		r.Mount("/health", healthTransport.CreateTransport())
		r.Mount("/metrics", metricsTransport.CreateTransport())
		r.Mount("/docs", swaggerTransport.CreateTransport())

		// Errors
		r.NotFound(func(res http.ResponseWriter, req *http.Request) {
			transport.RespondError(res, "Resource not found", 404)
		})
		r.MethodNotAllowed(func(res http.ResponseWriter, req *http.Request) {
			transport.RespondError(res, "Method not allowed", 405)
		})
	})

	return router
}
