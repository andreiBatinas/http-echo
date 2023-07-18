package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/gommon/log"
)

type EthereumClientOptions struct {
	RpcUrl string
}

// EthereumClient
type EthereumClient struct {
	Client *ethclient.Client
}

func NewEthereumClient(opts EthereumClientOptions) (*EthereumClient, error) {
	client, err := ethclient.Dial(opts.RpcUrl)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Retrieve the latest block number to verify the connection
	_, err = client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	ethClient := &EthereumClient{
		Client: client,
	}

	return ethClient, nil
}
