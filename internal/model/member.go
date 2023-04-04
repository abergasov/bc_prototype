package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Member struct {
	CircleID   uuid.UUID      `json:"circle_id" db:"circle_id"`
	MemberID   uuid.UUID      `json:"member_id" db:"member_id"`
	AddressSQL sql.NullString `json:"-" db:"address"`
	Address    string         `json:"address" db:"-"`
	Role       uint8          `json:"role" db:"role"`
}
