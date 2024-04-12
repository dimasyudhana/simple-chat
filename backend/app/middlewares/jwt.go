package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dimasyudhana/simple-chat/app/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := ValidateToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func ValidateToken(c *gin.Context) error {
	tokenString, err := ExtractToken(c)
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWT), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = claims["authorized"]
		_ = claims["userID"]
		return nil
	}

	return errors.New("Invalid token provided")
}

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT))
}

func ExtractToken(c *gin.Context) (string, error) {
	user, exists := c.Get("user")
	if exists {
		token := user.(*jwt.Token)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			userID := claims["userID"].(string)
			return userID, nil
		}
	}

	return "", errors.New("failed to extract jwt-token")
}
