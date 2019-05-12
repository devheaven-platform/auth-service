package main

import (
	"net/http"
	"os"

	"github.com/devheaven-platform/auth-service/controllers"
	"github.com/devheaven-platform/auth-service/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// main is invoked by the go compiler and is used
// to load the environment variables, database
// connection and the routes for the service.
func main() {
	// Load environment
	err := godotenv.Load()

	if err != nil {
		log.WithError(err).Fatal("An error occurred while loading the environment variables")
	}

	host := os.Getenv("GO_HOST")
	port := os.Getenv("GO_PORT")

	db, err := utils.OpenConnection()
	db.LogMode(false)

	if err != nil {
		log.WithError(err).Fatal("An error occurred while connecting to the database")
	}
	defer db.Close()

	// Create controllers
	healthController := controllers.CreateHealthController()
	metricsController := controllers.CreateMetricsController()
	swaggerController := controllers.CreateSwaggerController()

	// Create main router
	router := chi.NewRouter()

	// Add middleware
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RealIP,
		middleware.Recoverer,
		utils.NewStructuredLogger(log.StandardLogger()),
	)

	// Add routes
	router.Route("/", func(r chi.Router) {
		// General
		r.Mount("/health", healthController)
		r.Mount("/metrics", metricsController)
		r.Mount("/docs", swaggerController)

		// Service

		// Errors
		r.NotFound(func(res http.ResponseWriter, req *http.Request) {
			utils.RespondError(res, "Resource not found", 404)
		})
		r.MethodNotAllowed(func(res http.ResponseWriter, req *http.Request) {
			utils.RespondError(res, "Method not allowed", 405)
		})
	})

	log.WithFields(log.Fields{
		"host": host,
		"port": port,
	}).Info("Started server")
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
