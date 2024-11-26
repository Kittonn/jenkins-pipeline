package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kittonn/github-actions-ghcr/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	serverUrl := fmt.Sprintf(":%d", config.Port)

	e.Logger.Fatal(e.Start(serverUrl))
}
