package price

import (
	"errors"
	"fmt"
)

var (
	ErrPriceExists   = errors.New("price exists")
	ErrPriceNotFound = errors.New("price not found")
)

type Price struct {
	Name   string
	Amount float64
}

func (p *Price) String() string {
	return fmt.Sprintf("Price{Name: %s, Price: %.2f}", p.Name, p.Amount)
}

type PriceRepository interface {
	GetItem(name string) (*Price, error)
	GetAmount(name string) (float64, error)
	GetAll() ([]*Price, error)
	Create(price *Price) error
}

type InMemoryPrice struct {
	Storage map[string]Price
}

func NewInMemoryPrice(prices ...Price) (*InMemoryPrice, error) {
	i := InMemoryPrice{}
	i.Storage = make(map[string]Price)
	for _, price := range prices {
		err := i.Create(&price)
		if err != nil {
			return &InMemoryPrice{}, err
		}
	}
	return &i, nil
}

func (i *InMemoryPrice) Create(price *Price) error {
	if _, ok := i.Storage[price.Name]; ok {
		return ErrPriceExists
	}

	i.Storage[price.Name] = *price
	return nil
}

func (i *InMemoryPrice) GetItem(name string) (*Price, error) {
	// Don't like this -> but its okay for now
	if price, ok := i.Storage[name]; ok {
		return &price, nil
	}
	return &Price{}, ErrPriceNotFound
}

func (i *InMemoryPrice) GetAmount(name string) (float64, error) {
	price, err := i.GetItem(name)
	if err != nil {
		return 0.0, err
	}

	return price.Amount, nil
}

func (i *InMemoryPrice) GetAll() ([]*Price, error) {
	var prices []*Price

	for _, price := range i.Storage {
		prices = append(prices, &price)
	}

	return prices, nil
}
