package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
	"time"
	"test/internal/config"
)

func ExtractToken(authHeader string) string {
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}

func ValidateToken(tokenString string) error {
	if tokenString == "" {
		return errors.New("token string is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < jwt.TimeFunc().Unix() {
				return errors.New("token expired")
			}
		}
		return nil
	}

	return errors.New("invalid token")
}

func GenerateJWT(userID int) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   strconv.Itoa(userID),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}