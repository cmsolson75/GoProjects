package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/service"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/workflows"
)

type CustomerHandler struct {
	CustomerService service.CustomerService
	UserWorkflow    workflows.UserWorkflow
}

func NewCustomerHandler(customerService service.CustomerService, userWorkflow workflows.UserWorkflow) *CustomerHandler {
	return &CustomerHandler{
		CustomerService: customerService,
		UserWorkflow:    userWorkflow,
	}
}

// takes in name & email
// returns in JSON full customer
func (c *CustomerHandler) HandlePostNewCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer CustomerScheme
	json.NewDecoder(r.Body).Decode(&customer)
	c.UserWorkflow.CreateUser(customer.Name, customer.Email)
	json.NewEncoder(w).Encode(customer)
}

func (c *CustomerHandler) HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	customer, err := c.CustomerService.ViewCustomerByEmail(user.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Customer not found", http.StatusBadRequest)
		return
	}
	token, err := createToken(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error creating token.")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 5),
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, token)
}
