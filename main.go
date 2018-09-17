package main

import (
	"djdduty/gowasm/core"
	"djdduty/gowasm/states"
	"syscall/js"
	"time"
)

func main() {
	// Fetch DOM
	document := js.Global().Get("document")
	body := document.Get("body")

	// Setup canvas
	canvasElement := document.Call("getElementById", "SomeCanvas")
	width := body.Get("clientWidth").Int()
	height := body.Get("clientHeight").Int()
	canvasElement.Set("width", width)
	canvasElement.Set("height", height)

	// Get context
	context := canvasElement.Call("getContext", "2d")

	// Initial frame time value
	lastFrameTime := time.Now()

	// Channel to extend program life cycle to browser
	done := make(chan struct{}, 0)

	inputManager := core.NewInputManager()
	game := &core.Game{
		Width:    width,
		Height:   height,
		Input:    inputManager,
		Document: document,
	}

	initialState := states.NewMenuState(context, game)

	game.SetState(initialState)

	var renderFrame js.Callback
	renderFrame = js.NewCallback(func(args []js.Value) {
		deltaTime := time.Since(lastFrameTime)
		lastFrameTime = time.Now()

		game.Update(deltaTime)
		game.Render(deltaTime)
		js.Global().Call("requestAnimationFrame", renderFrame)
	})
	defer renderFrame.Release()

	onKeyDown := js.NewCallback(func(args []js.Value) {
		evt := args[0]
		//fmt.Println(evt.Get("code").String())
		inputManager.SetKeyState(evt.Get("code").String(), true)
	})
	defer onKeyDown.Release()

	onKeyUp := js.NewCallback(func(args []js.Value) {
		evt := args[0]
		inputManager.SetKeyState(evt.Get("code").String(), false)
	})
	defer onKeyUp.Release()

	document.Call("addEventListener", "keydown", onKeyDown)
	document.Call("addEventListener", "keyup", onKeyUp)
	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
