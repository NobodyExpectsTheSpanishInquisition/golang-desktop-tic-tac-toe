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
	play   *Play
}

func NewGame(window *window) *Game {
	return &Game{window: window}
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
	g.play = NewPlay()
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

	fieldOne = widget.NewButton("", func() { g.makeMove(fieldOne, "O", 1, true) })
	fieldTwo = widget.NewButton("", func() { g.makeMove(fieldTwo, "O", 2, true) })
	fieldThree = widget.NewButton("", func() { g.makeMove(fieldThree, "O", 3, true) })
	fieldFour = widget.NewButton("", func() { g.makeMove(fieldFour, "O", 4, true) })
	fieldFive = widget.NewButton("", func() { g.makeMove(fieldFive, "O", 5, true) })
	fieldSix = widget.NewButton("", func() { g.makeMove(fieldSix, "O", 6, true) })
	fieldSeven = widget.NewButton("", func() { g.makeMove(fieldSeven, "O", 7, true) })
	fieldEight = widget.NewButton("", func() { g.makeMove(fieldEight, "O", 8, true) })
	fieldNine = widget.NewButton("", func() { g.makeMove(fieldNine, "O", 9, true) })

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

func (g Game) makeMove(field *widget.Button, symbol string, fieldNum int, playerMove bool) {
	markField(field, symbol)
	if false == playerMove {
		g.play.MakeComputerMove(fieldNum)
		return
	}

	g.play.MakePlayerMove(fieldNum)
}

func markField(field *widget.Button, symbol string) {
	if "" == field.Text {
		field.SetText(symbol)
	}
}

type window struct {
	window fyne.Window
}

func NewWindow(w fyne.Window) *window {
	return &window{window: w}
}
