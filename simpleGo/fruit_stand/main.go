package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cmsolson75/GoProjects/simpleGo/shopping_cart/application"
	"github.com/cmsolson75/GoProjects/simpleGo/shopping_cart/service"
	"github.com/cmsolson75/GoProjects/simpleGo/shopping_cart/store/cart"
	"github.com/cmsolson75/GoProjects/simpleGo/shopping_cart/store/price"
)

func main() {
	userCart, _ := cart.NewCart()
	cartService := service.NewCartService(userCart)
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		log.Println("Error:", err)
	}
	defer db.Close()
	priceRepo := price.SQLitePrice{Storage: db}
	priceService := service.NewPriceService(&priceRepo)
	app := application.NewApp(cartService, priceService)
	app.Run()
}
