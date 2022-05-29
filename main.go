package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	portStr := mustEnvVar("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("failed to parse PORT env var as int; %s; %v", portStr, err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func mustEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("you must define env var: %s", key)
	}
	return val
}
