package price

import (
	"reflect"
	"testing"
)

// Need to test for error case
//

func TestInMemoryPriceCreate(t *testing.T) {
	testTable := []struct {
		testName          string
		initInMemoryPrice []Price
		createPrice       Price
		want              Price
		expectedErr       error
	}{
		{
			testName:          "basic test",
			initInMemoryPrice: []Price{},
			createPrice:       Price{Name: "item", Amount: 10.20},
			want:              Price{Name: "item", Amount: 10.20},
			expectedErr:       nil,
		},
		{
			testName:          "item exists error case",
			initInMemoryPrice: []Price{{Name: "item", Amount: 10.20}},
			createPrice:       Price{Name: "item", Amount: 10.20},
			want:              Price{},
			expectedErr:       ErrPriceExists,
		},
	}

	for _, tt := range testTable {
		// I would usually call i -> inMemoryPriceInstance but I am
		// Trying to be ideomatic to go
		i, err := NewInMemoryPrice(tt.initInMemoryPrice...)
		if err != nil {
			t.Fatal("error encountered when none expected")
		}

		// Main test point
		err = i.Create(&tt.createPrice)

		// Fail correctly
		if tt.expectedErr != nil {
			if err != tt.expectedErr {
				t.Fatalf("test: %s\n unexpected error encountered \n got %s want %s",
					tt.testName,
					err.Error(),
					tt.expectedErr.Error())
			}
			// exit to stop test breaks: this is the correct behavior
			// to make the tests more flexable.
			break
		}

		got, err := i.GetItem(tt.createPrice.Name)
		if err != nil {
			t.Fatal("error encountered when none expected")
		}

		if !reflect.DeepEqual(*got, tt.want) {
			t.Errorf("got %q want %q", got.String(), tt.want.String())
		}

	}
}

func TestInMemoryPriceGetItem(t *testing.T) {
	testTable := []struct {
		testName          string
		initInMemoryPrice []Price
		itemName          string
		priceWant         Price
		expectedErr       error
	}{
		{
			testName:          "normal use test",
			initInMemoryPrice: []Price{{Name: "item", Amount: 10.20}},
			itemName:          "item",
			priceWant:         Price{Name: "item", Amount: 10.20},
			expectedErr:       nil,
		},
		{
			testName:          "test error case price not found",
			initInMemoryPrice: []Price{},
			itemName:          "item",
			priceWant:         Price{},
			expectedErr:       ErrPriceNotFound,
		},
	}
	for _, tt := range testTable {
		i, err := NewInMemoryPrice(tt.initInMemoryPrice...)
		if err != nil {
			t.Fatal("error encountered when none expected")
		}

		got, err := i.GetItem(tt.itemName)
		// This is to catch hidden errors
		if err != tt.expectedErr {
			t.Fatalf("test: %s\n unexpected error encountered", tt.testName)
		}

		// Fail correctly: the deep equal will error if a error case
		// makes it through.
		if tt.expectedErr != nil {
			break
		}

		if !reflect.DeepEqual(*got, tt.priceWant) {
			t.Errorf("got %q want %q", got.String(), tt.priceWant.String())
		}

	}
}

func TestNewInMemoryPrice(t *testing.T) {
	testTable := []struct {
		testName    string
		initStorage []Price
		storageWant map[string]Price
		expectedErr error
	}{
		{
			testName:    "normal use test",
			initStorage: []Price{{Name: "item", Amount: 10.20}},
			storageWant: map[string]Price{"item": {Name: "item", Amount: 10.20}},
			expectedErr: nil,
		},
		{
			testName:    "normal use test",
			initStorage: []Price{{Name: "item", Amount: 10.20}, {Name: "item", Amount: 10.20}},
			storageWant: map[string]Price{},
			expectedErr: ErrPriceExists,
		},
	}

	for _, tt := range testTable {
		got, err := NewInMemoryPrice(tt.initStorage...)

		// This is to catch hidden errors
		if err != tt.expectedErr {
			t.Fatalf("test: %s\n unexpected error encountered", tt.testName)
		}
		if tt.expectedErr != nil {
			break
		}

		if !reflect.DeepEqual(got.Storage, tt.storageWant) {
			t.Errorf("got %v want %v", got, tt.storageWant)
		}

	}

}

func TestInMemoryGetAmount(t *testing.T) {
	testTable := []struct {
		testName          string
		initInMemoryPrice []Price
		itemName          string
		amountWant        float64
		expectedErr       error
	}{
		{
			testName:          "normal use test",
			initInMemoryPrice: []Price{{Name: "item", Amount: 10.20}},
			itemName:          "item",
			amountWant:        10.20,
			expectedErr:       nil,
		},
		{
			testName:          "price not found error",
			initInMemoryPrice: []Price{},
			itemName:          "item",
			amountWant:        0.0,
			expectedErr:       ErrPriceNotFound,
		},
	}
	for _, tt := range testTable {
		i, err := NewInMemoryPrice(tt.initInMemoryPrice...)
		if err != nil {
			t.Fatal("error encountered when none expected")
		}

		got, err := i.GetAmount(tt.itemName)
		// This is to catch hidden errors
		if err != tt.expectedErr {
			t.Fatalf("test: %s\n unexpected error encountered", tt.testName)
		}

		// Fail correctly: the deep equal will error if a error case
		// makes it through.
		if tt.expectedErr != nil {
			break
		}

		if got != tt.amountWant {
			t.Errorf("got %.2f want %.2f", got, tt.amountWant)
		}

	}
}

func TestInMemoryGetAll(t *testing.T) {
	testTable := []struct {
		testName          string
		initInMemoryPrice []Price
		want              []*Price
		expectedErr       error
	}{
		{
			testName:          "single price in storage",
			initInMemoryPrice: []Price{{Name: "item", Amount: 10.20}},
			want:              []*Price{{Name: "item", Amount: 10.20}},
			expectedErr:       nil,
		},
		{
			testName: "multiple price's in storage",
			initInMemoryPrice: []Price{
				{Name: "one", Amount: 10.20},
				{Name: "two", Amount: 10.22},
				{Name: "three", Amount: 10.23},
			},
			want: []*Price{
				{Name: "one", Amount: 10.20},
				{Name: "two", Amount: 10.22},
				{Name: "three", Amount: 10.23},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testTable {
		i, err := NewInMemoryPrice(tt.initInMemoryPrice...)
		if err != nil {
			t.Fatal("unexpected error detected")
		}
		prices, err := i.GetAll()
		if err != tt.expectedErr {
			t.Fatalf("test: %s\n unexpected error encountered", tt.testName)
		}
		if tt.expectedErr != nil {
			break
		}

		if !reflect.DeepEqual(prices, tt.want) {
			t.Errorf("got %v want %v", prices, tt.want)
		}
	}
}
