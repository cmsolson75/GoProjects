package service

import (
	"testing"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/data/repository"
)

func TestAccountService(t *testing.T) {
	// setup
	ctx := repository.NewDBContext()
	customerRepo := repository.NewInMemoryCustomer(ctx)
	accountRepo := repository.NewInMemoryAccount(ctx)

	// service layer
	customerService := NewCustomerService(customerRepo)
	accountService := NewAccountService(accountRepo)

	// we are here to test accounts not customers but we need some

	customerService.AddCustomer("test1@email.com", "test")
	// Can we open a new account
	customer, err := customerService.ViewCustomer(0)
	if err != nil {
		t.Fatal("error encountered where none expected")
	}
	accountService.CreateNewAccount(customer)
	accountService.Deposit(customer, 3000)

	amount, err := accountService.ViewBalance(customer)
	if err != nil {
		t.Fatal("error encountered where none expected")
	}

	want := 3000
	if amount != want {
		t.Errorf("got %d want %d", amount, want)
	}

	accountService.Withdraw(customer, 1000)
	want = 2000
	amount, err = accountService.ViewBalance(customer)
	if err != nil {
		t.Fatal("error encountered where none expected")
	}

	if amount != want {
		t.Errorf("got %d want %d", amount, want)
	}

}
