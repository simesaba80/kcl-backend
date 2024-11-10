package database

import (
	"time"
)

// User is a struct that represents a user.
type User struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	Height    int       `json:"height"`
	Age       int       `json:"age"`
	Job       string    `json:"job"`
	CreatedAt time.Time `json:"created_at"`
}

type Momentum struct {
	// Momentum is a struct that represents a user's Momentum data.
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	Steps        int    `json:"steps"`
	Calories     int    `json:"calories"`
	Distance     int    `json:"distance"`
	MaxHeartRate int    `json:"max_heart_rate"`
	MinHeartRate int    `json:"min_heart_rate"`
	Date         string `json:"date"`
}

type Sleep struct {
	// Sleep is a struct that represents a user's sleep data.
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Hours      int       `json:"hours"`
	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
	DeepSleep  int       `json:"deep_sleep"`
	LightSleep int       `json:"light_sleep"`
	RemSleep   int       `json:"rem_sleep"`
	Wake       int       `json:"wake"`
	Date       string    `json:"date"`
}

type Meal struct {
	// Meal is a struct that represents a user's meal data.
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	MealName       string    `json:"meal_name"`
	CaloriePer100G int       `json:"calories_per_100g"`
	Date           time.Time `json:"date"`
}
