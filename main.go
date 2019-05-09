package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load environment
	err := godotenv.Load()

	if err != nil {
		log.WithError(err).Fatal("An error occurred while loading the environment variables")
	}

	host := os.Getenv("GO_HOST")
	port := os.Getenv("GO_PORT")

	// Add health check
	// TODO: add health check

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
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
