package api

import (
	"net/http"
	"os"

	authService "github.com/devheaven-platform/auth-service/pkg/api/auth"
	authPlatform "github.com/devheaven-platform/auth-service/pkg/api/auth/platform"
	authTransport "github.com/devheaven-platform/auth-service/pkg/api/auth/transport"
	healthTransport "github.com/devheaven-platform/auth-service/pkg/api/health/transport"
	metricsTransport "github.com/devheaven-platform/auth-service/pkg/api/metrics/transport"
	swaggerTransport "github.com/devheaven-platform/auth-service/pkg/api/swagger/transport"
	usersService "github.com/devheaven-platform/auth-service/pkg/api/users"
	usersPlatform "github.com/devheaven-platform/auth-service/pkg/api/users/platform"
	usersTransport "github.com/devheaven-platform/auth-service/pkg/api/users/transport"
	"github.com/devheaven-platform/auth-service/pkg/utils/logging"
	"github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// CreateRouter is invoked by the main.go binary.
// This function creates the router. It takes an
// instance of a gorm database as parameter and
// returns an instance of chi router.
func CreateRouter(db *gorm.DB) chi.Router {
	auth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)

	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
	})

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		cors.Handler,
		middleware.RealIP,
		middleware.Recoverer,
		logging.NewStructuredLogger(log.StandardLogger()),
		jwtauth.Verifier(auth),
	)

	transport := transport.BaseTransport{}
	router.Route("/", func(r chi.Router) {
		r.Mount("/health", healthTransport.CreateTransport())
		r.Mount("/metrics", metricsTransport.CreateTransport())
		r.Mount("/docs", swaggerTransport.CreateTransport())
		r.Mount("/auth", authTransport.CreateTransport(authService.CreateService(authPlatform.CreatePlatform(db), auth)))
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
