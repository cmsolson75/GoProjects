# Compound Interest Calculator for Investment

## Goals

Learn the basics of the Math Library
Implement a TUI with tview


## TVIEW NOTES
[project github](https://github.com/rivo/tview)
[demos folder](https://github.com/rivo/tview/tree/master/demos)
[docs](https://pkg.go.dev/github.com/rivo/tview)
[tcell github](https://github.com/gdamore/tcell)
[SetRoot implementation](https://github.com/rivo/tview/blob/master/application.go#L783)
[article](https://rocket9labs.com/post/tview-and-you/)
[netris](https://code.rocket9labs.com/tslocum/netris)
[link](https://github.com/Skarlso/gtui/tree/main)
[GOOD EXAMPLE](https://github.com/broadcastle/crm/blob/master/code/tui/contact.go)
[link](https://github.com/dhulihan/grump/blob/main/ui/tracks.go)


## Refactor notes

There are alot of refactors to do. First I should not be using so many sub systems in my App struct, this is handling too much stuff, each of the systems should have there own struct. 

I think I am learning how methods should flow in Go, this will just take more practice. 

I need to start using Table Driven Tests instead of t.Run(), I like t.Run(), but setting up table driven tests will help me navigate edge cases.

I need to learn more about injecting functions, currently a lot of stuff is in line.