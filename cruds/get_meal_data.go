package cruds

import (
	"context"
	"time"

	"github.com/simesaba80/kcl-back/db"
)

func GetUserMealData(uid string, ctx context.Context) ([]db.Meal, error) {
	meal := []db.Meal{}
	date := time.Now().Format("2006-01-02")
	if err := db.DB.NewSelect().Model(&meal).Where("user_id = ?", uid).Where("date = ?", date).Order("created_at asc").Scan(ctx); err != nil {
		return meal, nil
	}
	return meal, nil
}
