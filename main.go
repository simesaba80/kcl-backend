package main

import (
	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/services"
	"github.com/simesaba80/kcl-back/utils"
)

func main() {
	e := echo.New()
	utils.LoadConfig()
	utils.Connectdb()

	e.GET("/", services.Hello)

	e.Start(":8080")
}
