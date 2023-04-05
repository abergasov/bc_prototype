package circle

import (
	"bc_prototype/internal/logger"
	"bc_prototype/internal/repository/contract"
	"bc_prototype/internal/repository/members"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	"go.uber.org/zap"
)

type Service struct {
	log          logger.AppLogger
	repoMembers  *members.Repo
	repoContract *contract.Repo
	ethClient    *ethclient.Client
	chainID      *big.Int
}

func InitService(log logger.AppLogger, rpcURL string, repoMembers *members.Repo, repoContract *contract.Repo) (*Service, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to ethereum node: %w", err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to get chain id: %w", err)
	}
	return &Service{
		repoMembers:  repoMembers,
		repoContract: repoContract,
		log:          log.With(zap.String("service", "sampler")),
		ethClient:    client,
		chainID:      chainID,
	}, nil
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
