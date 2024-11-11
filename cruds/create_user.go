package cruds

import (
	"context"
	"time"

	"github.com/simesaba80/kcl-back/db"
)

func AddUserToDB(user db.User, ctx context.Context) db.User {
	user.CreatedAt = time.Now()
	db.DB.NewInsert().Model(&user).Exec(ctx)
	return user

}
