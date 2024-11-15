package services

import "github.com/labstack/echo/v4"

func HelloWithAuth(c echo.Context) error {
	uid := c.Get("uid").(string)

	return c.String(200, "Hello, World!"+uid)
}
