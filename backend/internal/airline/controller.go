package airline

import "github.com/gofiber/fiber/v2"

// "github.com/gofiber/fiber/v2"
// "github.com/uowcs/CSIT214/internal/models"

type AirlineController struct {
	storage *AirlineStorage
}

func NewAirlineController(storage *AirlineStorage) *AirlineController {
	return &AirlineController{
		storage: storage,
	}
}

func (a *AirlineController) CreateSchemas(c *fiber.Ctx) error {
	err := a.storage.CreateSchemas()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Schemas created")
}
