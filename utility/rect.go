package utility

import (
	"github.com/progrium/macdriver/macos/foundation"
	"math"
)

const Infinity = math.MaxFloat64

func RectOf(size foundation.Size) foundation.Rect {
	return foundation.Rect{Size: size}
}

func SizeOf(width, height float64) foundation.Size {
	return foundation.Size{Width: width, Height: height}
}
