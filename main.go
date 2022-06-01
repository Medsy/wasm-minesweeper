package main

import (
	"syscall/js"

	"wasm/Game"
)

var ms *Game.Minesweeper

func newGame(this js.Value, i []js.Value) interface{} {
	w := i[0].Int()
	h := i[1].Int()
	m := i[2].Int()

	ms = Game.New(w, h, m)

	return js.ValueOf(ms.Print()).String()
}

func openCell(this js.Value, i []js.Value) interface{} {
	x := i[0].Int()
	y := i[1].Int()

	_, r, _ := ms.Open(x, y)
	if r {
		println("You lost!")
	}

	return js.ValueOf(ms.Print()).String()
}

func msPrint(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(ms.Print()).String()
}

func toggleFlag(this js.Value, i []js.Value) interface{} {
	x := i[0].Int()
	y := i[1].Int()

	ms.ToggleFlag(x, y)

	return js.ValueOf(ms.Print()).String()
}

func registerCallbacks() {
	js.Global().Set("openCell", js.FuncOf(openCell))
	js.Global().Set("newGame", js.FuncOf(newGame))
	js.Global().Set("msPrint", js.FuncOf(msPrint))
	js.Global().Set("toggleFlag", js.FuncOf(toggleFlag))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")

	// register functions
	registerCallbacks()
	<-c
}
