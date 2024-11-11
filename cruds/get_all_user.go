package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func GetAllUsers(ctx context.Context) ([]db.User, error) {
	users := []db.User{}
	if err := db.DB.NewSelect().Model(&users).Scan(ctx); err != nil {
		return users, err
	}
	return users, nil
}
