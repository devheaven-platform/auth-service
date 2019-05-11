package utils

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// OpenConnection is used to open a database connection
// depending on the environment. It returns and db instance
// and error if one occurred.
func OpenConnection() (db *gorm.DB, err error) {
	var dialect string
	var url string

	if os.Getenv("GO_ENV") == "development" {
		dialect = "sqlite3"
		url = "devheaven.db"
	} else {
		dialect = "postgres"
		url = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USERNAME"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
	}

	return gorm.Open(dialect, url)
}
