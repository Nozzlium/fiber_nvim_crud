package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/fiber_nvim_crud/entities"
	"github.com/nozzlium/fiber_nvim_crud/params"
	"github.com/nozzlium/fiber_nvim_crud/requestbody"
	"github.com/nozzlium/fiber_nvim_crud/services"
)

type ContactControllerImpl struct {
	ContactService services.ContactService
}

func NewContactControllerImpl(contactService services.ContactService) *ContactControllerImpl {
	return &ContactControllerImpl{
		ContactService: contactService,
	}
}

func (controller *ContactControllerImpl) Create(ctx *fiber.Ctx) error {
	var contact requestbody.CreateContact
	err := ctx.BodyParser(&contact)
	if err != nil {
		return err
	}

	resp, err := controller.ContactService.Create(ctx.Context(), params.Contact{
		Contact: entities.Contact{
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Phone:     contact.Phone,
			IsVip:     contact.IsVip,
		},
	})
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
	param := params.NewContactParam(ctx)
	resp, err := controller.ContactService.Find(ctx.Context(), param)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   resp,
	}, "application/json")
}

func (controller *ContactControllerImpl) FindById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}
	param := params.Contact{Contact: entities.Contact{ID: uint(id)}}
	resp, err := controller.ContactService.FindById(ctx.Context(), param)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "OK",
		"code":   fiber.StatusOK,
		"data":   resp,
	}, "application/json")
}

func (controller *ContactControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}

	var contact requestbody.EditContact
	err = ctx.BodyParser(&contact)
	if err != nil {
		return err
	}

	resp, err := controller.ContactService.Update(
		ctx.Context(),
		params.Contact{Contact: entities.Contact{
			ID:        uint(id),
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Phone:     contact.Phone,
		}},
	)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   resp,
	}, "application/json")
}

func (controller *ContactControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}

	resp, err := controller.ContactService.Delete(
		ctx.Context(),
		params.Contact{Contact: entities.Contact{ID: uint(id)}},
	)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   resp,
	}, "application/json")
}
