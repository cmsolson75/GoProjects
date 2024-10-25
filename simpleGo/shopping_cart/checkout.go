package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/charmbracelet/huh"
)

type ItemTable interface {
	GetValue(string) float64
}

type PriceTable struct {
	priceTable map[string]float64
}

func NewPriceTable(priceMap map[string]float64) *PriceTable {
	// return reference due to the size of priceTable being potentially huge
	return &PriceTable{priceTable: priceMap}
}

func (p *PriceTable) GetValue(itemName string) float64 {
	return p.priceTable[itemName]
}

func (p *PriceTable) Exists(item string) bool {
	_, ok := p.priceTable[item]
	return ok
}

// if adding in multiple items: could store it in dict (string: int)
type Cart struct {
	items []string
}

func (c *Cart) AddItem(item string) {
	c.items = append(c.items, item)
}

func (c *Cart) Checkout(itemTable ItemTable) float64 {
	var total float64
	for _, item := range c.items {
		total += itemTable.GetValue(item)
	}
	return total
}

var PMap = map[string]float64{
	"apple":  10.25,
	"orange": 4.00,
	"banana": 3.00,
}

func UserInput(reader *bufio.Reader, message string, out io.Writer) (string, error) {
	fmt.Fprint(out, message)
	item, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(item), nil
}

var priceTable = NewPriceTable(PMap)

func printPriceMap(priceMap map[string]float64) string {
	outputBuffer := bytes.Buffer{}

	w := tabwriter.NewWriter(&outputBuffer, 10, 0, 3, ' ', 0)
	fmt.Fprint(&outputBuffer, "Items you can buy\n")
	fmt.Fprint(&outputBuffer, "------------------\n")
	for key, value := range priceMap {
		fmt.Fprintf(w, "%s\t$%.2f\t\n", key, value)
		w.Flush()
	}
	fmt.Fprint(&outputBuffer, "\n")
	return outputBuffer.String()
}

func (c *PriceTable) FetchItems() []string {
	var output []string
	for name := range c.priceTable {
		output = append(output, name)
	}
	sort.Strings(output)

	return output

}

var fruit string

func main() {
	title := `

███████╗██████╗ ██╗   ██╗██╗████████╗    ███████╗████████╗ █████╗ ███╗   ██╗██████╗
██╔════╝██╔══██╗██║   ██║██║╚══██╔══╝    ██╔════╝╚══██╔══╝██╔══██╗████╗  ██║██╔══██╗
█████╗  ██████╔╝██║   ██║██║   ██║       ███████╗   ██║   ███████║██╔██╗ ██║██║  ██║
██╔══╝  ██╔══██╗██║   ██║██║   ██║       ╚════██║   ██║   ██╔══██║██║╚██╗██║██║  ██║
██║     ██║  ██║╚██████╔╝██║   ██║       ███████║   ██║   ██║  ██║██║ ╚████║██████╔╝
╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝   ╚═╝       ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝

	`
	huh.NewNote().
		Title(title).
		Description(printPriceMap(PMap)).
		Next(true).
		NextLabel("Continue").Run()
	cart := Cart{}
	items := priceTable.FetchItems()
	items = append(items, "checkout")

	for {
		huh.NewSelect[string]().
			Title("Pick Fruit").
			OptionsFunc(func() []huh.Option[string] {
				return huh.NewOptions(items...)
			}, "").Value(&fruit).Run()

		if fruit == "checkout" {
			break
		}

		// this should be handled better
		if priceTable.Exists(fruit) {
			fmt.Printf("Adding %s to your cart\n", fruit)
			cart.AddItem(fruit)
		} else {
			fmt.Printf("We don't sell %q\n", fruit)
		}

	}

	output := cart.Checkout(priceTable)
	fmt.Printf("You are charged $%.2f\n", output)

}
