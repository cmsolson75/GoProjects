package service

import (
	"bytes"
	"fmt"
	"sort"
	"text/tabwriter"

	"github.com/cmsolson75/GoProjects/simpleGo/shopping_cart/store/price"
)

type PriceService struct {
	Price price.PriceRepository
}

func NewPriceService(priceRepo price.PriceRepository) *PriceService {
	return &PriceService{Price: priceRepo}
}

func (p *PriceService) GetAmount(name string) (float64, error) {
	amount, err := p.Price.GetAmount(name)
	return amount, err

}

func (p *PriceService) GetPriceByName(name string) (*price.Price, error) {
	price, err := p.Price.GetItem(name)
	return price, err
}

func (p *PriceService) GetAllPrices() ([]*price.Price, error) {
	price, err := p.Price.GetAll()
	return price, err
}

func (p *PriceService) GetNames() ([]string, error) {
	prices, err := p.GetAllPrices()
	if err != nil {
		return []string{}, err
	}
	var names []string
	for _, price := range prices {
		names = append(names, price.Name)
	}
	sort.Strings(names)
	return names, nil
}

func (p *PriceService) GetPriceUIFormat() (string, error) {
	outputBuffer := bytes.Buffer{}
	prices, err := p.GetAllPrices()
	if err != nil {
		return outputBuffer.String(), err
	}

	w := tabwriter.NewWriter(&outputBuffer, 15, 0, 3, ' ', 0)

	for _, price := range prices {
		fmt.Fprintf(w, "%s\t$%.2f\t\n", price.Name, price.Amount)
		w.Flush()
	}
	fmt.Fprint(&outputBuffer, "\n")
	return outputBuffer.String(), nil

}
