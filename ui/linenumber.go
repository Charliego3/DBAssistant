package ui

import (
	"fmt"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var (
	LineNumberViewClass objc.UserClass[LineNumberView]
)

func init() {
	LineNumberViewClass = objc.NewClass[LineNumberView](
		objc.Sel("drawHashMarksAndLabelsInRect:"),
	)
	objc.RegisterClass(LineNumberViewClass)
}

type LineNumberView struct {
	appkit.RulerView `objc:"NSRulerView"`
}

func (view LineNumberView) DrawHashMarksAndLabelsInRect(rect foundation.Rect) {
	view.RulerView.DrawHashMarksAndLabelsInRect(rect)

	fmt.Printf("LineNumberView: %p", view.ScrollView())
}
