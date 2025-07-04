package mw

import (
	"log"

	"github.com/labstack/echo/v4"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		val := ctx.Request().Header.Get("User-Role")

		if val == "admin" {
			log.Println("red ALERT")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
