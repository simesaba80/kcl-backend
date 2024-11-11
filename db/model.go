package db

import (
	"time"
)

// User is a struct that represents a user.
type User struct {
	ID        int       `bun:"id,pk,autoincrement"`
	UID       string    `bun:"uid,unique,notnull"`
	Name      string    `bun:"name,notnull"`
	Sex       string    `bun:"sex,notnull"`
	Height    int       `bun:"height,nullzero"`
	Weight    int       `bun:"weight,nullzero"`
	Age       int       `bun:"age,nullzero"`
	Job       string    `bun:"job,nullzero"`
	CreatedAt time.Time `bun:"created_at"`
}

type Momentum struct {
	// Momentum is a struct that represents a user's Momentum data.
	ID           int    `bun:"id,pk,autoincrement"`
	UserID       int    `bun:"user_id,notnull"`
	Steps        int    `bun:"steps,notnull"`
	Calories     int    `bun:"calories,notnull"`
	Distance     int    `bun:"distance,notnull"`
	MaxHeartRate int    `bun:"max_heart_rate,notnull"`
	MinHeartRate int    `bun:"min_heart_rate,notnull"`
	Date         string `bun:"date,notnull"`
}

type Sleep struct {
	// Sleep is a struct that represents a user's sleep data.
	ID         int       `bun:"id, pk, autoincrement"`
	UserID     int       `bun:"user_id,notnull"`
	Hours      int       `bun:"hours,notnull"`
	StartedAt  time.Time `bun:"started_at,notnull"`
	EndedAt    time.Time `bun:"ended_at,notnull"`
	DeepSleep  int       `bun:"deep_sleep,notnull"`
	LightSleep int       `bun:"light_sleep,notnull"`
	RemSleep   int       `bun:"rem_sleep,notnull"`
	Wake       int       `bun:"wake,notnull"`
	Date       string    `bun:"date,notnull"`
}

type Meal struct {
	// Meal is a struct that represents a user's meal data.
	ID             int    `bun:"id, pk, autoincrement"`
	UserID         int    `bun:"user_id,notnull"`
	MealName       string `bun:"meal_name,notnull"`
	CaloriePer100G int    `bun:"calorie_per_100g,notnull"`
	Date           time.Time
}
