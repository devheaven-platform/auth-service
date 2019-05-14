package main

import (
	"net/http"
	"os"

	"github.com/devheaven-platform/auth-service/pkg/api"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

// main is invoked by the go compiler and is used
// to load the environment variables, database
// connection and the routes for the service.
func main() {
	// Environment
	host := os.Getenv("GO_HOST")
	port := os.Getenv("GO_PORT")

	router := api.CreateRouter()

	log.WithFields(log.Fields{
		"host": host,
		"port": port,
	}).Info("Started server")
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
