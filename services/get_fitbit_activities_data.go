package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func GetUserActivitiesData(c echo.Context) error {
	fitbitUserID := c.Param("fitbit_user_id")
	ctx := c.Request().Context()
	user, err := cruds.GetUserByFitbitUserID(fitbitUserID, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	Activity, err := cruds.GetUserActivityData(user.UID, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	fmt.Println(Activity)
	type response struct {
		Steps    int    `json:"steps"`
		Calories int    `json:"calories"`
		Date     string `json:"date"`
	}
	return c.JSON(200, response{
		Steps:    Activity.Steps,
		Calories: Activity.Calories,
		Date:     Activity.Date,
	})
}
