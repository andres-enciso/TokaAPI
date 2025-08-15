package auth

import (
	"encoding/base64"
	"net/http"
	"strings"

	"TokaAPI/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Basic ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(h, "Basic "))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		parts := strings.SplitN(string(raw), ":", 2)
		if len(parts) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		username, password := parts[0], parts[1]

		var u models.User
		if err := db.Where("username = ?", username).First(&u).Error; err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", u.Username)
		c.Next()
	}
}
