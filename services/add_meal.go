package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
	"github.com/simesaba80/kcl-back/db"
)

// json入れ子の場合は構造体も入れ子に
func AddMealData(c echo.Context) error {
	// AddMeal is a function that adds a meal data to the database.
	type body struct {
		UserID        string  `json:"user_id"`
		MealName      string  `json:"meal_name"`
		Calories      int     `json:"calories"`
		Protein       float64 `json:"protein"`
		Fat           float64 `json:"fat"`
		Carbohydrates float64 `json:"carbohydrates"`
		Salt          float64 `json:"salt"`
		Date          string  `json:"date"`
	}
	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return c.String(400, "Bad Request")
	}
	ctx := c.Request().Context()
	meal := db.Meal{
		UserID:        obj.UserID,
		MealName:      obj.MealName,
		Calories:      obj.Calories,
		Protein:       obj.Protein,
		Fat:           obj.Fat,
		Carbohydrates: obj.Carbohydrates,
		Salt:          obj.Salt,
		Date:          obj.Date,
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
