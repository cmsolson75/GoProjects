package model

type Account struct {
	ID         int `json:"id"`
	CustomerID int `json:"customerID"` // Foreign Key from Customer Model
	Balance    int `json:"balance"`    // this will get changed but for now is simpler to test whole numbers
}
