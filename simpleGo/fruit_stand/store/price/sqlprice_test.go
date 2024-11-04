package price

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	db, err := sql.Open("sqlite3",
		"/Users/cameronolson/Developer/go_code/GoProjects/simpleGo/fruit_stand/store.db")
	if err != nil {
		t.Fatal("Error Opening:", err)
	}
	defer db.Close()

	p := SQLitePrice{Storage: db}
	price, err := p.GetItem("apple")
	if err != nil {
		t.Fatal("Error Getting Item:", err)
	}
	fmt.Println(price.Amount)

	prices, err := p.GetAll()
	if err != nil {
		t.Fatal("error getting all:", err)
	}
	fmt.Println(prices)
}
