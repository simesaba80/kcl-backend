package cruds

import (
	"context"
	"time"

	"github.com/simesaba80/kcl-back/database"
)

func AddUserToDB(user database.User, ctx context.Context) database.User {
	user.CreatedAt = time.Now()
	database.DB.NewInsert().Model(&user).Exec(ctx)
	return user

}
