package states

import (
	"djdduty/gowasm/core"
	"djdduty/gowasm/gmath"
	"syscall/js"
	"time"
)

// GameState ...
type GameState struct {
	*GameContext
}

// GameContext ...
type GameContext struct {
	ctx  js.Value
	game *core.Game
	x    float32
	y    float32
	img  js.Value
}

// NewGameState creates a new menu state context
func NewGameState(ctx js.Value, game *core.Game) GameState {
	context := &GameContext{
		ctx:  ctx,
		game: game,
		x:    float32(game.Width) * 0.5,
		y:    float32(game.Height) * 0.5,
	}
	return GameState{context}
}

// Init ...
func (state GameState) Init() error {
	img := state.game.Document.Call("createElement", "img")
	img.Set("src", "smiley.png")
	state.img = img
	return nil
}

// Update ...
func (state GameState) Update(dt time.Duration) error {
	deltaSeconds := float64(dt/time.Microsecond) * 0.000001
	velocity := gmath.Vector2{}

	if state.game.Input.KeyDown("ArrowUp") {
		velocity.Y -= 150
	}

	if state.game.Input.KeyDown("ArrowDown") {
		velocity.Y += 150
	}

	if state.game.Input.KeyDown("ArrowLeft") {
		velocity.X -= 150
	}

	if state.game.Input.KeyDown("ArrowRight") {
		velocity.X += 150
	}

	if state.game.Input.KeyDown("Escape") {
		menuState := NewMenuState(state.ctx, state.game)
		state.game.SetState(menuState)
	}

	state.x += velocity.X * float32(deltaSeconds)
	state.y += velocity.Y * float32(deltaSeconds)
	return nil
}

// Render ...
func (state GameState) Render(dt time.Duration) error {
	//fmt.Println("rendering game state")
	state.ctx.Call("clearRect", 0, 0, state.game.Width, state.game.Height)
	state.ctx.Call("drawImage", state.img, state.x-50, state.y-50, 100, 100)
	return nil
}
