package repository

import (
	"errors"
	"log"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/model"
)

type AccountRepository interface {
	CreateAccount(customer *model.Customer)
	ReadAccountByCustomerID(customerID int) (*model.Account, error)
	AddBalance(customer *model.Customer, amount int) error
	RemoveBalance(customer *model.Customer, amount int) error
	ViewAccountBalance(customer *model.Customer) (int, error)
}

type InMemoryAccount struct {
	DB *DBContext
}

func NewInMemoryAccount(db *DBContext) *InMemoryAccount {
	return &InMemoryAccount{DB: db}
}

func (i *InMemoryAccount) CreateAccount(customer *model.Customer) {
	i.DB.Lock()
	defer i.DB.Unlock()
	account := model.Account{
		ID:         i.DB.AutoIncrementAccount.ID(),
		CustomerID: customer.ID,
		Balance:    0,
	}
	i.DB.Account = append(i.DB.Account, &account)
}

func (i *InMemoryAccount) ReadAccountByCustomerID(customerID int) (*model.Account, error) {
	for _, account := range i.DB.Account {
		if account.CustomerID == customerID {
			return account, nil
		}
	}
	return &model.Account{}, errors.New("customer eather doesnt exist or doesnt have an account with us")
}

func (i *InMemoryAccount) ViewAccountBalance(customer *model.Customer) (int, error) {
	i.DB.Lock()
	defer i.DB.Unlock()
	account, err := i.ReadAccountByCustomerID(customer.ID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return account.Balance, nil
}

func (i *InMemoryAccount) AddBalance(customer *model.Customer, amount int) error {
	if amount < 0 {
		return errors.New("cannot add a negative value to account balance")
	}
	account, err := i.ReadAccountByCustomerID(customer.ID)
	if err != nil {
		return err
	}
	account.Balance += amount
	return nil
}

func (i *InMemoryAccount) RemoveBalance(customer *model.Customer, amount int) error {
	i.DB.Lock()
	defer i.DB.Unlock()
	account, err := i.ReadAccountByCustomerID(customer.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	if amount > account.Balance {
		// should make message more helpful
		err = errors.New("cannot overdraft from account")
		log.Println(err)
		return err
	}
	account.Balance -= amount
	return nil
}
