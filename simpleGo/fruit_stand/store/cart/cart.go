package cart

import (
	"errors"

	"github.com/cmsolson75/GoProjects/simpleGo/shopping_cart/store/price"
)

// Doesnt need to handle Checkout
// Checkout is dependent on Cart
type CartRepository interface {
	AddItem(item string, number int) error
	DeleteItem(name string)
	UpdateItem(item string, number int)
	GetItems() map[string]int
}

type Cart struct {
	items map[string]int
}

func NewCart(itemSliceMap ...map[string]int) (*Cart, error) {
	// handle nil errors
	items := make(map[string]int)
	cart := Cart{items}
	// handle cart init if present
	for _, itemMap := range itemSliceMap {
		for item, number := range itemMap {
			err := cart.AddItem(item, number)
			if err != nil {
				return &Cart{}, err
			}
		}
	}
	return &cart, nil
}

var (
	ErrNegativeInput = errors.New("invalid input: negative number")
)

// function for adding items to cart
// input item name and number you want to add.
func (c *Cart) AddItem(item string, number int) error {
	if number < 0 {
		return ErrNegativeInput
	}
	c.items[item] += number
	return nil
}

func (c *Cart) DeleteItem(item string) {
	delete(c.items, item)
}

// update total number for a specified item
func (c *Cart) UpdateItem(item string, number int) {
	c.items[item] = number
}

// function returns internal elements
// of Cart, used for fetching data for UI
func (c *Cart) GetItems() map[string]int {
	return c.items
}

// Checkout method: gives you total value of cart given a price table
// Move to Service
func (c *Cart) Checkout(p price.PriceRepository) (float64, error) {
	var totalAmount float64
	for item, num := range c.items {
		a, err := p.GetAmount(item)
		if err != nil {
			return 0.0, err
		}
		totalAmount += (a * float64(num))

	}
	return totalAmount, nil
}
