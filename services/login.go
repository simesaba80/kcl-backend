package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func Login(c echo.Context) error {
	uid := c.Param("uid")
	ctx := c.Request().Context()
	result, err := cruds.GetUserByUID(uid, ctx)
	if err != nil {
		fmt.Println(err)
		return c.String(500, "Internal Server Error")
	}
	fmt.Println(result)
	return c.String(200, "Login")
}
