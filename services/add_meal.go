package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
	"github.com/simesaba80/kcl-back/db"
)

func AddMealData(c echo.Context) error {
	// AddMeal is a function that adds a meal data to the database.
	type body struct {
		UserID         string `json:"user_id"`
		MealName       string `json:"meal_name"`
		CaloriePer100G int    `json:"calorie_per_100g"`
		Date           string `json:"date"`
	}
	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return c.String(400, "Bad Request")
	}
	ctx := c.Request().Context()
	meal := db.Meal{
		UserID:         obj.UserID,
		MealName:       obj.MealName,
		CaloriePer100G: obj.CaloriePer100G,
		Date:           obj.Date,
	}
	if meal.UserID == "" || meal.MealName == "" {
		return echo.NewHTTPError(400, "Bad Request")
	}
	result := cruds.CreateMealData(ctx, meal)
	if result.ID == 0 {
		return echo.NewHTTPError(500, "Internal Server Error")
	}
	fmt.Println(result)
	return c.String(201, "AddMealData")
}
