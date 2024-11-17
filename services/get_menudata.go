package services

import (
	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/cruds"
)

func GetMenuData(c echo.Context) error {
	menuName := c.Param("menu_name")
	menus, err := cruds.SearchMenu(menuName, c.Request().Context())
	if err != nil {
		return c.String(500, "Failed to get menu data")
	}
	return c.JSON(200, menus)
}
