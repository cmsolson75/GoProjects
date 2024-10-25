package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func testingPriceMap() map[string]float64 {
	return map[string]float64{
		"apple":   10.25,
		"orange":  4.00,
		"bannana": 3.00,
	}
}

func TestNewPriceTable(t *testing.T) {
	priceTable := testingPriceMap()
	pt := NewPriceTable(priceTable)

	if !reflect.DeepEqual(priceTable, pt.priceTable) {
		t.Errorf("\ngot %s, \nwant %s", prettyMapPrint(pt.priceTable), prettyMapPrint(priceTable))
	}
}

func TestGetValue(t *testing.T) {
	priceMap := testingPriceMap()
	pt := NewPriceTable(priceMap)

	got := pt.GetValue("apple")
	want := 10.25

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestCartAddItem(t *testing.T) {
	c := Cart{}
	c.AddItem("apple")
	c.AddItem("orange")
	got := c.items
	want := []string{"apple", "orange"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCartCheckout(t *testing.T) {
	cart := Cart{items: []string{"apple", "orange", "bannana"}}
	priceMap := testingPriceMap()
	p := NewPriceTable(priceMap)

	got := cart.Checkout(p)

	want := 17.25

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestUserInput(t *testing.T) {
	want := "apple"
	mockStdin := bytes.NewBufferString(fmt.Sprintf("%s\n", want))
	reader := bufio.NewReader(mockStdin)
	got, _ := UserInput(reader, "", &bytes.Buffer{})
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func prettyMapPrint(m map[string]float64) string {
	data, _ := json.MarshalIndent(m, "", " ")
	return string(data)
}
