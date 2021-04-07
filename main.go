package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/safecornerscoffee/golang-crud-api/internal/config"
	"github.com/safecornerscoffee/golang-crud-api/internal/database"
	"github.com/safecornerscoffee/golang-crud-api/user"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", func(c echo.Context) (err error) {
		return c.String(http.StatusOK, "healthy")
	})

	db, err := database.NewDB(config.NewConfig())
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	user.RegisterRoutes(db, e)

	e.Logger.Fatal(e.Start(":8080"))
}
