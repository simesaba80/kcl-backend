package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func GetUserMealData(uid string, ctx context.Context) ([]db.Meal, error) {
	meal := []db.Meal{}
	if err := db.DB.NewSelect().Model(&meal).Where("user_id = ?", uid).Order("created_at asc").Scan(ctx); err != nil {
		return meal, nil
	}
	return meal, nil
}
