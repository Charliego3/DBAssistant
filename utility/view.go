package utility

import (
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

func Controller(view appkit.IView) appkit.ViewController {
	controller := appkit.NewViewController()
	controller.SetView(view)
	return controller
}

func Active(constraints ...appkit.LayoutConstraint) {
	for _, constraint := range constraints {
		constraint.SetActive(true)
	}
}

type SeparatorOption struct {
	Super  appkit.IView
	Color  appkit.Color
	Width  float64
	Height float64
}

func SeparatorLine(opt SeparatorOption) appkit.Box {
	box := appkit.NewBox()
	box.SetFrameSize(SizeOf(opt.Width, opt.Height))
	box.SetBoxType(appkit.BoxCustom)
	box.SetBorderWidth(0)
	box.SetTranslatesAutoresizingMaskIntoConstraints(false)
	if opt.Color.IsNil() {
		AddAppearanceObserver(func() {
			box.SetFillColor(ColorWithAppearance(
				ColorHex("#DADADA"),
				ColorHex("#000000"),
			))
		})
	} else {
		box.SetFillColor(opt.Color)
	}
	if !opt.Super.IsNil() {
		opt.Super.AddSubview(box)
	}
	return box
}

func SymbolButton(symbol string, config appkit.ImageSymbolConfiguration, handler func(sender objc.Object)) appkit.Button {
	button := appkit.NewButtonWithImage(SymbolImage(symbol, config))
	button.SetTranslatesAutoresizingMaskIntoConstraints(false)
	button.SetBordered(false)
	action.Set(button, handler)
	return button
}
