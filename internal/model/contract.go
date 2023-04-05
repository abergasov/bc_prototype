package model

import (
	"time"

	"github.com/google/uuid"
)

type Contract struct {
	CircleID        uuid.UUID `json:"circle_id" db:"circle_id"`
	ContractAddress string    `json:"contract_address" db:"contract_address"`
	DeployTxHash    string    `json:"deploy_hash" db:"deploy_hash"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
