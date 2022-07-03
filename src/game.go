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

	fieldOne = widget.NewButton("", func() { g.makeMove(0, true) })
	fieldTwo = widget.NewButton("", func() { g.makeMove(1, true) })
	fieldThree = widget.NewButton("", func() { g.makeMove(2, true) })
	fieldFour = widget.NewButton("", func() { g.makeMove(3, true) })
	fieldFive = widget.NewButton("", func() { g.makeMove(4, true) })
	fieldSix = widget.NewButton("", func() { g.makeMove(5, true) })
	fieldSeven = widget.NewButton("", func() { g.makeMove(6, true) })
	fieldEight = widget.NewButton("", func() { g.makeMove(7, true) })
	fieldNine = widget.NewButton("", func() { g.makeMove(8, true) })

	g.play.SetFields(
		[9]*Field{
			NewField(0, fieldOne),
			NewField(1, fieldTwo),
			NewField(2, fieldThree),
			NewField(3, fieldFour),
			NewField(4, fieldFive),
			NewField(5, fieldSix),
			NewField(6, fieldSeven),
			NewField(7, fieldEight),
			NewField(8, fieldNine),
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

func (g Game) makeMove(fieldNum int, playerMove bool) {
	if false == playerMove {
		g.play.MakeComputerMove(fieldNum)
		return
	}

	g.play.MakePlayerMove(fieldNum)
}

type window struct {
	window fyne.Window
}

func NewWindow(w fyne.Window) *window {
	return &window{window: w}
}
