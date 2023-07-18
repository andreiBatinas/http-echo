package http

import (
	"net/http"

	"alluvial/interview/domain"
	_http "alluvial/interview/infrastructure/http"
	"alluvial/interview/infrastructure/log"

	"github.com/labstack/echo/v4"
)

type accountService struct {
	http           *_http.Server
	accountUseCase domain.AccountUseCase
	log            *log.Logger
}

// AccountServiceDi dependencies
type AccountServiceDi struct {
	HTTP        *_http.Server
	UserUseCase domain.AccountUseCase
	Log         *log.Logger
}

func (srv accountService) Handler() {
	api := srv.http.API(_http.NewRouter("/balance"))

	api.GET("/:address", func(c echo.Context) error {
		in := &domain.AccountAddress{}
		c.Bind(in)

		res, err := srv.accountUseCase.Balance(*in)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, res)
	})
}

// NewAccountService
func NewAccountService(di AccountServiceDi) {
	srv := &accountService{
		http:           di.HTTP,
		accountUseCase: di.UserUseCase,
		log:            di.Log,
	}

	di.HTTP.SubscribeRouter(srv)
}
