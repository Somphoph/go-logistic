package http

import "github.com/gofiber/fiber/v2"

func Routes(api fiber.Router) {
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World.")
	})
	api.Get("/xo", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OX")
	})
}
