package ui

import (
	"DataHarbor/utility"
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

type MySQLView struct{}

func (m *MySQLView) Sidebar() appkit.ViewController {
	view := appkit.NewView()
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	layout.SetMinWidth(view, 200)
	controller := appkit.NewViewController()
	controller.SetView(view)
	return controller
}

func (m *MySQLView) Content() appkit.ViewController {
	editor := NewEditor()
	view := appkit.NewView()
	view.AddSubview(editor.scroll)
	view.SetWantsLayer(true)
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	layout.PinEdgesToSuperView(editor.scroll, foundation.EdgeInsets{})
	layout.PinEdgesToSuperView(view, foundation.EdgeInsets{Top: 10})
	utility.AddAppearanceObserver("setTextEditorBackgroundColor", func() {
		view.Layer().SetBackgroundColor(utility.ColorWithAppearance(
			appkit.Color_WhiteColor(),
			utility.ColorWithRGBA(41, 42, 47, 1),
		).CGColor())
	})

	controller := appkit.NewViewController()
	controller.SetView(view)
	return controller
}
