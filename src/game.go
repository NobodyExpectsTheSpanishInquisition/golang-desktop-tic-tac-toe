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

	fieldOne = widget.NewButton("", func() { g.makeMove(0) })
	fieldTwo = widget.NewButton("", func() { g.makeMove(1) })
	fieldThree = widget.NewButton("", func() { g.makeMove(2) })
	fieldFour = widget.NewButton("", func() { g.makeMove(3) })
	fieldFive = widget.NewButton("", func() { g.makeMove(4) })
	fieldSix = widget.NewButton("", func() { g.makeMove(5) })
	fieldSeven = widget.NewButton("", func() { g.makeMove(6) })
	fieldEight = widget.NewButton("", func() { g.makeMove(7) })
	fieldNine = widget.NewButton("", func() { g.makeMove(8) })

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

	fieldBox := container.NewGridWithColumns(
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

func (g Game) makeMove(fieldNum int) {
	g.play.MakePlayerMove(fieldNum)
	if true == g.play.DidPlayerWin(fieldNum) {
		g.window.window.SetContent(container.NewWithoutLayout(widget.NewLabel("YOU WON")))
		return
	}
	g.play.MakeComputerMove(g.play.GetComputerField())
	if true == g.play.DidComputerWin(fieldNum) {
		g.window.window.SetContent(container.NewWithoutLayout(widget.NewLabel("YOU LOSE")))
		return
	}
}

type window struct {
	window fyne.Window
}

func NewWindow(w fyne.Window) *window {
	return &window{window: w}
}
