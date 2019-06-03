package main

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/devheaven-platform/auth-service/pkg/utils/db"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

// main is invoked by the go compiler and is used
// to migrate the database.
func main() {
	db, err := db.OpenConnection()
	db.LogMode(true)

	if err != nil {
		log.WithError(err).Fatal("An error occurred while connecting to the database")
	}
	defer db.Close()

	// Create models
	db.DropTableIfExists(&domain.User{}, &domain.Email{}, &domain.Role{})
	db.CreateTable(&domain.Email{}, &domain.Role{}, &domain.User{})

	// Create user
	id, _ := uuid.NewRandom()
	user := domain.User{
		ID: id,
		Emails: []domain.Email{
			{Email: "user@devheaven.nl"},
		},
		Roles: []domain.Role{
			{Role: "ROLE_USER"},
		},
		Password: "Test1234",
	}

	// Create developer
	id, _ = uuid.NewRandom()
	developer := domain.User{
		ID: id,
		Emails: []domain.Email{
			{Email: "developer@devheaven.nl"},
		},
		Roles: []domain.Role{
			{Role: "ROLE_USER"},
			{Role: "ROLE_DEVELOPER"},
		},
		Password: "Test1234",
	}

	// Create hr
	id, _ = uuid.NewRandom()
	hr := domain.User{
		ID: id,
		Emails: []domain.Email{
			{Email: "hr@devheaven.nl"},
		},
		Roles: []domain.Role{
			{Role: "ROLE_USER"},
			{Role: "ROLE_HR"},
		},
		Password: "Test1234",
	}

	// Create manager
	id, _ = uuid.NewRandom()
	manager := domain.User{
		ID: id,
		Emails: []domain.Email{
			{Email: "manager@devheaven.nl"},
			{Email: "devheavenplatform@gmail.com"},
		},
		Roles: []domain.Role{
			{Role: "ROLE_USER"},
			{Role: "ROLE_MANAGER"},
		},
		Password: "Test1234",
	}

	db.Create(&user)
	db.Create(&developer)
	db.Create(&hr)
	db.Create(&manager)
}
