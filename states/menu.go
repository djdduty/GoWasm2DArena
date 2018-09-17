package states

import (
	"djdduty/gowasm/core"
	"syscall/js"
	"time"
)

// MenuState ...
type MenuState struct {
	*MenuContext
}

// MenuContext ...
type MenuContext struct {
	ctx  js.Value
	x    float32
	y    float32
	game *core.Game
}

// NewMenuState creates a new menu state context
func NewMenuState(ctx js.Value, game *core.Game) MenuState {
	context := &MenuContext{
		ctx:  ctx,
		game: game,
		x:    0,
		y:    0,
	}

	return MenuState{context}
}

// Init ...
func (state MenuState) Init() error {
	return nil
}

// Update ...
func (state MenuState) Update(dt time.Duration) error {
	if state.game.Input.KeyDown("Space") {
		gameState := NewGameState(state.ctx, state.game)
		gameState.Init()
		state.game.SetState(gameState)
	}
	return nil
}

// Render ...
func (state MenuState) Render(dt time.Duration) error {
	//fmt.Println(fmt.Sprintf("%f %f", state.x, state.y))
	state.ctx.Call("clearRect", 0, 0, state.game.Width, state.game.Height)
	state.ctx.Set("textAlign", "center")
	state.ctx.Set("font", "48px Arial")
	state.ctx.Call("fillText", "Press Space to play", state.game.Width/2, state.game.Height/2)
	return nil
}
