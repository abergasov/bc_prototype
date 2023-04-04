package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Server) listOfMembers(ctx *fiber.Ctx) error {
	list, err := s.service.GetMembers(testCircle)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(list)
}

func (s *Server) saveMemberAddress(c *fiber.Ctx) error {
	type update struct {
		Member  uuid.UUID `json:"member"`
		Address string    `json:"address"`
	}
	payload := new(update)
	if err := c.BodyParser(payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := s.service.SetMemberAddress(payload.Member, payload.Address); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(map[string]bool{"ok": true})
}
