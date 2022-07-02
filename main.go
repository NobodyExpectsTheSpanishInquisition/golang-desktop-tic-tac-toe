package main

import (
	"TicTacToe/src"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("TIC-TAC-TOE")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 400))
	w.SetFixedSize(true)

	game := src.NewGame(src.NewWindow(w))

	game.Run()
}
