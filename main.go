package main

import (
	"go-react-app/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	web.RegisterHandlers(app)

	app.GET("/api/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Herman")
	})

	app.Logger.Fatal(app.Start(":3000"))
}
