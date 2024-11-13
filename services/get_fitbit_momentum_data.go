package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func GetUserMomentumData(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	user, err := cruds.GetUserByID(id, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	momentums, err := cruds.GetUserMomentumData(user.UID, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	fmt.Println(momentums)
	return c.String(200, "GetUserMomentumData")
}
