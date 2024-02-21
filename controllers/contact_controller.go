package controllers

import "github.com/gofiber/fiber/v2"

type ContactController interface {
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
