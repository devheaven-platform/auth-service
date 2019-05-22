package api

import (
	"net/http"

	authService "github.com/devheaven-platform/auth-service/pkg/api/auth"
	authPlatform "github.com/devheaven-platform/auth-service/pkg/api/auth/platform"
	authTransport "github.com/devheaven-platform/auth-service/pkg/api/auth/transport"
	healthTransport "github.com/devheaven-platform/auth-service/pkg/api/health/transport"
	metricsTransport "github.com/devheaven-platform/auth-service/pkg/api/metrics/transport"
	swaggerTransport "github.com/devheaven-platform/auth-service/pkg/api/swagger/transport"
	usersService "github.com/devheaven-platform/auth-service/pkg/api/users"
	usersPlatform "github.com/devheaven-platform/auth-service/pkg/api/users/platform"
	usersTransport "github.com/devheaven-platform/auth-service/pkg/api/users/transport"
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

	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RealIP,
		middleware.Recoverer,
		logging.NewStructuredLogger(log.StandardLogger()),
	)

	transport := transport.BaseTransport{}
	router.Route("/", func(r chi.Router) {
		r.Mount("/health", healthTransport.CreateTransport())
		r.Mount("/metrics", metricsTransport.CreateTransport())
		r.Mount("/docs", swaggerTransport.CreateTransport())
		r.Mount("/auth", authTransport.CreateTransport(authService.CreateService(authPlatform.CreatePlatform(db))))
		r.Mount("/users", usersTransport.CreateTransport(usersService.CreateService(usersPlatform.CreatePlatform(db))))

		r.NotFound(func(res http.ResponseWriter, req *http.Request) {
			transport.RespondError(res, "Resource not found", 404)
		})
		r.MethodNotAllowed(func(res http.ResponseWriter, req *http.Request) {
			transport.RespondError(res, "Method not allowed", 405)
		})
	})

	return router
}
