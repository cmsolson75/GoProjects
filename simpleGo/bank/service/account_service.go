package service

import (
	"github.com/cmsolson75/GoProjects/simpleGo/bank/data/repository"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/model"
)

type AccountService struct {
	Account repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) *AccountService {
	return &AccountService{
		Account: accountRepo,
	}
}

func (a *AccountService) CreateNewAccount(customer *model.Customer) {
	a.Account.CreateAccount(customer)
}

func (a *AccountService) Deposit(customer *model.Customer, ammount int) error {
	return a.Account.AddBalance(customer, ammount)
}

func (a *AccountService) Withdraw(customer *model.Customer, ammount int) error {
	return a.Account.RemoveBalance(customer, ammount)
}

func (a *AccountService) ViewBalance(customer *model.Customer) (int, error) {
	return a.Account.ViewAccountBalance(customer)
}

func (a *AccountService) ViewBalanceByID(id int) (int, error) {
	account, err := a.Account.ReadAccountByCustomerID(id)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil

}
