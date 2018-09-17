package core

import "time"

// State represents a scene or state in the game, I.E Menu state, game state, etc
type State interface {
	Init() error
	Update(dt time.Duration) error // dt is in microseconds
	Render(dt time.Duration) error // dt is in microseconds
}
