package service

import (
	"bytes"
	"fmt"
	"sort"
	"text/tabwriter"

	"github.com/cmsolson75/GoProjects/simpleGo/shopping_cart/store/cart"
)

type CartService struct {
	Cart cart.CartRepository
}

func NewCartService(cartRepo cart.CartRepository) *CartService {
	return &CartService{Cart: cartRepo}
}

func (c *CartService) Checkout(p PriceService) (float64, error) {
	var total float64
	for item, number := range c.Cart.GetItems() {
		a, err := p.GetAmount(item)
		if err != nil {
			return 0.0, err
		}
		total += (a * float64(number))
	}
	return total, nil
}

func (c *CartService) GetNames() []string {
	var items []string

	for name := range c.Cart.GetItems() {
		items = append(items, name)

	}
	sort.Strings(items)
	return items
}

func (c *CartService) GetAllItems() map[string]int {
	return c.Cart.GetItems()
}

func (c *CartService) AddItem(item string, number int) error {
	err := c.Cart.AddItem(item, number)
	if err != nil {
		return err
	}
	return nil
}

func (c *CartService) DeleteItemByName(name string) {
	c.Cart.DeleteItem(name)
}

func (c *CartService) UpdateItem(item string, number int) {
	c.Cart.UpdateItem(item, number)
}

func (c *CartService) GetCartUIFormat() string {
	outputBuffer := bytes.Buffer{}
	w := tabwriter.NewWriter(&outputBuffer, 15, 0, 3, ' ', 0)
	fmt.Fprint(&outputBuffer, "\n")
	fmt.Fprintf(w, "%s\t%s\t\n", "Name", "Count")
	w.Flush()

	for name, count := range c.GetAllItems() {
		fmt.Fprintf(w, "%s\t%d\t\n", name, count)
		w.Flush()
	}

	return outputBuffer.String()
}
