package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
	"github.com/simesaba80/kcl-back/db"
)

func CreateUser(c echo.Context) error {
	type body struct {
		UID          string `json:"uid"`
		Name         string `json:"name"`
		FitbitUserID string `json:"fitbit_user_id"`
		Sex          string `json:"sex"`
		Height       int    `json:"height"`
		Weight       int    `json:"weight"`
		Age          int    `json:"age"`
		Job          string `json:"job"`
		AcessToken   string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return c.String(400, "Bad Request")
	}
	ctx := c.Request().Context()
	user := db.User{
		UID:          obj.UID,
		Name:         obj.Name,
		FitbitUserID: obj.FitbitUserID,
		Sex:          obj.Sex,
		Height:       obj.Height,
		Weight:       obj.Weight,
		Age:          obj.Age,
		Job:          obj.Job,
		AccessToken:  obj.AcessToken,
		RefreshToken: obj.RefreshToken,
	}
	fmt.Println(user)
	if user.UID == "" || user.Name == "" {
		return echo.NewHTTPError(400, "Bad Request")
	}
	result := cruds.AddUserToDB(user, ctx)
	if result.ID == 0 {
		return c.String(500, "Internal Server Error")
	}
	fmt.Println(result)
	return c.String(200, "CreateUser")
}
