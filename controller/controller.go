package controller

import "github.com/gofiber/fiber/v3"

type Controller interface {
	Init(app *fiber.App)
}
