package db

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"TokaAPI/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	driver := os.Getenv("DB_DRIVER") // "mssql" del env

	var database *gorm.DB
	var err error

	switch driver {
	case "sqlite":
		dsn := os.Getenv("DB_FILE")
		if dsn == "" {
			dsn = "data.db"
		}
		database, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default: // mssql del env
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		name := os.Getenv("DB_NAME")
		encrypt := os.Getenv("DB_ENCRYPT")       // "disable" porquie se utilizo docker
		trust := os.Getenv("DB_TRUSTSERVERCERT") // "true" en Docker
		if host == "" {
			host = "mssql"
		}
		if port == "" {
			port = "1433"
		}

		// contrase;a con caracteres especiales
		u := &url.URL{
			Scheme: "sqlserver",
			User:   url.UserPassword(user, pass),
			Host:   fmt.Sprintf("%s:%s", host, port),
		}
		q := u.Query()
		if name != "" {
			q.Set("database", name)
		}
		if encrypt != "" {
			q.Set("encrypt", encrypt)
		}
		if trust != "" {
			q.Set("trustservercertificate", trust)
		}
		u.RawQuery = q.Encode()

		database, err = gorm.Open(sqlserver.Open(u.String()), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("db no conexion error: %v", err)
	}

	if err := database.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalf("migrations fails: %v", err)
	}

	log.Println("migrations jalandooo")

	return database
}
