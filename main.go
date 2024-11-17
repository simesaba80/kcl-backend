package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/simesaba80/kcl-back/db"
	"github.com/simesaba80/kcl-back/originalmiddleware"
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
	e.PUT("/user/update/:fitbit_user_id", services.UpdateUser)
	e.GET("/activities/:fitbit_user_id", services.GetUserActivitiesData)
	e.GET("/sleep/:fitbit_user_id", services.GetUserSleepData)
	e.POST("/meal/create", services.AddMealData)
	e.GET("/meal/:fitbit_user_id", services.GetUserMealData)
	e.GET("/insert", db.Insert)
	e.GET("/getmenudata/:menu_name", services.GetMenuData)

	r := e.Group("/auth")                            // 認証が必要なエンドポイントはすべて/auth以下にまとめる
	r.Use(originalmiddleware.FirebaseAuthMiddleware) // 認証ミドルウェアを適用
	r.GET("/hello", services.HelloWithAuth)

	e.Start(":8080")
}
