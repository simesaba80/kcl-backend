package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func GetUserActivityData(id string, ctx context.Context) (db.Activities, error) {
	activity := db.Activities{}
	if err := db.DB.NewSelect().Model(&activity).Where("user_id = ?", id).Order("created_at desc").Limit(1).Scan(ctx); err != nil {
		return activity, nil
	}
	return activity, nil
}
