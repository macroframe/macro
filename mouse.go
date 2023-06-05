package macro

import (
	"github.com/go-vgo/robotgo"
	"time"
)

type mouse struct {
	mouseDelay time.Duration
}

// Mouse includes all mouse related functions.
var Mouse = mouse{}

func init() {
	Mouse.SetMouseDelay(100 * time.Millisecond)
}

// --- Config ---

func (m *mouse) SetMouseDelay(delay time.Duration) {
	m.mouseDelay = delay
	robotgo.MouseSleep = int(m.mouseDelay / time.Millisecond)
}

// --- Mouse Functions ---

// Move moves the mouse to the given coordinates.
func (m *mouse) Move(x, y int) {
	robotgo.Move(x, y)
}

// MoveSmooth moves the mouse to the given coordinates smoothly.
func (m *mouse) MoveSmooth(x, y int) {
	robotgo.MoveSmooth(x, y)
}
