package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func GetUserActivitiesData(id string, ctx context.Context) ([]db.Activities, error) {
	activities := []db.Activities{}
	if err := db.DB.NewSelect().Model(&activities).Where("user_id = ?", id).Order("created_at asc").Scan(ctx); err != nil {
		return activities, nil
	}
	return activities, nil
}
