package helper

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	config "go-project/configs"
	r "go-project/repositories"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"iat": time.Now().Unix(),
	})
	return token.SignedString(privateKey)
}

func ExtractTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	split := strings.Split(bearerToken, " ")
	if len(split) == 2 {
		return split[1]
	}
	return ""
}

func ExtractToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}

func ValidateToken(c *gin.Context) error {
	token, err := ExtractToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

func ValidUser(c *gin.Context) error {
	if err := ValidateToken(c); err != nil {
		return err
	}
	token, _ := ExtractToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	userRepo := r.NewUserRepository(config.DB)
	_, err := userRepo.FindOne(strconv.Itoa(int(userId)))
	if err != nil {
		return errors.New("unauthorized")
	}
	return nil
}
