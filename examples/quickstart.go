package examples

import "github.com/samluiz/goselcli/menu"

func Start() {
	m := menu.NewMenu("Select an option!")

	m.AddOption("Option 1", "1")
	m.AddOption("Option 2", "2")
	m.AddOption("Option 3", "3")
	m.AddOption("Option 4", "4")
	m.AddOption("Option 5", "5")

	choice := m.Display()

	println("You chose:", choice)
}
