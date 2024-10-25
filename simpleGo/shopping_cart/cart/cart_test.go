package cart

import (
	"reflect"
	"testing"
)

func TestCartAddItem(t *testing.T) {
	testTable := []struct {
		testName      string
		initItems     []map[string]int
		wantItems     map[string]int
		itemAdd       string
		numberAdd     int
		expectedError error
	}{
		{
			testName:      "basic add test",
			initItems:     []map[string]int{},
			wantItems:     map[string]int{"item": 2},
			itemAdd:       "item",
			numberAdd:     2,
			expectedError: nil,
		},
		{
			testName:      "add to existing value",
			initItems:     []map[string]int{{"item": 3}},
			wantItems:     map[string]int{"item": 5},
			itemAdd:       "item",
			numberAdd:     2,
			expectedError: nil,
		},
		{
			testName:      "negative number error",
			initItems:     []map[string]int{{"item": 3}},
			wantItems:     map[string]int{"item": 5},
			itemAdd:       "item",
			numberAdd:     -2,
			expectedError: ErrNegativeInput,
		},
	}

	for _, tt := range testTable {
		cart, _ := NewCart(tt.initItems...)
		err := cart.AddItem(tt.itemAdd, tt.numberAdd)

		if tt.expectedError != nil {
			// This pattern might be better than
			// version in price_test.go
			if err == tt.expectedError {
				// don't test next case due to error breaks
				break
			} else {
				t.Fatalf("test: %s \n unexpected error: \ngot %s want %s",
					tt.testName,
					err.Error(),
					tt.expectedError.Error())
			}
		}
		got := cart.items
		want := tt.wantItems

		if !reflect.DeepEqual(got, want) {
			t.Errorf("test: %s \ngot %v want %v", tt.testName, got, want)
		}

	}
}

func TestCartDeleteItem(t *testing.T) {
	testTable := []struct {
		testName   string
		initItems  []map[string]int
		deleteItem string
		wantItems  map[string]int
	}{
		{
			testName: "basic functionality",
			initItems: []map[string]int{
				{"item": 4},
				{"other": 2},
			},
			deleteItem: "item",
			wantItems:  map[string]int{"other": 2},
		},
	}
	for _, tt := range testTable {
		cart, _ := NewCart(tt.initItems...)
		cart.DeleteItem(tt.deleteItem)
		got := cart.items
		want := tt.wantItems

		if !reflect.DeepEqual(got, want) {
			t.Errorf("test: %s \ngot %v want %v", tt.testName, got, want)
		}
	}
}

func TestNewCart(t *testing.T) {
	testTable := []struct {
		testName  string
		initItems []map[string]int
		wantItems map[string]int
	}{
		{
			testName: "multiple of same item in variadic input",
			initItems: []map[string]int{
				{"item": 2},
				{"fun": 5},
				{"item": 4}},
			wantItems: map[string]int{
				"item": 6,
				"fun":  5,
			},
		},
		{
			testName:  "empty init test",
			initItems: []map[string]int{},
			wantItems: map[string]int{},
		},
	}

	for _, tt := range testTable {
		// error case tested in previos test.
		got, err := NewCart(tt.initItems...)
		// check for unexpected error
		if err != nil {
			t.Fatalf("test: %s error encountered when none expected", tt.testName)
		}
		want := tt.wantItems
		if !reflect.DeepEqual(got.items, want) {
			t.Errorf("test: %s \ngot %v want %v", tt.testName, got.items, want)
		}
	}
}

func TestViewItems(t *testing.T) {
	testTable := []struct {
		testName  string
		initItems []map[string]int
		wantItems map[string]int
	}{
		{
			testName: "basic test",
			initItems: []map[string]int{
				{"item": 2},
				{"fun": 5},
			},
			wantItems: map[string]int{
				"item": 2,
				"fun":  5,
			},
		},
	}

	for _, tt := range testTable {
		got, err := NewCart(tt.initItems...)
		if err != nil {
			t.Fatalf("test: %s error encountered when none expected", tt.testName)
		}

		if reflect.DeepEqual(got, tt.wantItems) {
			t.Errorf("test: %s \ngot %v want %v", tt.testName, got.items, tt.wantItems)
		}
	}
}
