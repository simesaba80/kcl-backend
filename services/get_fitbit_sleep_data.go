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
	type response struct {
		Minutes    int    `json:"minutes"`
		DeepSleep  int    `json:"deep_sleep"`
		LightSleep int    `json:"light_sleep"`
		RemSleep   int    `json:"rem_sleep"`
		Wake       int    `json:"wake"`
		Date       string `json:"date"`
	}
	return c.JSON(200, response{
		Minutes:    sleep.Minutes,
		DeepSleep:  sleep.DeepSleep,
		LightSleep: sleep.LightSleep,
		RemSleep:   sleep.RemSleep,
		Wake:       sleep.Wake,
		Date:       sleep.Date,
	})
}
