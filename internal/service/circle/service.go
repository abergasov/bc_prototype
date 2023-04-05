package circle

import (
	"bc_prototype/internal/logger"
	"bc_prototype/internal/repository/contract"
	"bc_prototype/internal/repository/members"
	"fmt"

	"go.uber.org/zap"
)

type Service struct {
	log          logger.AppLogger
	repoMembers  *members.Repo
	repoContract *contract.Repo
}

func InitService(log logger.AppLogger, repoMembers *members.Repo, repoContract *contract.Repo) *Service {
	return &Service{
		repoMembers:  repoMembers,
		repoContract: repoContract,
		log:          log.With(zap.String("service", "sampler")),
	}
}

func (s *Service) EraseAll() error {
	if err := s.repoMembers.EraseAll(); err != nil {
		return fmt.Errorf("failed to erase all members: %w", err)
	}
	if err := s.repoContract.EraseAll(); err != nil {
		return fmt.Errorf("failed to erase circle contract: %w", err)
	}
	return nil
}
