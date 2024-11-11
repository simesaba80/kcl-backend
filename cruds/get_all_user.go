package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/database"
)

func GetAllUsers(ctx context.Context) ([]database.User, error) {
	users := []database.User{}
	if err := database.DB.NewSelect().Model(&users).Scan(ctx); err != nil {
		return users, err
	}
	return users, nil
}
