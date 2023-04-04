package circle

import "github.com/google/uuid"

func (s *Service) GetCircleContract(circleID uuid.UUID) (string, error) {
	return s.repoContract.GetCircleContract(circleID)
}

func (s *Service) DeployContract(circleID uuid.UUID) (string, error) {
	panic("implement me")
}
