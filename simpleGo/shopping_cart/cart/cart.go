package cart

import "errors"

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

// function returns internal elements
// of Cart, used for fetching data for UI
func (c *Cart) GetItems() map[string]int {
	return c.items
}

// checkout
//
