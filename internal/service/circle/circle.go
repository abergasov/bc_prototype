package circle

import (
	"bc_prototype/internal/contracts"
	"bc_prototype/internal/model"
	_ "embed"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) GetCircleContract(circleID uuid.UUID) (*model.Contract, error) {
	return s.repoContract.GetCircleContract(circleID)
}

func (s *Service) SaveCircleContract(circleID uuid.UUID, transactionHash, contractAddress string) error {
	// todo validate
	if err := s.repoContract.SaveCircleContract(circleID, transactionHash, contractAddress); err != nil {
		return fmt.Errorf("failed to save circle contract: %w", err)
	}
	return nil
}

func (s *Service) DeployContract(circleID uuid.UUID) (string, string, error) {
	hexAddress, err := s.repoMembers.GetDirectorsAddress(circleID)
	if err != nil {
		return "", "", fmt.Errorf("failed to get directors address: %w", err)
	}
	if hexAddress == "" {
		return "", "", fmt.Errorf("directors address is empty")
	}

	//	contract, err := contracts.NewContracts(address, s.ethClient)
	//	contract.
	//	contract, _, _, err := contracts.DeployContracts(address, s.ethClient)
	//if err != nil {
	//	return "", fmt.Errorf("failed to create contract: %w", err)
	//}
	//contract.

	//// Set up the contract deployer and the transaction options
	//auth := bind.NewClefTransactor(nil, s.chainID)
	//
	////auth.GasLimit = uint64(3000000)
	//auth.Value = big.NewInt(0)

	//contractABI, err := contracts.ContractsMetaData.GetAbi()
	//if err != nil {
	//	return "", fmt.Errorf("")
	//}
	//// Generate the raw transaction data for deploying the contract
	//txData, err := contractABI.Pack("", circleID)
	//if err != nil {
	//	return "", err
	//}
	return contracts.ContractsMetaData.Bin, contracts.ContractsMetaData.ABI, nil
}

//func (s *Service) GenerateRawDeployTxData(client *Client, contractBin []byte, owner common.Address) ([]byte, error) {
//	// Set up the contract deployer and the transaction options
//	auth, err := bind.NewKeyedTransactorWithChainID(nil, s.chainID)
//	if err != nil {
//		return nil, err
//	}
//	//auth.GasLimit = uint64(3000000)
//	auth.Value = big.NewInt(0)
//
//	contractABI, err := contracts.ContractsMetaData.GetAbi()
//	if err != nil {
//		return nil, fmt.Errorf("")
//	}
//	// Generate the raw transaction data for deploying the contract
//	txData, err := contractABI.Pack("", contractBin, circleID)
//	if err != nil {
//		return nil, err
//	}
//
//	// Create the transaction object
//	tx := types.NewContractCreation(
//		uint64(0),
//		auth.Value,
//		auth.GasLimit,
//		big.NewInt(0),
//		txData,
//	)
//
//	// Set the receiver to the null address, which indicates that this is a contract deployment
//	tx.To = nil
//
//	// Set the nonce and gas price
//	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
//	if err != nil {
//		return nil, err
//	}
//	tx.Nonce = big.NewInt(int64(nonce))
//
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		return nil, err
//	}
//	tx.GasPrice = gasPrice
//
//	// Calculate the gas limit
//	estimatedGasLimit, err := client.EstimateGas(context.Background(), tx)
//	if err != nil {
//		return nil, err
//	}
//	tx.GasLimit = estimatedGasLimit
//
//	// Sign the transaction and return the raw transaction data
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(client.chainID), auth.PrivateKey)
//	if err != nil {
//		return nil, err
//	}
//
//	return signedTx.MarshalBinary()
//}
