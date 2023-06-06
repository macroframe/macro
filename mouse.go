package macro

import (
	"github.com/go-vgo/robotgo"
	"github.com/macroframe/macro/internal"
	"github.com/macroframe/macro/pkg/mouse"
	"time"
)

type mouseType struct {
	mouseDelay time.Duration
}

// Mouse includes all mouseType related functions.
var Mouse = mouseType{}

func init() {
	Mouse.SetMouseDelay(25 * time.Millisecond)
}

// --- Config ---

// SetMouseDelay sets the default delay between mouseType actions.
// If the delay is too short, the operating system might fail to register all actions.
// The default value is 25ms.
func (m *mouseType) SetMouseDelay(delay time.Duration) {
	m.mouseDelay = delay
	robotgo.MouseSleep = int(m.mouseDelay / time.Millisecond)
}

// --- Mouse Movement ---

// Move moves the mouseType to the given coordinates.
func (m *mouseType) Move(x, y int) {
	robotgo.Move(x, y)
}

// MoveSmooth moves the mouseType to the given coordinates smoothly.
func (m *mouseType) MoveSmooth(x, y int) {
	robotgo.MoveSmooth(x, y)
}

// MoveRelative moves the mouseType relative to the current position.
func (m mouseType) MoveRelative(x, y int) {
	robotgo.MoveRelative(x, y)
}

// MoveRelativeSmooth moves the mouseType relative to the current position smoothly.
func (m mouseType) MoveRelativeSmooth(x, y int) {
	robotgo.MoveSmoothRelative(x, y)
}

// Drag moves the mouseType to the given coordinates while holding the left mouseType button.
func (m mouseType) Drag(x, y int) {
	robotgo.DragSmooth(x, y)
}

// --- Mouse Clicks ---

// Click clicks a specific mouse.Button.
// A count can be given to click multiple times.
func (m mouseType) Click(mouseButton mouse.Button, count ...int) {
	sum := internal.Sum(count...)
	if sum == 0 {
		sum = 1
	}

	button := internal.MouseButtonToRobotgoString(mouseButton)

	for i := 0; i < sum; i++ {
		robotgo.Click(button)
	}
}

// Hold presses a specific mouse.Button.
// Use Release to release the button.
func (m mouseType) Hold(mouseButton mouse.Button) {
	robotgo.Toggle(internal.MouseButtonToRobotgoString(mouseButton))
}

// Release releases a specific mouse.Button.
// Use Hold to press the button.
func (m mouseType) Release(mouseButton mouse.Button) {
	robotgo.Toggle(internal.MouseButtonToRobotgoString(mouseButton), "up")
}
