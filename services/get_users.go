package services

import (
	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func GetUsers(c echo.Context) error {
	users, err := cruds.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.String(500, "Internal Server Error")
	}
	return c.JSON(200, users)
}
