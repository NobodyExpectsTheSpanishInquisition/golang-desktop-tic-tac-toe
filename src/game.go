package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Game struct {
	window *window
}

func (g Game) Run() {
	g.renderMainMenu()

	g.window.window.ShowAndRun()
}

func (g Game) renderMainMenu() {
	title := widget.NewTextGrid()
	title.SetText("TIC-TAC-TOE")

	exitBtn := widget.NewButton(
		"Exit", func() {
			g.exit()
		},
	)

	c := container.NewVBox(title, exitBtn)
	g.window.window.SetContent(c)
}

func (g Game) exit() {
	g.window.window.Close()
}

func NewGame(window *window) *Game {
	return &Game{window: window}
}

type window struct {
	window fyne.Window
}

func NewWindow(w fyne.Window) *window {
	return &window{window: w}
}
