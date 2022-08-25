package main

import (
	"fmt"
	"github.com/Somphoph/go-logistic/pkg/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start. err: %v", err))
	}
}

func run() error {
	app := fiber.New(fiber.Config{StrictRouting: true, AppName: "Logistic"})
	app.Route("/", http.Routes, "/")
	return app.Listen(":3000")
}
