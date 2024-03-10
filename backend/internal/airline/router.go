package airline

import "github.com/gofiber/fiber/v2"

func AddAirlineRoutes(app *fiber.App, controller *AirlineController) {
	Airline := app.Group("/")
	Airline.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})
	Airline.Get("/create-schemas", controller.CreateSchemas)
}
