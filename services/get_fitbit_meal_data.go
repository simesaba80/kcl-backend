package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func GetUserMealData(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	user, err := cruds.GetUserByID(id, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	meal, err := cruds.GetUserMealData(user.UID, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	fmt.Println(meal)
	return c.String(200, "GetUserMealData")
}
