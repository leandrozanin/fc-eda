package gateway

import "github.com/leandrozanin/fc-eda/fc-ms-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
