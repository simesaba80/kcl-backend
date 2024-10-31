package main

import (
	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/services"
)

func main() {
	e := echo.New()

	e.GET("/", services.Hello)

	e.Start(":8080")
}
