package service

import (
	"reflect"
	"testing"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/data/repository"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/model"
)

func TestCustomerService(t *testing.T) {
	// All in one test to keep setup time down
	ctx := repository.NewDBContext()
	// Mock DB
	customerRepo := repository.NewInMemoryCustomer(ctx)

	service := NewCustomerService(customerRepo)

	// test add customer
	service.AddCustomer("test@email.com", "test")
	got := &ctx.Customer
	want := []*model.Customer{{ID: 0, Email: "test@email.com", Name: "test"}}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("test AddCustomer:\n\tgot %v want %v", got, want)
	}

	// test view customer
	customer, err := service.ViewCustomer(0)
	if err != nil {
		t.Fatal("test ViewCustomer:\n\tgot error where none expected")
	}

	wantCustomer := model.Customer{ID: 0, Email: "test@email.com", Name: "test"}

	if !reflect.DeepEqual(*customer, wantCustomer) {
		t.Errorf("test ViewCustomer:\n\tgot %v want %v", customer, wantCustomer)
	}

	service.RemoveCustomer(0)
	emptyCustomerDB := []*model.Customer{}
	if !reflect.DeepEqual(*got, emptyCustomerDB) {
		t.Errorf("test RemoveCustomer:\n\tgot %v want %v", got, emptyCustomerDB)
	}
}
