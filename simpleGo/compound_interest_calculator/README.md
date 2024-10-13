# Compound Interest Calculator for Investment

## Goals

Learn the basics of the Math Library
Implement a TUI with tview


## TVIEW NOTES
[project github](https://github.com/rivo/tview)

Just go through the getting started

INSTALL: `go get github.com/rivo/tview@master`

```
package main

import (
	"github.com/rivo/tview"
)

func main() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
```


[demos folder](https://github.com/rivo/tview/tree/master/demos)


What will I need
- Form Input
- Buttons: For processing the input: this will trigger event???
- Table Output: For future
- Some kind of text output for final number: This will always be there.
- Modal: For quiting the app


[docs](https://pkg.go.dev/github.com/rivo/tview)

Just start implementing the demos and seeting what stuff does.


### Box Demo

uses tcell:[github](https://github.com/gdamore/tcell)

### GENERAL NOTES

app is the main interface with the system create a new app with the following command

`app := tview.NewApplication()`

Standard
```
if app.SetRoot(thing, true).Run(); err != nil {
    panic(err)
}

```
this is in all the demos

It seems like Run is what is calling the app.
- This is true

What is SetRoot Doing?
- [implementation](https://github.com/rivo/tview/blob/master/application.go#L783)
- It says it is sets the primative of the application, I wonder what is standard here.


What is a primative? Is this like a box or something?
- So it seems like SetRoot is letting you display what tview object you want to view on the screen

SO the core of TVIEW is using a GRID, this lets you add primatives to the screen

primative is any object that tview creates

Nice Tutorial: [article](https://rocket9labs.com/post/tview-and-you/)


Core Idea: Use a Grid to build your app, pages are intresting as an adition but a grid is all I will need.

Need to spend a bit more time understanding the Grid Organization: Its a bit confusing.


```
Example of my code
package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	grid := tview.NewGrid()
	nameLabel := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetWrap(true).
		SetWordWrap(true)

	button := tview.NewButton("Hit to get a message").SetSelectedFunc(func() {
		nameLabel.SetText("You are cute!")
	})
	button.SetBorder(true)

	grid.AddItem(button, 0, 0, 2, 2, 0, 0, true)

	grid.AddItem(nameLabel, 0, 2, 2, 2, 0, 0, false)
	app.SetRoot(grid, true).EnableMouse(true)
	app.Run()
}

```

Alot of these are using width, what does that mean? I can't tell this is a mission.


```
func main() {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddTextArea("Address", "", 40, 0, 0, nil).
		AddTextView("Notes", "This is just a demo", 40, 2, true, false).
		AddCheckbox("Age 18+", false, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

```


Alot of tview uses this method chain style


Do you use event functions to set variables you use in future calculations.



Box is the core of tview: lets do the grid example to learn more about it, then read up on Box.


These notes will need to get cleaned up in the future. Thats a future problem though.


### Overview Notes 

#### The Primitive

This is the top most interface for the library, the only implemented primative is Box.


#### Widgets

Widgets are structs with embeded a Box and build on top of it.

TexView struct
- *box
- buffer []string
- align int // this is for the text alignment that has aliases in Left


Some widgets can hold other widgets in a nested manner
- Grid
- Flexbox
- Form
- List
- Pages
- Table
- TreeView

Chaining: Common for widget commands.


#### App Structure & Thread Safty

TO display a primitive (and its contents) we call SetRoot
- Arg 0: what primative is root
- Arg 1: will this be resized to fit the screen


## Example Apps
[netris](https://code.rocket9labs.com/tslocum/netris)



## General Flow

You can use GetText() and SetText(): This gives you elements


```
package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	label := tview.NewTextView().SetText("Please enter your name:")

	input := tview.NewInputField()
	// var submittedName bool

	appGrid := tview.NewGrid()

	btn := tview.NewButton("Submit").SetSelectedFunc(func() {
		name := input.GetText()
		m := tview.NewModal().
			SetText(fmt.Sprintf("Greetings, %s!", name)).
			AddButtons([]string{"done"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "done" {
				input.SetText("")
				app.SetRoot(appGrid, true)
			}
		})
		app.SetRoot(m, true)
	})

	appGrid.SetColumns(-1, 24, 16, -1).
		SetRows(-1, 2, 3, -1).
		AddItem(label, 1, 1, 1, 1, 0, 0, false).
		AddItem(input, 1, 2, 1, 1, 0, 0, false).
		AddItem(btn, 2, 1, 1, 2, 0, 0, false)

	app.EnableMouse(true)

	app.SetRoot(appGrid, true).SetFocus(input)

	err := app.Run()
	if err != nil {
		panic(err)
	}

}



```



## Current Problems

Lets just go for it; I don't understand the grid setting, but I can figure it out.

So I can't nessasaraly do it with a form yet, that is the next idea, but for now we can do it with a grid.


I need to look at Grid Implementations, and Form Implementation



I am starting to get it. Its still challenging, but I am getting it.

I have a working TUI for the app, its bad, so I will improve it but its working.


```
This is the working version

func main() {
	app := tview.NewApplication()

	principle := tview.NewTextView().SetText("principle:")
	principleInput := tview.NewInputField()

	rate := tview.NewTextView().SetText("rate:")
	rateInput := tview.NewInputField()

	time := tview.NewTextView().SetText("time:")
	timeInput := tview.NewInputField()

	appGrid := tview.NewGrid()

	btn := tview.NewButton("Submit").SetSelectedFunc(func() {
		p, _ := strconv.ParseFloat(principleInput.GetText(), 64)
		r, _ := strconv.ParseFloat(rateInput.GetText(), 64)
		t, _ := strconv.ParseFloat(timeInput.GetText(), 64)

		amount, _ := CompoundInterestCalculation(p, r, t)
		m := tview.NewModal().
			SetText(fmt.Sprintf("$%.2f", amount)).
			AddButtons([]string{"done"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "done" {
				principleInput.SetText("")
				rateInput.SetText("")
				timeInput.SetText("")
				app.SetRoot(appGrid, true)
			}
		})
		app.SetRoot(m, true)
	})

	btn.SetBorder(true)

	appGrid.SetColumns(1).
		SetRows(2, 2, 2, 2, 3).
		AddItem(principle, 1, 1, 1, 1, 0, 0, false).
		AddItem(principleInput, 1, 2, 1, 1, 0, 0, false).
		AddItem(rate, 2, 1, 1, 1, 0, 0, false).
		AddItem(rateInput, 2, 2, 1, 1, 0, 0, false).
		AddItem(time, 3, 1, 1, 1, 0, 0, false).
		AddItem(timeInput, 3, 2, 1, 1, 0, 0, false).
		AddItem(btn, 4, 1, 2, 2, 0, 0, false)

	app.EnableMouse(true)

	app.SetRoot(appGrid, true)

	err := app.Run()
	if err != nil {
		panic(err)
	}

}




```


More Go Examples
[link](https://github.com/Skarlso/gtui/tree/main)


I Need to focus on 
- Flex & Grid
- Getting Data Out of Forms


[GOOD EXAMPLE](https://github.com/broadcastle/crm/blob/master/code/tui/contact.go)



I really need to learn
- Flex & Grid!!
- This is a pain in the butt, I think I need to think of a form as just a input, 


You can wrap the App in a Struct and tie methods too it.

Will learn alot from just reading through projects that people have done in tview. This is a old package and doesnt feel ideal compaired to something like Bubble Tea, but I will keep going. I don't know if its worth using this, but there are definently alot of UI handlers I should Write. 


Take a break
- Come back: Look through flex box tutorial
- Make form: get user data
- Output: Make some output that gets the data from the form and uses it.


Okay so I have the app working; but I need some kind of output box.


I need to make a box that shows a text view on the inside. So I can put a boarder on it.