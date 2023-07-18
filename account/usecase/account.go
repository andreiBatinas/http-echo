package usecase

import (
	"context"
	"math/big"

	"alluvial/interview/domain"
	ethereum "alluvial/interview/infrastructure/geth"
	"alluvial/interview/infrastructure/log"

	"github.com/ethereum/go-ethereum/common"
)

// AccountUseCaseDi
type AccountUseCaseDi struct {
	EthereumClient *ethereum.EthereumClient
	Log            *log.Logger
}

type accountUseCase struct {
	ethereumClient *ethereum.EthereumClient
	log            *log.Logger
}

func (ac accountUseCase) Balance(in domain.AccountAddress) (*domain.Account, *domain.AccountErrorResponse) {

	isValidAddress := common.IsHexAddress(in.Address)
	if !isValidAddress {
		ac.log.Error("Error address is not valid")
		return nil, &domain.AccountErrorResponse{
			Error: "address is not valid",
		}
	}

	address := common.HexToAddress(in.Address)

	balance, err := ac.ethereumClient.Client.BalanceAt(context.Background(), address, nil)

	if err != nil {
		ac.log.Error("Internal Error balanceAt")
		return nil, &domain.AccountErrorResponse{
			Error: "Internal error balance At",
		}
	}

	etherBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))

	response := &domain.Account{
		Balance: etherBalance.String(),
	}
	return response, nil
}

// NewAccountUseCase returns account use case object
func NewAccountUseCase(di AccountUseCaseDi) domain.AccountUseCase {
	return accountUseCase{
		ethereumClient: di.EthereumClient,
		log:            di.Log,
	}
}
