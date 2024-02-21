package params

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/fiber_nvim_crud/entities"
)

type Contact struct {
	PageNo   uint
	PageSize uint
	Contact  entities.Contact
}

func NewContactParam(ctx *fiber.Ctx) Contact {
	return Contact{
		PageNo:   uint(ctx.QueryInt("pageNo", 1)),
		PageSize: uint(ctx.QueryInt("pageSize", 10)),
	}
}
