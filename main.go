package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Сервер работает")

	s := echo.New()

	s.Use(MW)

	s.GET("/status", Handler)

	err := s.Start(":8080")

	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {
	d := time.Date(2026, time.April, 28, 0, 0, 0, 0, time.UTC)

	durr := time.Until(d)

	s := fmt.Sprintf("Кол-во дней: %d", int64(durr.Hours())/24)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}

func MW(next echo.HandlerFunc) echo.HandlerFunc {
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
