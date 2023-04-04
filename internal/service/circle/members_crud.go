package circle

import (
	"bc_prototype/internal/model"

	"github.com/google/uuid"
)

func (s *Service) GetMembers(circleID uuid.UUID) ([]*model.Member, error) {
	return s.repoMembers.GetCircleMembers(circleID)
}

func (s *Service) GetMember(circleID, memberID uuid.UUID) (*model.Member, error) {
	return s.repoMembers.GetMemberByID(circleID, memberID)
}

func (s *Service) SetMemberAddress(memberID uuid.UUID, address string) error {
	return s.repoMembers.SetMemberAddress(memberID, address)
}
