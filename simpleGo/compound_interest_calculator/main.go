package main

import (
	"fmt"
	"strconv"

	"github.com/cmsolson75/GoProjects/simpleGo/compound_interest_calculator/calc"
	"github.com/rivo/tview"
)

func inputHandler(principle, rate, time string) string {
	p, _ := strconv.ParseFloat(principle, 64)
	r, _ := strconv.ParseFloat(rate, 64)
	t, _ := strconv.ParseFloat(time, 64)

	amount, _ := calc.CompoundInterestCalculation(p, r, t)
	return fmt.Sprintf("In %s years, you will have $%.2f", time, amount)
}

var (
	principle string
	rate      string
	time      string
)

func main() {
	app := tview.NewApplication()
	outputView := tview.NewTextView().SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	appDataForm := tview.NewForm()
	appDataForm.AddTextView("Step 1:", "Starting Investment", 400, 1, false, false).
		AddInputField("principle:", principle, 20, nil, func(text string) {
			principle = text
		}).AddTextView("Step 2:", "Interest Rate", 20, 1, false, false).
		AddInputField("rate (%):", rate, 20, nil, func(text string) {
			rate = text
		}).AddTextView("Step 3:", "Investment Time", 20, 1, false, false).
		AddInputField("time:", time, 20, nil, func(text string) {
			time = text
		}).
		AddButton("submit", func() {
			famount := inputHandler(principle, rate, time)
			outputView.SetText(famount)
		})

	outputView.SetBorder(true).SetTitle("Output")
	appDataForm.SetBorder(true).SetTitle("Input Data")

	flex := tview.NewFlex().
		// AddItem(appDataForm, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(appDataForm, 0, 1, true).
			AddItem(outputView, 0, 2, false), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Table"), 0, 1, false), 0, 2, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
