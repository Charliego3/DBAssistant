package ui

import (
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
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
	controller := appkit.NewViewController()
	controller.SetView(NewEditor())
	return controller
}
