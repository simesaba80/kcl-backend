package cruds

import (
	"context"
	"time"

	"github.com/simesaba80/kcl-back/db"
)

func CreateMealData(ctx context.Context, meal db.Meal) db.Meal {
	// CreateMealData is a function that creates a meal data in the database.
	meal.CreatedAt = time.Now()
	db.DB.NewInsert().Model(&meal).Exec(ctx)
	return meal
}
