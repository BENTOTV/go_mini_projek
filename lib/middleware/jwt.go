package middleware

import (
	"myapp/config"
	models "myapp/model"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	//"github.com/labstack/echo/v4"
)

func GenerateTokenUser(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

func GenerateTokenAdmin(admin models.Admin) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"admin_id":   admin.ID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}
