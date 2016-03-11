package main

import (
	"fmt"
	"io"
	"os"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request().(*standard.Request)
		req.ParseMultipartForm(16 << 20) // Max memory 16 MiB

		// Read form fields
		name := c.Form("name")
		email := c.Form("email")

		// Read files
		files := req.MultipartForm.File["files"]
		for _, f := range files {
			// Source file
			src, err := f.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			// Destination file
			dst, err := os.Create(f.Filename)
			if err != nil {
				return err
			}
			defer dst.Close()

			if _, err = io.Copy(dst, src); err != nil {
				return err
			}
		}
		return c.String(http.StatusOK, fmt.Sprintf("Thank You! %s <%s>, %d files uploaded successfully.",
			name, email, len(files)))
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))

	e.Post("/upload", upload())

	e.Run(standard.New(":1323"))
}
