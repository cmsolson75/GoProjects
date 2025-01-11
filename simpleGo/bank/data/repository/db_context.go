package repository

import (
	"sync"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/data/utils"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/model"
)

// Customer: int is main id
// Account: have int be customer id: making it simple to search
type DBContext struct {
	sync.Mutex
	Customer              []*model.Customer
	Account               []*model.Account
	AutoIncrementCustomer utils.AutoIncrementInt
	AutoIncrementAccount  utils.AutoIncrementInt
}

func NewDBContext() *DBContext {
	return &DBContext{
		Customer:              []*model.Customer{},
		Account:               []*model.Account{},
		AutoIncrementCustomer: utils.AutoIncrementInt{},
		AutoIncrementAccount:  utils.AutoIncrementInt{},
	}
}
