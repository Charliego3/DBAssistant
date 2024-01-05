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
		Window: appkit.NewWindowWithSizeAndStyle(
			homeWindowFrame.Width, homeWindowFrame.Height,
			appkit.WindowStyleMaskTitled|
				appkit.WindowStyleMaskClosable|
				appkit.WindowStyleMaskMiniaturizable|
				appkit.WindowStyleMaskResizable|
				appkit.WindowStyleMaskFullSizeContentView,
		),
	}
	objc.Retain(w)

	types := enums.DataSourceTypeMySQL
	view := getSourceView(types)
	delegate := new(appkit.SplitViewDelegate)
	delegate.SetSplitViewConstrainMinCoordinateOfSubviewAt(func(splitView appkit.SplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
		fmt.Println(splitView, proposedMinimumPosition, dividerIndex)
		return 1
	})
	delegate.SetSplitViewResizeSubviewsWithOldSize(func(splitView appkit.SplitView, oldSize foundation.Size) {
		fmt.Println(splitView, oldSize)
	})
	delegate.SetSplitViewWillResizeSubviews(func(notification foundation.Notification) {
		fmt.Println("SetSplitViewWillResizeSubviews")
	})
	splitView := appkit.NewSplitView()
	splitView.SetDelegate(delegate)
	//splitView.SetDelegateObject(getSplitViewDelegate())
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
	//w.SetContentSize(utility.SizeOf(800, 600))
	w.MakeKeyAndOrderFront(nil)
}

type SplitViewDelegate struct {
	appkit.SplitViewDelegateObject `objc:"NSSplitViewDelegate"`
}

func (s SplitViewDelegate) SplitViewResizeSubviewsWithOldSize(splitView appkit.SplitView, oldSize foundation.Size) {
	fmt.Println(splitView, oldSize)
}

func (s SplitViewDelegate) HasSplitViewResizeSubviewsWithOldSize() bool {
	return true
}

func getSplitViewDelegate() SplitViewDelegate {
	class := objc.NewClass[SplitViewDelegate](
		objc.Sel("splitView:resizeSubviewsWithOldSize:"),
	)
	objc.RegisterClass(class)
	return class.New()
}
