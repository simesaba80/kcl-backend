package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func UpdateUser(c echo.Context) error {
	type body struct {
		Name   string `json:"name"`
		Sex    string `json:"sex"`
		Height int    `json:"height"`
		Weight int    `json:"weight"`
		Age    int    `json:"age"`
		Job    string `json:"job"`
	}
	id := c.Param("id")
	ctx := c.Request().Context()
	user, err := cruds.GetUserByID(id, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return c.String(400, "Bad Request")
	}
	user.Name = obj.Name
	user.Sex = obj.Sex
	user.Height = obj.Height
	user.Weight = obj.Weight
	user.Age = obj.Age
	user.Job = obj.Job
	result, err := cruds.UpdateUser(id, user, ctx)
	if err != nil {
		return c.String(500, "Internal Server Error")
	}
	fmt.Println(result)
	return c.String(200, "UpdateUser")
}
