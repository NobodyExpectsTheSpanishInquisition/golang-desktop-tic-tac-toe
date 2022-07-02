package main

import (
	"TicTacToe/src"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("TIC-TAC-TOE")
	w.CenterOnScreen()
	w.SetFullScreen(true)

	game := src.NewGame(src.NewWindow(w))

	game.Run()
}
