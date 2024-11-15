package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/simesaba80/kcl-back/db"
	"github.com/simesaba80/kcl-back/services"
	"github.com/simesaba80/kcl-back/utils"
)

func main() {
	e := echo.New()
	utils.LoadConfig()
	db.Connectdb()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/", services.Hello)
	e.GET("/user", services.GetUsers)
	e.GET("/login/:uid", services.Login)
	e.POST("/user/create", services.CreateUser)
	e.PUT("/user/update/:id", services.UpdateUser)
	e.GET("/momentum/:id", services.GetUserMomentumData)
	e.GET("/sleep/:id", services.GetUserSleepData)
	e.POST("/meal/create", services.AddMealData)
	e.GET("/meal/:id", services.GetUserMealData)

	e.Start(":8080")
}
