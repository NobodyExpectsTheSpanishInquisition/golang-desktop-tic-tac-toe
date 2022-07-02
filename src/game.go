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
			g.start()
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

func (g Game) start() {
	g.renderPlayField()
}

func (g Game) renderPlayField() {
	var fieldOne *widget.Button
	var fieldTwo *widget.Button
	var fieldThree *widget.Button
	var fieldFour *widget.Button
	var fieldFive *widget.Button
	var fieldSix *widget.Button
	var fieldSeven *widget.Button
	var fieldEight *widget.Button
	var fieldNine *widget.Button

	fieldOne = widget.NewButton(
		"", func() {
			markField(fieldOne, "O")
		},
	)
	fieldTwo = widget.NewButton(
		"", func() {
			markField(fieldTwo, "O")
		},
	)
	fieldThree = widget.NewButton(
		"", func() {
			markField(fieldThree, "O")
		},
	)
	fieldFour = widget.NewButton(
		"", func() {
			markField(fieldFour, "O")
		},
	)
	fieldFive = widget.NewButton(
		"", func() {
			markField(fieldFive, "O")
		},
	)
	fieldSix = widget.NewButton(
		"", func() {
			markField(fieldSix, "O")
		},
	)
	fieldSeven = widget.NewButton(
		"", func() {
			markField(fieldSeven, "O")
		},
	)
	fieldEight = widget.NewButton(
		"", func() {
			markField(fieldEight, "O")
		},
	)
	fieldNine = widget.NewButton(
		"", func() {
			markField(fieldNine, "O")
		},
	)

	fieldBox := container.NewGridWithRows(
		3,
		fieldOne,
		fieldTwo,
		fieldThree,
		fieldFour,
		fieldFive,
		fieldSix,
		fieldSeven,
		fieldEight,
		fieldNine,
	)

	g.window.window.SetContent(fieldBox)
}

func markField(field *widget.Button, symbol string) {
	if "" == field.Text {
		field.SetText(symbol)
	}
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
