package main

import (
	"fmt"
	"strconv"

	"github.com/cmsolson75/GoProjects/simpleGo/compound_interest_calculator/calc"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getTableData(principle, rate, time float64) [][]string {
	var tableData [][]string

	tableData = append(tableData, []string{"Year 0", fmt.Sprintf("$%.2f", principle)})

	for t := 1; t < int(time)+1; t++ {
		a, _ := calc.CompoundInterestCalculation(principle, rate, float64(t))
		tableData = append(tableData, []string{fmt.Sprintf("Year %d", t), fmt.Sprintf("$%.2f", a)})
	}

	return tableData
}

// This should be a struct with a method on it.
func inputHandler(principle, rate, time string) string {
	p, _ := strconv.ParseFloat(principle, 64)
	r, _ := strconv.ParseFloat(rate, 64)
	t, _ := strconv.ParseFloat(time, 64)

	amount, _ := calc.CompoundInterestCalculation(p, r, t)
	return fmt.Sprintf("In %s years, you will have $%.2f", time, amount)
}

// make struct
var (
	principle string
	rate      string
	time      string
)

// struct methods: init from strings
// struct methods: convert to ammount

// app struct
// app, form, table, view, this is simple to do.
// This means you have a bunch of methods off of this that make construction EASY.

type DataTable struct {
	table *tview.Table
}

func (d *DataTable) constructTable(headers []string, data [][]string) {

	n := len(headers)
	for i := 0; i < n; i++ {
		d.table.SetCell(0, i, tview.NewTableCell(headers[i]).SetAlign(tview.AlignLeft))
	}

	rows := len(data)

	for r := 0; r < rows; r++ {
		for c := 0; c < n; c++ {
			d.table.SetCell(r+1, c, tview.NewTableCell(data[r][c]).SetAlign(tview.AlignLeft))
		}
	}
}

// Make text const for simple UI text editing.
// title, form text (struct: p, r, t),

func main() {
	app := tview.NewApplication()
	dataTable := tview.NewTable().SetBorders(true).SetBordersColor(tcell.ColorLightCyan)
	dt := DataTable{table: dataTable}
	outputView := tview.NewTextView().SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	appDataForm := tview.NewForm()
	appDataForm.
		AddInputField("Initial Investment:", principle, 20, nil, func(text string) {
			principle = text
		}).
		AddInputField("Estimated Interest Rate:", rate, 20, nil, func(text string) {
			rate = text
		}).
		AddInputField("Length of Time in Years:", time, 20, nil, func(text string) {
			time = text
		}).
		AddButton("SUBMIT", func() {
			famount := inputHandler(principle, rate, time)
			outputView.SetText(famount)

			p, _ := strconv.ParseFloat(principle, 64)
			r, _ := strconv.ParseFloat(rate, 64)
			t, _ := strconv.ParseFloat(time, 64)

			headers := []string{"Years", fmt.Sprintf("Future Value (%.2f%%)", r)}

			data := getTableData(p, r, t)
			dt.table.Clear()
			dt.constructTable(headers, data)
		})

	outputView.SetBorder(true).SetTitle("Output").SetBorderColor(tcell.ColorPurple)
	appDataForm.SetBorder(true).SetTitle("Input Data").SetBorderColor(tcell.ColorYellow)
	dt.table.SetBorder(true).SetBorderColor(tcell.ColorDarkCyan).SetTitle("Table Output")

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(appDataForm, 0, 1, true).
			AddItem(outputView, 0, 2, false), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(dt.table, 0, 1, false), 0, 2, false)

	frame := tview.NewFrame(flex).AddText("COMPOUND INTEREST CALCULATOR", true, 1, tcell.ColorWheat)

	if err := app.SetRoot(frame, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
