package middleware

import (
	"banner-service/pkg/config"
	"banner-service/pkg/models"
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

// JWTAuthentication проверяет наличие и валидность JWT в запросе
func JWTAuthentication(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			notAuth := []string{"/signin", "/user_banner"} // Эндпоинты, не требующие аутентификации
			requestPath := r.URL.Path

			// Пропускаем запросы без необходимости аутентификации
			for _, value := range notAuth {
				if value == requestPath {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Проверка наличия и формата токена
			tokenHeader := r.Header.Get("Authorization")
			if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
				http.Error(w, "Missing auth token", http.StatusUnauthorized)
				return
			}

			tokenPart := strings.Split(tokenHeader, " ")[1] // Токен
			claims := &models.Claims{}

			token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JWTSecret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Malformed authentication token", http.StatusForbidden)
				return
			}

			// Проверка роли для админских эндпоинтов
			if strings.Contains(requestPath, "/admin/") && claims.Role != "admin" {
				http.Error(w, "Forbidden - Admins only", http.StatusForbidden)
				return
			}

			// Добавляем информацию о пользователе в контекст запроса
			ctx := context.WithValue(r.Context(), "user", claims)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
