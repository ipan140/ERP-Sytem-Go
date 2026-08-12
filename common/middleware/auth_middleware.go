package middleware

import (
	"net/http"
	"strings"

	"ERP-System/common/utils"
	"ERP-System/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return utils.SendError(c, http.StatusUnauthorized, "Missing Authorization header", "")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return utils.SendError(c, http.StatusUnauthorized, "Invalid Authorization format", "")
			}

			tokenString := parts[1]
			secret := config.AppConfig.JWTSecret

			token, err := jwt.ParseWithClaims(tokenString, &utils.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return utils.SendError(c, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			}

			if claims, ok := token.Claims.(*utils.JWTCustomClaims); ok && token.Valid {
				c.Set("user_id", claims.UserID)
				c.Set("company_id", claims.CompanyID)
				c.Set("role", claims.Role)
			}

			return next(c)
		}
	}
}
