package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	crud "github.com/simesaba80/kcl-back/cruds"
	"github.com/simesaba80/kcl-back/database"
)

func CreateUser(c echo.Context) error {
	type body struct {
		UID    string `json:"uid"`
		Name   string `json:"name"`
		Sex    string `json:"sex"`
		Email  string `json:"email"`
		Height int    `json:"height"`
		Age    int    `json:"age"`
		Job    string `json:"job"`
	}
	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return c.String(400, "Bad Request")
	}
	ctx := c.Request().Context()
	user := database.User{
		UID:    obj.UID,
		Name:   obj.Name,
		Sex:    obj.Sex,
		Email:  obj.Email,
		Height: obj.Height,
		Age:    obj.Age,
		Job:    obj.Job,
	}
	fmt.Println(user)
	if user.UID == "" || user.Name == "" {
		return echo.NewHTTPError(400, "Bad Request")
	}
	result := crud.AddUserToDB(user, ctx)
	if result.ID == 0 {
		return c.String(500, "Internal Server Error")
	}
	fmt.Println(result)
	return c.String(200, "CreateUser")
}
