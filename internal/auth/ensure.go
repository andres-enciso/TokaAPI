package auth

import (
	"log"

	"TokaAPI/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func EnsureAdminUser(db *gorm.DB, username, password string) {
	if username == "" || password == "" {
		log.Println("ADMIN_USER/ADMIN_PASS not set")
		return
	}

	var count int64
	if err := db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		log.Fatalf(" admin error: %v", err)
	}
	if count > 0 {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("bcrypt error: %v", err)
	}

	if err := db.Create(&models.User{
		Username: username,
		Password: string(hash),
	}).Error; err != nil {
		log.Fatalf("erorr al crear admin error: %v", err)
	}
	log.Printf("admin user '%s' created", username)
}
