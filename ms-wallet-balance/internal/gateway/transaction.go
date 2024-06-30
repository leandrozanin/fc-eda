package gateway

import "github.com/leandrozanin/fc-eda/ms-wallet-balance/internal/entity"

type BalanceGateway interface {
	Save(balance *entity.Balance) error
	Update(balance *entity.Balance) error
	FindByAccountID(accountID string) (*entity.Balance, error)
}
