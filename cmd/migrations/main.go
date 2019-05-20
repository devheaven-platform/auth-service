package main

import (
	"github.com/devheaven-platform/auth-service/pkg/domain"
	"github.com/devheaven-platform/auth-service/pkg/utils/db"
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
	db.AutoMigrate(&domain.Email{}, &domain.Role{}, &domain.User{})

	// Create user
	user := domain.User{
		Firstname: "User",
		Lastname:  "",
		Emails: []domain.Email{
			{Email: "user@devheaven.nl"},
		},
		Roles: []domain.Role{
			{Role: "USER"},
		},
		Password: "Test1234",
		Enabled:  true,
	}

	// Create developer
	developer := domain.User{
		Firstname: "Developer",
		Lastname:  "",
		Emails: []domain.Email{
			{Email: "developer@devheaven.nl"},
		},
		Roles: []domain.Role{
			{Role: "USER"},
		},
		Password: "Test1234",
		Enabled:  true,
	}

	// Create hr
	hr := domain.User{
		Firstname: "HR",
		Lastname:  "",
		Emails: []domain.Email{
			{Email: "hr@devheaven.nl"},
		},
		Roles: []domain.Role{
			{Role: "USER"},
		},
		Password: "Test1234",
		Enabled:  true,
	}

	// Create manager
	manager := domain.User{
		Firstname: "Manager",
		Lastname:  "",
		Emails: []domain.Email{
			{Email: "manager@devheaven.nl"},
		},
		Roles: []domain.Role{
			{Role: "USER"},
		},
		Password: "Test1234",
		Enabled:  true,
	}

	db.Save(&user)
	db.Save(&developer)
	db.Save(&hr)
	db.Save(&manager)
}
