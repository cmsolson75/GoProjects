package application

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/cmsolson75/GoProjects/simpleGo/shopping_cart/service"
)

var appTitle = `
███████╗██████╗ ██╗   ██╗██╗████████╗    ███████╗████████╗ █████╗ ███╗   ██╗██████╗
██╔════╝██╔══██╗██║   ██║██║╚══██╔══╝    ██╔════╝╚══██╔══╝██╔══██╗████╗  ██║██╔══██╗
█████╗  ██████╔╝██║   ██║██║   ██║       ███████╗   ██║   ███████║██╔██╗ ██║██║  ██║
██╔══╝  ██╔══██╗██║   ██║██║   ██║       ╚════██║   ██║   ██╔══██║██║╚██╗██║██║  ██║
██║     ██║  ██║╚██████╔╝██║   ██║       ███████║   ██║   ██║  ██║██║ ╚████║██████╔╝
╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝   ╚═╝       ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝
`

var cartTitle = `
╔═════════════════════════════════╗
║ ██████╗ █████╗ ██████╗ ████████╗║
║██╔════╝██╔══██╗██╔══██╗╚══██╔══╝║
║██║     ███████║██████╔╝   ██║   ║
║██║     ██╔══██║██╔══██╗   ██║   ║
║╚██████╗██║  ██║██║  ██║   ██║   ║ 
║ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ║
╚═════════════════════════════════╝ 
`

var priceTitle = `
██████╗ ██████╗ ██╗ ██████╗███████╗███████╗
██╔══██╗██╔══██╗██║██╔════╝██╔════╝██╔════╝
██████╔╝██████╔╝██║██║     █████╗  ███████╗
██╔═══╝ ██╔══██╗██║██║     ██╔══╝  ╚════██║
██║     ██║  ██║██║╚██████╗███████╗███████║
╚═╝     ╚═╝  ╚═╝╚═╝ ╚═════╝╚══════╝╚══════╝
                                           
`

type App struct {
	cart   *service.CartService
	prices *service.PriceService
}

func NewApp(cart *service.CartService, prices *service.PriceService) *App {
	return &App{cart: cart, prices: prices}
}

func (a *App) Run() {
	a.ViewForm(appTitle, "Welcome to my fruit stand: ")

	for {
		option := a.MenuForm()
		switch option {
		case "add":
			name, number, err := a.AddForm()
			if err != nil {
				fmt.Println("Error: ", err)
			}
			err = a.cart.AddItem(name, number)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		case "update":
			name, number, err := a.UpdateForm()
			if err != nil {
				fmt.Println("Error:", err)
			}
			a.cart.UpdateItem(name, number)
		case "delete":
			name := a.DeleteForm()
			a.cart.DeleteItemByName(name)
		case "prices":
			prices, err := a.prices.GetPriceUIFormat()
			if err != nil {
				fmt.Println("Error:", err)
			}
			a.ViewForm(priceTitle, prices)
		case "view":
			cartView := a.cart.GetCartUIFormat()

			a.ViewForm(cartTitle, cartView)
		case "checkout":
			checkoutTotal, err := a.cart.Checkout(*a.prices)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("$%.2f\n", checkoutTotal)
			return
		case "quit":
			if a.ConfirmForm() {
				return
			}

		}
	}
}

func (a *App) DeleteForm() string {
	var item string
	cartItems := a.cart.GetNames()
	// for safe exit
	cartItems = append(cartItems, "none")

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pick Fruit").
				OptionsFunc(func() []huh.Option[string] {
					return huh.NewOptions(cartItems...)
				}, "").Value(&item),
		),
	)
	form.Run()
	return item

}

func (a *App) UpdateForm() (string, int, error) {
	var item, number string
	cartItems := a.cart.GetNames()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pick Fruit").
				OptionsFunc(func() []huh.Option[string] {
					return huh.NewOptions(cartItems...)
				}, "").Value(&item),
			huh.NewInput().
				Title("What number do you want to update too.").
				Value(&number).
				Validate(func(str string) error {
					if _, err := strconv.Atoi(str); err != nil {
						return errors.New("only input whole numbers")
					}
					return nil
				}),
		),
	)
	form.Run()
	num, err := strconv.Atoi(number)
	if err != nil {
		return "", 0, err
	}
	return item, num, err

}

func (a *App) AddForm() (string, int, error) {
	var name, stringNumber string
	fruits, err := a.prices.GetNames()
	if err != nil {
		return "", 0, err
	}
	addForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pick Fruit").
				OptionsFunc(func() []huh.Option[string] {
					return huh.NewOptions(fruits...)
				}, "").Value(&name),
			huh.NewInput().
				Title("How many to add?").
				Value(&stringNumber).
				Validate(func(str string) error {
					if _, err := strconv.Atoi(str); err != nil {
						return errors.New("only input whole numbers")
					}
					return nil
				}),
		),
	)
	addForm.Run()

	number, err := strconv.Atoi(stringNumber)
	if err != nil {
		return "", 0, err
	}
	return name, number, nil
}

func (a *App) MenuForm() (option string) {
	home := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What do you want to do?").
				Options(
					huh.NewOption("Add item to cart", "add"),
					huh.NewOption("Update Item from cart", "update"),
					huh.NewOption("Delete Item from cart", "delete"),
					huh.NewOption("View Cart", "view"),
					huh.NewOption("View Prices", "prices"),
					huh.NewOption("Checkout", "checkout"),
					huh.NewOption("Quit", "quit"),
				).Value(&option),
		),
	)
	home.Run()
	return
}

func (a *App) ConfirmForm() (choice bool) {
	huh.NewConfirm().
		Title("Are you sure?").
		Value(&choice).Run()
	return
}

func (a *App) ViewForm(title, description string) {
	huh.NewNote().
		Title(title).
		Description(description).
		Next(true).
		NextLabel("Continue").Run()
}
