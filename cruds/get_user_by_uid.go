package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func GetUserByUID(uid string, ctx context.Context) (db.User, error) {
	user := db.User{}
	if err := db.DB.NewSelect().Model(&user).Where("uid = ?", uid).Scan(ctx); err != nil {
		return user, err
	}
	return user, nil
}
