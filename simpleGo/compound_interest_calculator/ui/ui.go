package ui

import (
	"fmt"

	"github.com/cmsolson75/GoProjects/simpleGo/compound_interest_calculator/calculator"
	"github.com/cmsolson75/GoProjects/simpleGo/compound_interest_calculator/state"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	AppSession   *tview.Application
	Table        *tview.Table
	Form         *tview.Form
	Text         *tview.TextView
	Output       string
	State        *state.AppState
	TableData    [][]string
	TableHeaders []string
	Frame        *tview.Frame
}

func (a *App) ErrorModal(err error) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Error: %s", err.Error())).
		AddButtons([]string{"BACK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "BACK" {
				a.AppSession.SetRoot(a.Frame, true)
			}
		})
	a.AppSession.SetRoot(modal, true)
}

func (a *App) ExitModal() {
	modal := tview.NewModal().
		SetText("Are you sure you want to quit?").
		AddButtons([]string{"QUIT", "BACK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
			case "QUIT":
				a.AppSession.Stop()
			case "BACK":
				a.AppSession.SetRoot(a.Frame, true)
			}
		})
	a.AppSession.SetRoot(modal, true)
}

func (a *App) CalculateTableData() error {
	// Init Empty
	a.TableData = [][]string{}
	cid, err := calculator.InitCompoundInterestDataString(*a.State)
	if err != nil {
		return err
	}

	amount, err := cid.Compute()
	if err != nil {
		// need to return error
		return err
	}

	a.Output = fmt.Sprintf("In %s years, you will have $%.2f", a.State.TimeYears, amount)

	// Init first row with starting data
	a.TableData = append(a.TableData, []string{"Year 0", fmt.Sprintf("$%.2f", cid.Principle)})

	for t := 1; t < int(cid.Time)+1; t++ {
		amount, err := cid.ComputeWithTime(t)
		if err != nil {
			return err
		}
		a.TableData = append(a.TableData, []string{fmt.Sprintf("Year %d", t), fmt.Sprintf("$%.2f", amount)})
	}
	return nil
}

func (a *App) ConstructTable() {
	// Clear current state
	a.Table.Clear()

	a.TableHeaders = []string{"Years", fmt.Sprintf("Future Value (%s%%)", a.State.InterestRate)}

	// Calculate new state
	n := len(a.TableHeaders)
	for i := 0; i < n; i++ {
		a.Table.SetCell(0, i, tview.NewTableCell(a.TableHeaders[i]).SetAlign(tview.AlignLeft))
	}

	rows := len(a.TableData)

	for r := 0; r < rows; r++ {
		for c := 0; c < n; c++ {
			a.Table.SetCell(r+1, c, tview.NewTableCell(a.TableData[r][c]).SetAlign(tview.AlignLeft))
		}
	}
}

// Will need to return error on all of these
func (a *App) CreateForm() {
	a.Form = tview.NewForm().
		AddInputField("Initial Investment:", a.State.InitialInvestment, 20, nil, func(text string) {
			a.State.InitialInvestment = text
		}).
		AddInputField("Estimated Interest Rate:", a.State.InterestRate, 20, nil, func(text string) {
			a.State.InterestRate = text
		}).
		AddInputField("Length of Time in Years:", a.State.TimeYears, 20, nil, func(text string) {
			a.State.TimeYears = text
		}).
		AddButton("SUBMIT", func() {
			err := a.State.CheckEmptyInput()
			if err != nil {
				a.ErrorModal(err)
				return
			}
			err = a.CalculateTableData()
			if err != nil {
				a.ErrorModal(err)
				return
			}
			// if we are error free
			a.Text.SetText(a.Output)
			a.ConstructTable()
		}).
		AddButton("QUIT", func() {
			a.ExitModal()
		})
}

func NewApp() *App {
	return &App{
		State:      &state.AppState{},
		AppSession: tview.NewApplication(),
		Text: tview.NewTextView().
			SetDynamicColors(true).
			SetRegions(true).
			SetWordWrap(true),
		Table: tview.NewTable().
			SetBorders(true).
			SetBordersColor(tcell.ColorLightCyan),
	}
}

func (a *App) Run() {
	a.CreateForm()

	a.Text.SetBorder(true).SetTitle("Output").SetBorderColor(tcell.ColorPurple)
	a.Form.SetBorder(true).SetTitle("Input Data").SetBorderColor(tcell.ColorYellow)
	a.Table.SetBorder(true).SetBorderColor(tcell.ColorDarkCyan).SetTitle("Table Output")

	flexBox := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(a.Form, 0, 1, true).
			AddItem(a.Text, 0, 2, false), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(a.Table, 0, 1, false), 0, 2, false)

	a.Frame = tview.NewFrame(flexBox).AddText("COMPOUND INTEREST CALCULATOR", true, 1, tcell.ColorWheat)

	if err := a.AppSession.SetRoot(a.Frame, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
