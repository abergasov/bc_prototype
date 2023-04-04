package members

import (
	"bc_prototype/internal/model"
	"bc_prototype/internal/storage/database"
	"fmt"

	"github.com/google/uuid"
)

type Repo struct {
	db database.DBConnector
}

func InitRepo(db database.DBConnector) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetMemberByID(circleID, memberID uuid.UUID) (*model.Member, error) {
	var res model.Member
	err := r.db.Client().QueryRowx("SELECT * FROM circle_members WHERE circle_id = $1 AND member_id = $2", circleID, memberID).StructScan(&res)
	return &res, err
}

func (r *Repo) GetCircleMembers(circleID uuid.UUID) ([]*model.Member, error) {
	res := make([]*model.Member, 0, 10)
	rows, err := r.db.Client().Queryx("SELECT * FROM circle_members WHERE circle_id = $1", circleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var member model.Member

		if err = rows.StructScan(&member); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		member.Address = member.AddressSQL.String
		res = append(res, &member)
	}
	return res, err
}

func (r *Repo) SetMemberAddress(memberID uuid.UUID, address string) error {
	_, err := r.db.Client().Exec("UPDATE circle_members SET address = $1 WHERE member_id = $2", address, memberID)
	return err
}
