package middleware

import (
	"net/http"
	"strings"
	"ticket-booking-system/internal/auth"
	"ticket-booking-system/internal/httpresponse"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(jwtService auth.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, httpresponse.Response{
					Success: false,
					Message: "Unauthorized: Missing authorization header",
				})
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, httpresponse.Response{
					Success: false,
					Message: "Unauthorized: Invalid authorization header",
				})
			}

			token := parts[1]

			claims, err := jwtService.ValidateToken(token)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, httpresponse.Response{
					Success: false,
					Message: "Unauthorized: token is invalid or expired",
				})
			}

			c.Set("user_id", claims.ID)
			c.Set("user_name", claims.Name)
			c.Set("user_email", claims.Email)

			return next(c)
		}
	}
}
