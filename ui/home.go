package ui

import (
	"DataHarbor/enums"
	"DataHarbor/utility"
	"fmt"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type SourceView interface {
	Sidebar() appkit.ViewController
	Content() appkit.ViewController
}

func getSourceView(types enums.DataSourceType) SourceView {
	switch types {
	case enums.DataSourceTypeMySQL:
		return new(MySQLView)
	default:
		return nil
	}
}

var homeWindowFrame = utility.SizeOf(600, 400)

type HomeWindow struct {
	appkit.Window
	app appkit.Application
}

func ActiveHomeWindow(app appkit.Application) {
	w := &HomeWindow{
		app: app,
		Window: appkit.NewWindowWithContentRectStyleMaskBackingDefer(
			utility.RectOf(homeWindowFrame),
			appkit.WindowStyleMaskTitled|
				appkit.WindowStyleMaskClosable|
				appkit.WindowStyleMaskMiniaturizable|
				appkit.WindowStyleMaskResizable|
				appkit.WindowStyleMaskFullSizeContentView,
			appkit.BackingStoreBuffered,
			false,
		),
	}
	objc.Retain(w)

	types := enums.DataSourceTypeMySQL
	view := getSourceView(types)
	delegate := new(appkit.SplitViewDelegate)
	delegate.SetSplitViewConstrainMinCoordinateOfSubviewAt(func(splitView appkit.SplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
		fmt.Println(splitView, proposedMinimumPosition, dividerIndex)
		return 500
	})
	delegate.SetSplitViewResizeSubviewsWithOldSize(func(splitView appkit.SplitView, oldSize foundation.Size) {
		fmt.Println(splitView, oldSize)
	})
	splitView := appkit.NewSplitView()
	splitView.SetDelegate(delegate)
	splitView.SetVertical(true)
	controller := appkit.NewSplitViewController()
	controller.SetSplitView(splitView)
	controller.AddSplitViewItem(appkit.SplitViewItem_SidebarWithViewController(view.Sidebar()))
	controller.AddSplitViewItem(appkit.SplitViewItem_SplitViewItemWithViewController(view.Content()))

	addToolbar(w.Window)
	w.Center()
	w.SetContentMinSize(homeWindowFrame)
	w.SetToolbarStyle(appkit.WindowToolbarStyleUnifiedCompact)
	w.SetTitlebarSeparatorStyle(appkit.TitlebarSeparatorStyleLine)
	w.SetTitlebarAppearsTransparent(false)
	w.SetContentViewController(controller)
	w.SetContentSize(utility.SizeOf(800, 600))
	w.MakeKeyAndOrderFront(nil)
}
