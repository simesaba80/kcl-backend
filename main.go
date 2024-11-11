package main

import (
	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/database"
	"github.com/simesaba80/kcl-back/services"
	"github.com/simesaba80/kcl-back/utils"
)

func main() {
	e := echo.New()
	utils.LoadConfig()
	database.Connectdb()

	e.GET("/", services.Hello)
	e.GET("/user", services.GetUsers)
	e.GET("/login/:uid", services.Login)
	e.POST("/user/create", services.CreateUser)

	e.Start(":8080")
}
