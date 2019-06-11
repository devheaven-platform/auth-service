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
	id, _ := uuid.Parse("8c5df3bc-9fa6-4d73-b79a-9a1cbb35740c")
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
	id, _ = uuid.Parse("b0203081-5dfe-4bb7-87d1-e2c59e2af7b6")
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
	id, _ = uuid.Parse("6b59c645-82c3-4e08-b089-f4236a2141b6")
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
	id, _ = uuid.Parse("75129bb5-5c12-48a1-8410-bb2630fff9ed")
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
