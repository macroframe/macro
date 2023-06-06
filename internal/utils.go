package internal

import "github.com/macroframe/macro/pkg/mouse"

// Number is a type that can be used for math operations.
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Sum returns the sum of all numbers.
func Sum[T Number](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

func MouseButtonToRobotgoString(button mouse.Button) string {
	switch button {
	case mouse.Left:
		return "left"
	case mouse.Right:
		return "right"
	case mouse.Middle:
		return "center"
	}
	return ""
}
