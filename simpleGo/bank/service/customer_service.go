package service

import (
	"errors"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/data/repository"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/model"
)

type CustomerService struct {
	CustomerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		CustomerRepo: customerRepo,
	}
}

// Might need to make a data model for sending over the wire
func (c *CustomerService) AddCustomer(email string, name string) *model.Customer {
	return c.CustomerRepo.CreateCustomer(email, name)
}

func (c *CustomerService) RemoveCustomer(id int) {
	c.CustomerRepo.DeleteCustomerByID(id)
}

func (c *CustomerService) ViewCustomer(id int) (*model.Customer, error) {
	return c.CustomerRepo.ReadCustomerByID(id)
}

func (c *CustomerService) ViewCustomerByEmail(email string) (*model.Customer, error) {
	for _, customer := range c.CustomerRepo.ViewAllCustomers() {
		if customer.Email == email {
			return customer, nil
		}
	}
	return &model.Customer{}, errors.New("customer not found")
}
