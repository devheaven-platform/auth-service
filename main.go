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
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		r.Mount("/health", healthController)
	})

	// Add prometheus
	http.Handle("/metrics", promhttp.Handler())

	// Add swagger
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))
	http.HandleFunc("/docs/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./spec/swagger.yaml")
	})

	log.WithFields(log.Fields{
		"host": host,
		"port": port,
	}).Info("Started server")
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
