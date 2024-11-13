package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func GetUserMomentumData(id string, ctx context.Context) ([]db.Momentum, error) {
	momentum := []db.Momentum{}
	if err := db.DB.NewSelect().Model(&momentum).Where("user_id = ?", id).Order("created_at asc").Scan(ctx); err != nil {
		return momentum, nil
	}
	return momentum, nil
}
