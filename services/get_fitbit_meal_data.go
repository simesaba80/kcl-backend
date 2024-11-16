package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func GetUserMealData(c echo.Context) error {
	fitbitUserID := c.Param("fitbit_user_id")
	fmt.Println("fitbitUserID:" + fitbitUserID)
	ctx := c.Request().Context()
	user, err := cruds.GetUserByFitbitUserID(fitbitUserID, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	meal, err := cruds.GetUserMealData(user.UID, ctx)
	if err != nil {
		return c.String(404, "Not Found")
	}
	fmt.Println(meal)
	type response struct {
		Calories      int     `json:"calories"`
		Protein       float64 `json:"protein"`
		Fat           float64 `json:"fat"`
		Carbohydrates float64 `json:"carbohydrates"`
		Salt          float64 `json:"salt"`
		Calcium       float64 `json:"calcium"`
		Date          string  `json:"date"`
	}
	responseData := response{}
	for i := 0; i < len(meal); i++ {
		responseData.Calories += meal[i].Calories
		responseData.Protein += meal[i].Protein
		responseData.Fat += meal[i].Fat
		responseData.Carbohydrates += meal[i].Carbohydrates
		responseData.Salt += meal[i].Salt
		responseData.Calcium += meal[i].Calcium
	}
	responseData.Date = meal[0].Date
	return c.JSON(200, responseData)
}
