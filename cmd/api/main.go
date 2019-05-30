package main

import (
	"net/http"
	"os"

	"github.com/devheaven-platform/auth-service/pkg/api"
	"github.com/devheaven-platform/auth-service/pkg/utils/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

// main is invoked by the go compiler and is used
// to load the environment variables, database
// connection and the routes for the service.
func main() {
	host := os.Getenv("GO_HOST")
	port := os.Getenv("GO_PORT")

	db, err := db.OpenConnection()
	db.LogMode(true)

	if err != nil {
		log.WithError(err).Fatal("An error occurred while connecting to the database")
	}
	defer db.Close()

	router := api.CreateRouter(db)

	log.WithFields(log.Fields{
		"host": host,
		"port": port,
	}).Info("Started server")
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
