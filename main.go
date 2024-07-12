package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
)

const padding = 10

type Todo struct {
	text string
	done bool
}

type App struct {
	screen      tcell.Screen
	todos       []Todo
	cursorY     int
	insertMode  bool
	currentText string
	bgColor     tcell.Color
	fgColor     tcell.Color
	transparent bool
}

func newApp() *App {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	bgColor := tcell.NewRGBColor(10, 14, 20)
	fgColor := tcell.NewRGBColor(179, 177, 173)

	return &App{
		screen:      screen,
		todos:       []Todo{},
		cursorY:     0,
		insertMode:  false,
		bgColor:     bgColor,
		fgColor:     fgColor,
		transparent: true,
	}
}

func (a *App) draw() {
	a.screen.Clear()
	style := tcell.StyleDefault.Foreground(a.fgColor)
	if !a.transparent {
		style = style.Background(a.bgColor)
	}
	// if highlight {
	// 	style = style.Reverse(true)
	// }

	// Draw border
	width, height := a.screen.Size()
	drawBox(a.screen, padding-1, padding-1, width-padding, height-padding, style)

	for i, todo := range a.todos {
		if i >= height-2*padding {
			break // Don't draw beyond the visible area
		}
		status := "[ ]"
		if todo.done {
			status = "[x]"
		}
		row := fmt.Sprintf("%s %s", status, todo.text)
		drawText(a.screen, padding, padding+i, style, row)
	}

	if a.insertMode {
		drawText(a.screen, padding, padding+len(a.todos), style, "> "+a.currentText)
	}

	a.screen.ShowCursor(padding, padding+a.cursorY)
	a.screen.Show()
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	for x := x1; x <= x2; x++ {
		s.SetContent(x, y1, tcell.RuneHLine, nil, style)
		s.SetContent(x, y2, tcell.RuneHLine, nil, style)
	}
	for y := y1 + 1; y < y2; y++ {
		s.SetContent(x1, y, tcell.RuneVLine, nil, style)
		s.SetContent(x2, y, tcell.RuneVLine, nil, style)
	}
	s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
}

func drawText(s tcell.Screen, x, y int, style tcell.Style, text string) {
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, style)
	}
}

func (a *App) handleInput() {
	for {
		ev := a.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				a.insertMode = false
				a.currentText = ""
			} else if a.insertMode {
				a.handleInsertMode(ev)
			} else {
				a.handleNormalMode(ev)
			}
		case *tcell.EventResize:
			a.screen.Sync()
		}
		a.draw()
	}
}

func (a *App) handleInsertMode(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyEnter:
		if a.currentText != "" {
			a.todos = append(a.todos, Todo{text: a.currentText, done: false})
			a.currentText = ""
			a.insertMode = false
			a.cursorY = len(a.todos) - 1
		}
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if len(a.currentText) > 0 {
			a.currentText = a.currentText[:len(a.currentText)-1]
		}
	default:
		if ev.Rune() != 0 {
			a.currentText += string(ev.Rune())
		}
	}
}

func (a *App) handleNormalMode(ev *tcell.EventKey) {
	switch ev.Rune() {
	case 'q':
		a.screen.Fini()
		os.Exit(0)
	case 'j':
		if a.cursorY < len(a.todos)-1 {
			a.cursorY++
		}
	case 'k':
		if a.cursorY > 0 {
			a.cursorY--
		}
	case 'i':
		a.insertMode = true
		a.cursorY = len(a.todos)
	case ' ':
		if a.cursorY < len(a.todos) {
			a.todos[a.cursorY].done = !a.todos[a.cursorY].done
		}
	}
}

func main() {
	// Clear the terminal
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	app := newApp()
	defer app.screen.Fini()

	app.draw()
	app.handleInput()
}
