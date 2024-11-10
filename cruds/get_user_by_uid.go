package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/database"
)

func GetUserByUID(uid string, ctx context.Context) (database.User, error) {
	user := database.User{}
	if err := database.DB.NewSelect().Model(&user).Where("uid = ?", uid).Scan(ctx); err != nil {
		return user, err
	}
	return user, nil
}
