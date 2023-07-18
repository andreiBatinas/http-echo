package main

import (
	"log"

	"alluvial/interview/config"
	ethereum "alluvial/interview/infrastructure/geth"
	"alluvial/interview/infrastructure/http"
	_logger "alluvial/interview/infrastructure/log"

	"github.com/sirupsen/logrus"

	_serviceHttp "alluvial/interview/account/http"
	"alluvial/interview/account/usecase"
)

func main() {
	config.Read()

	mainLog := _logger.NewLogger(_logger.LoggerOptions{
		Level:  logrus.DebugLevel,
		Module: "main",
	})

	httpSrv := http.NewHttpServer(
		http.ServerOptions{
			Port: config.Configuration.HTTP.Port,
			Log: _logger.NewLogger(_logger.LoggerOptions{
				Level:  logrus.DebugLevel,
				Module: "http",
			}),
		},
	)

	ethClient, err := ethereum.NewEthereumClient(ethereum.EthereumClientOptions{
		RpcUrl: config.Configuration.RPC.Url,
	})

	if err != nil {
		mainLog.Error("Error connecting to ethereum rpc")
		log.Fatal(err)
	}

	accountUseCase := usecase.NewAccountUseCase(usecase.AccountUseCaseDi{
		EthereumClient: ethClient,
		Log: _logger.NewLogger(_logger.LoggerOptions{
			Level:  logrus.DebugLevel,
			Module: "accountUseCase",
		}),
	})

	_serviceHttp.NewAccountService(_serviceHttp.AccountServiceDi{
		HTTP:        httpSrv,
		UserUseCase: accountUseCase,
		Log: _logger.NewLogger(_logger.LoggerOptions{
			Level:  logrus.DebugLevel,
			Module: "accountService",
		}),
	})

	httpSrv.Start()

	mainLog.Info("Server started at port %d", config.Configuration.HTTP.Port)

}
