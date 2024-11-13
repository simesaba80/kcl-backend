package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func UpdateUser(id string, user db.User, ctx context.Context) (db.User, error) {
	if _, err := db.DB.NewUpdate().Model(&user).Where("id = ?", id).Exec(ctx); err != nil {
		return user, err
	}
	return user, nil
}
