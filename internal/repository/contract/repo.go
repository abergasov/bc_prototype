package contract

import (
	"bc_prototype/internal/storage/database"
	"database/sql"

	"github.com/google/uuid"
)

type Repo struct {
	db database.DBConnector
}

func InitRepo(db database.DBConnector) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetCircleContract(circleID uuid.UUID) (string, error) {
	var contractAddress string
	err := r.db.Client().QueryRowx("SELECT contract_address FROM circle_contract WHERE circle_id = $1", circleID).Scan(&contractAddress)
	if err != nil && err == sql.ErrNoRows {
		return "", nil
	}
	return contractAddress, err
}

func (r *Repo) EraseAll() error {
	_, err := r.db.Client().Exec("DELETE FROM circle_contract")
	return err
}
