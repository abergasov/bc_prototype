package contract

import (
	"bc_prototype/internal/model"
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

func (r *Repo) GetCircleContract(circleID uuid.UUID) (*model.Contract, error) {
	var contract model.Contract
	err := r.db.Client().QueryRowx("SELECT * FROM circle_contract WHERE circle_id = $1", circleID).StructScan(&contract)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	return &contract, err
}

func (r *Repo) EraseAll() error {
	_, err := r.db.Client().Exec("DELETE FROM circle_contract")
	return err
}

func (r *Repo) SaveCircleContract(circleID uuid.UUID, transactionHash, contractAddress string) error {
	_, err := r.db.Client().Exec("INSERT INTO circle_contract (circle_id, contract_address, deploy_hash, created_at) VALUES ($1, $2, $3, NOW())", circleID, contractAddress, transactionHash)
	return err
}
