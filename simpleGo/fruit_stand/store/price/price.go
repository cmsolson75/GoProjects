package price

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
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

type SQLitePrice struct {
	Storage *sql.DB
}

func (s *SQLitePrice) GetItem(name string) (*Price, error) {
	price := Price{}
	err := s.Storage.QueryRow("SELECT name, amount FROM produce WHERE name = ?", name).Scan(&price.Name, &price.Amount)
	if err == sql.ErrNoRows {
		return nil, ErrPriceNotFound
	} else if err != nil {
		return nil, err
	}
	return &price, nil
}

func (s *SQLitePrice) GetAmount(name string) (float64, error) {
	price, err := s.GetItem(name)
	if err != nil {
		log.Println("Error:", err)
		return 0.0, err
	}
	return price.Amount, nil
}

func (s *SQLitePrice) GetAll() ([]*Price, error) {
	rows, err := s.Storage.Query("SELECT * FROM produce")
	if err != nil {
		log.Println("Error Getting All Prices:", err)
		return []*Price{}, err
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Println("Rows Error:", err)
		return []*Price{}, err
	}

	prices := make([]*Price, 0)
	for rows.Next() {
		price := Price{}
		err = rows.Scan(&price.Name, &price.Amount)
		if err != nil {
			log.Println("Error:", err)
			return []*Price{}, err
		}
		prices = append(prices, &price)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error:", err)
		return []*Price{}, err
	}
	return prices, nil
}

var priceData = []Price{
	{Name: "banana's", Amount: 1.19},
	{Name: "strawberry's", Amount: 3.99},
	{Name: "blueberry's", Amount: 5.99},
	{Name: "raspberry's", Amount: 3.99},
	{Name: "blackberry's", Amount: 5.29},
	{Name: "apple", Amount: 1.59},
	{Name: "grape's", Amount: 3.99},
	{Name: "orange", Amount: 1.29},
}

var PublicPriceTable, _ = NewInMemoryPrice(priceData...)
