package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/fiber_nvim_crud/controllers"
	"github.com/nozzlium/fiber_nvim_crud/repositories"
	"github.com/nozzlium/fiber_nvim_crud/services"
)

var contactRepository repositories.ContactRepository = repositories.NewContactRepositoryMock()
var db *sql.DB = &sql.DB{}
var contactService services.ContactService = services.NewContactService(
	contactRepository,
	db,
)
var contactController controllers.ContactController = controllers.NewContactControllerImpl(
	contactService,
)
var app *fiber.App = fiber.New()

func GetApp() *fiber.App {
	app.Post("/contact", contactController.Create)
	app.Get("/contact", contactController.Find)
	app.Get("/contact/:id", contactController.FindById)
	app.Put("/contact/:id", contactController.Update)
	app.Delete("contact/:id", contactController.Delete)
	return app
}
