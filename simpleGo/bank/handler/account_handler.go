package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/model"
	"github.com/cmsolson75/GoProjects/simpleGo/bank/service"
)

type AccountHandler struct {
	AccountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: accountService,
	}
}

func (a *AccountHandler) HandleGetAccount(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing or invalid token")
		log.Println(err)
		return
	}
	tokenString := cookie.Value
	claims, err := validateToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid token")
		return
	}
	userID := int(claims["id"].(float64))
	log.Println(userID)
	balance, err := a.AccountService.ViewBalanceByID(userID)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprint(w, balance)
}

func (a *AccountHandler) HandlePostDeposit(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing or invalid token")
		log.Println(err)
		return
	}
	tokenString := cookie.Value
	claims, err := validateToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var total struct {
		Amount int `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&total); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	userID := int(claims["id"].(float64))
	customer := model.Customer{
		Name:  "",
		ID:    userID,
		Email: "",
	}
	a.AccountService.Account.AddBalance(&customer, total.Amount)
	w.WriteHeader(http.StatusOK)
}

func (a *AccountHandler) HandlePostWithdraw(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing or invalid token")
		log.Println(err)
		return
	}
	tokenString := cookie.Value
	claims, err := validateToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid token")
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var total struct {
		Amount int `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&total); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		log.Println(err)
		return
	}

	userID := int(claims["id"].(float64))
	customer := model.Customer{
		Name:  "",
		ID:    userID,
		Email: "",
	}
	err = a.AccountService.Account.RemoveBalance(&customer, total.Amount)
	if err != nil {
		http.Error(w, "Invalid info", http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
