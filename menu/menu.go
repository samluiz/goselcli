package menu

import (
	"fmt"
	"log"
	"os"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/buger/goterm"
)

type Menu struct {
	Prompt    string
	CursorPos int
	Options   []*Option
}

type Option struct {
	Text string
	ID   string
}

func NewMenu(prompt string) *Menu {
	return &Menu{
		Prompt:  prompt,
		Options: make([]*Option, 0),
	}
}

func (m *Menu) AddOption(name string, id string) *Menu {
	option := &Option{
		Text: name,
		ID:   id,
	}

	m.Options = append(m.Options, option)
	return m
}

func (m *Menu) renderMenuOptions(rerender bool) {
	if rerender && len(m.Options) > 1 {
		for i := 0; i < len(m.Options); i++ {
			fmt.Printf("\033[2K")
		}
		fmt.Println("\033[?25l")
		fmt.Printf("\033[%dA", len(m.Options))
	}

	for i, option := range m.Options {
		var newline = "\n"
		if i == len(m.Options)-1 {
			newline = ""
		}

		optionText := option.Text
		cursor := " "
		if i == m.CursorPos {
			cursor = goterm.Color(">", goterm.BLUE)
			optionText = goterm.Color(optionText, goterm.BLUE)
		}

		fmt.Printf("\r%s %s%s", cursor, optionText, newline)
	}
}

func (m *Menu) Display() string {
	defer func() {
		fmt.Printf("\033[?25h")
	}()

	if len(m.Options) == 0 {
		log.Fatal("No options added to menu")
	}

	var option string

	fmt.Printf("\n%s\n", goterm.Color(goterm.Bold(m.Prompt)+":", goterm.CYAN))

	m.renderMenuOptions(false)

	for {
		keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			switch key.Code {
			case keys.Up:
				m.moveCursorUp()
			case keys.Down:
				m.moveCursorDown()
			case keys.RuneKey:
				if key.String() == "k" {
					m.moveCursorUp()
				} else if key.String() == "j" {
					m.moveCursorDown()
				} else if key.String() == "q" {
					os.Exit(0)
				} else {
					return false, nil
				}
			case keys.Enter:
				option = m.Options[m.CursorPos].ID
				fmt.Println("\r")
				return true, nil
			case keys.CtrlC, keys.CtrlD, keys.CtrlX, keys.Escape:
				os.Exit(0)
			}
			return false, nil
		})
		return option
	}
}

func (m *Menu) moveCursorUp() {
	m.CursorPos = (m.CursorPos + len(m.Options) - 1) % len(m.Options)
	m.renderMenuOptions(true)
}

func (m *Menu) moveCursorDown() {
	m.CursorPos = (m.CursorPos + 1) % len(m.Options)
	m.renderMenuOptions(true)
}
