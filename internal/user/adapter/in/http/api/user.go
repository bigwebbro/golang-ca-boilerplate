package api

import (
	"github.com/gofiber/fiber/v2"
)

func (a Adapter) CreateUser(c *fiber.Ctx) error {
	var r = &CreateUserRequest{}
	err := c.BodyParser(&r)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate CreateUserRequest

	cmd, errs := r.toCommand()
	if len(errs) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, errs...)
	}

	u, err := a.ucUser.Create(cmd)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(u)
}
