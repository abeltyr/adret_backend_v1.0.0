package companyRouter

import (
	"adr/backend/src/service/companyService"
	"adr/backend/src/utils"

	"github.com/gofiber/fiber/v2"
)

func Find(c *fiber.Ctx) error {

	id := c.Params("id")

	client, ctx, err := utils.PrismaClient()
	if err != nil {
		return err
	}

	data, err := companyService.Find(id, client, ctx)
	if err != nil {
		return err
	}
	return c.JSON(data)
}
