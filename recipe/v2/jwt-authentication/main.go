package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

const (
	Bearer     = "Bearer"
	SigningKey = "somethingsupersecret"
)

// A JSON Web Token middleware
func JWTAuth(key string) echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {

			// Skip WebSocket
			if (c.Request().Header().Get(echo.Upgrade)) == echo.WebSocket {
				return nil
			}

			auth := c.Request().Header().Get("Authorization")
			l := len(Bearer)
			he := echo.NewHTTPError(http.StatusUnauthorized)

			if len(auth) > l+1 && auth[:l] == Bearer {
				t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {

					// Always check the signing method
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// Return the key for validation
					return []byte(key), nil
				})
				if err == nil && t.Valid {
					// Store token claims in echo.Context
					c.Set("claims", t.Claims)
					return nil
				}
			}
			return he
		})
	}
}

func accessible() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "No auth required for this route.\n")
	}

}
func restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Access granted with JWT.\n")
	}
}

func main() {
	// Echo instance
	e := echo.New()

	// Logger
	e.Use(middleware.Logger())

	// Unauthenticated route
	e.Get("/", accessible())

	// Restricted group
	r := e.Group("/restricted")
	r.Use(JWTAuth(SigningKey))
	r.Get("", restricted())

	// Start server
	e.Run(standard.New(":1323"))
}
