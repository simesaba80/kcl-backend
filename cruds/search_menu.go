package cruds

import (
	"context"

	"github.com/simesaba80/kcl-back/db"
)

func SearchMenu(mealName string, ctx context.Context) ([]db.Menu, error) {
	menus := []db.Menu{}
	if err := db.DB.NewSelect().Model(&menus).Where("meal_name LIKE ?", "%"+mealName+"%").Scan(ctx); err != nil {
		return menus, err
	}
	return menus, nil
}
