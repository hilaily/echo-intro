package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Setup
	e := echo.New()
	e.Use(middleware.Static("public"))

	e.Get("/jsonp", echo.HandlerFunc(func(c echo.Context) error {
		callback := c.Query("callback")
		var content struct {
			Response  string    `json:"response"`
			Timestamp time.Time `json:"timestamp"`
			Random    int       `json:"random"`
		}
		content.Response = "Sent via JSONP"
		content.Timestamp = time.Now().UTC()
		content.Random = rand.Intn(1000)
		return c.JSONP(http.StatusOK, callback, &content)
	}))

	// Start server
	e.Run(standard.New(":1323"))
}
