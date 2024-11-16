package cruds

import (
	"context"
	"fmt"

	"github.com/simesaba80/kcl-back/db"
)

func GetUserByFitbitUserID(fitbitUserID string, ctx context.Context) (db.User, error) {
	user := db.User{}
	fmt.Println("fitbitUserID:" + fitbitUserID)
	if err := db.DB.NewSelect().Model(&user).Where("fitbit_user_id = ?", fitbitUserID).Scan(ctx); err != nil {
		return user, err
	}
	return user, nil
}
