package core

import (
	"syscall/js"
	"time"
)

// Game ...
type Game struct {
	currentState State
	Width        int
	Height       int
	Input        *InputManager
	Document     js.Value
}

// SetState set the current state of the game
func (game *Game) SetState(newState State) error {
	game.currentState = newState
	return nil
}

// Update ...
func (game *Game) Update(dt time.Duration) error {
	return game.currentState.Update(dt)
}

// Render ...
func (game *Game) Render(dt time.Duration) error {
	return game.currentState.Render(dt)
}
