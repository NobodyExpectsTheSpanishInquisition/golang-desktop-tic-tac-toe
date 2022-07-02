package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Game struct {
	window *window
}

func (g Game) Run() {
	g.renderMainMenu()

	g.window.window.ShowAndRun()
}

func (g Game) renderMainMenu() {
	title := canvas.NewText("TIC-TAC-TOE", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 42

	symbol := canvas.NewText("X-O-X", color.White)
	symbol.Alignment = fyne.TextAlignCenter
	symbol.TextSize = 100

	startBtn := widget.NewButton(
		"Start", func() {
		},
	)

	exitBtn := widget.NewButton(
		"Exit", func() {
			g.exit()
		},
	)

	c := container.NewVBox(
		title,
		symbol,
		layout.NewSpacer(),
		container.NewCenter(container.NewVBox(startBtn, exitBtn)),
		layout.NewSpacer(),
	)
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
