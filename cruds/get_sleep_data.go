package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func GetUserSleepData(uid string, ctx context.Context) (db.Sleep, error) {
	sleep := db.Sleep{}
	if err := db.DB.NewSelect().Model(&sleep).Where("user_id = ?", uid).Order("created_at desc").Limit(1).Scan(ctx); err != nil {
		return sleep, nil
	}
	return sleep, nil
}
