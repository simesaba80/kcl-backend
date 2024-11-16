package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func GetUserSleepData(c echo.Context) error {
	fitbitUserID := c.Param("fitbit_user_id")
	ctx := c.Request().Context()
	user, err := cruds.GetUserByFitbitUserID(fitbitUserID, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	sleep, err := cruds.GetUserSleepData(user.UID, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	fmt.Println(sleep)
	return c.String(200, "GetUserSleepData")
}
