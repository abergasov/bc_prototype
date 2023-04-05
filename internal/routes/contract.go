package routes

import (
	"bc_prototype/internal/utils"

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
	return c.JSON(contract)
}

func (s *Server) deployContract(c *fiber.Ctx) error {
	circleID, err := uuid.Parse(c.Params("circleID"))
	if err != nil {
		return err
	}
	contract, contractABI, err := s.service.DeployContract(circleID)
	if err != nil {
		return err
	}
	return c.JSON(map[string]interface{}{
		"contract_code": contract,
		"contract_abi":  contractABI,
		"params": []string{
			utils.UUIDToInt(circleID).String(),
		},
	})
}

func (s *Server) saveContract(c *fiber.Ctx) error {
	type update struct {
		CircleID        uuid.UUID `json:"circle_id"`
		TransactionHash string    `json:"transaction_hash"`
		ContractAddress string    `json:"contract_address"`
	}
	payload := new(update)
	if err := c.BodyParser(payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := s.service.SaveCircleContract(payload.CircleID, payload.TransactionHash, payload.ContractAddress); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(map[string]bool{"ok": true})
}
