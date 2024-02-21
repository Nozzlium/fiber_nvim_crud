package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/fiber_nvim_crud/requestparam"
	"github.com/nozzlium/fiber_nvim_crud/services"
)

type ContactControllerImpl struct {
	ContactService services.ContactService
}

func (controller *ContactControllerImpl) Create(ctx *fiber.Ctx) error {
	var contact requestparam.Contact
	err := ctx.BodyParser(&contact)
	if err != nil {
		return err
	}

	resp, err := controller.ContactService.Create(ctx.Context(), contact)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   resp,
	}, "application/json")
}

func (controller *ContactControllerImpl) Find(ctx *fiber.Ctx) error {
	ctx.QueryInt
}

func (controller *ContactControllerImpl) FindById(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *ContactControllerImpl) Update(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *ContactControllerImpl) Delete(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
