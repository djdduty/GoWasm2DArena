package core

// InputManager ...
type InputManager struct {
	keys map[string]bool
}

// SetKeyState ...
func (inputManager *InputManager) SetKeyState(keyCode string, isDown bool) {
	inputManager.keys[keyCode] = isDown
}

// KeyDown ...
func (inputManager *InputManager) KeyDown(keyCode string) bool {
	return inputManager.keys[keyCode]
}

// NewInputManager ...
func NewInputManager() *InputManager {
	return &InputManager{
		keys: make(map[string]bool),
	}
}
