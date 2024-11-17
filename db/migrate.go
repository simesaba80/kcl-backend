package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func Insert(c echo.Context) error {
	// Read the JSON file
	file, err := os.Open("/app/db/foods_data.json")
	if err != nil {
		log.Fatalf("failed to open JSON file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read JSON file: %v", err)
	}

	// Unmarshal the JSON data
	var menus []Menu
	err = json.Unmarshal(byteValue, &menus)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON data: %v", err)
	}

	// Insert the data into the menu table
	for _, menu := range menus {
		_, err := DB.NewInsert().Model(&menu).Exec(context.Background())
		if err != nil {
			log.Fatalf("failed to insert menu data: %v", err)
		}
	}

	fmt.Println("Data successfully inserted into the menu table")
	return c.String(201, "Data successfully inserted into the menu table")
}
