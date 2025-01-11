package workflows

import "github.com/cmsolson75/GoProjects/simpleGo/bank/service"

type UserWorkflow struct {
	CustomerService service.CustomerService
	AccountService  service.AccountService
}

func NewUserWorkflow(customerService service.CustomerService, accountService service.AccountService) *UserWorkflow {
	return &UserWorkflow{
		CustomerService: customerService,
		AccountService:  accountService,
	}
}

func (u *UserWorkflow) CreateUser(name string, email string) {
	customer := u.CustomerService.AddCustomer(email, name)
	u.AccountService.CreateNewAccount(customer)
}
