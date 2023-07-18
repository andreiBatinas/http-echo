package domain

// Account model structure
type Account struct {
	Balance string `json:"balance"`
}

// AccountAddress struct
type AccountAddress struct {
	Address string `param:"address"`
}

// AccountUseCase usecase
type AccountUseCase interface {
	Balance(AccountAddress) (*Account, *AccountErrorResponse)
}

// AccountErrorResponse
type AccountErrorResponse struct {
	Error string `json:"error"`
}
