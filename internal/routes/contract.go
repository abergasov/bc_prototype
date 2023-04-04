package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Server) getCircleContract(c *fiber.Ctx) error {
	circleID, err := uuid.Parse(c.Params("circleID"))
	if err != nil {
		return err
	}
	contract, err := s.service.GetCircleContract(circleID)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{
		"address": contract,
	})
}

func (s *Server) deployContract(c *fiber.Ctx) error {
	circleID, err := uuid.Parse(c.Params("circleID"))
	if err != nil {
		return err
	}
	contract, err := s.service.DeployContract(circleID)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{
		"address": contract,
	})
}
