package main

import (
	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/db"
	"github.com/simesaba80/kcl-back/services"
	"github.com/simesaba80/kcl-back/utils"
)

func main() {
	e := echo.New()
	utils.LoadConfig()
	db.Connectdb()

	e.GET("/", services.Hello)
	e.GET("/user", services.GetUsers)
	e.GET("/login/:uid", services.Login)
	e.POST("/user/create", services.CreateUser)
	e.PUT("/user/update/:id", services.UpdateUser)
	e.GET("/momentum/:id", services.GetUserMomentumData)

	e.Start(":8080")
}
