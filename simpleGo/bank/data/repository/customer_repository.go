package repository

import (
	"errors"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/model"
)

type CustomerRepository interface {
	CreateCustomer(email string, name string) *model.Customer
	ReadCustomerByID(id int) (*model.Customer, error)
	DeleteCustomerByID(id int)
	ViewAllCustomers() []*model.Customer
}

type InMemoryCustomer struct {
	DB *DBContext
}

func NewInMemoryCustomer(db *DBContext) *InMemoryCustomer {
	return &InMemoryCustomer{DB: db}
}

func (i *InMemoryCustomer) CreateCustomer(email string, name string) *model.Customer {
	i.DB.Lock()
	defer i.DB.Unlock()
	customer := model.Customer{
		ID:    i.DB.AutoIncrementCustomer.ID(),
		Email: email,
		Name:  name,
	}
	i.DB.Customer = append(i.DB.Customer, &customer)
	return &customer
}

// This is a bit cringe
func (i *InMemoryCustomer) DeleteCustomerByID(id int) {
	i.DB.Lock()
	defer i.DB.Unlock()
	for idx, customer := range i.DB.Customer {
		if customer.ID == id {
			i.DB.Customer = append(i.DB.Customer[:idx], i.DB.Customer[idx+1:]...)
		}
	}
	for idx, account := range i.DB.Account {
		if account.CustomerID == id {
			i.DB.Account = append(i.DB.Account[:idx], i.DB.Account[idx+1:]...)
		}
	}
}

func (i *InMemoryCustomer) ReadCustomerByID(id int) (*model.Customer, error) {
	i.DB.Lock()
	defer i.DB.Unlock()
	for _, customer := range i.DB.Customer {
		if customer.ID == id {
			return customer, nil
		}
	}
	return &model.Customer{}, errors.New("customer not in database")
}

func (i *InMemoryCustomer) ViewAllCustomers() []*model.Customer {
	i.DB.Lock()
	defer i.DB.Unlock()
	return i.DB.Customer
}
