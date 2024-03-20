package main

import (
	_ "embed"
	"fmt"

	"github.com/charliego3/assistant/utility"

	_ "github.com/charliego3/assistant/images"

	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var (
	// homeWindowFrame the main window initialize frame size
	homeWindowFrame = utility.SizeOf(700, 500)

	// contentSplitViewClass custom split view
	contentSplitViewClass objc.UserClass[ContentSplitView]

	mouseViewClass objc.UserClass[MouseView]
)

type ContentSplitView struct {
	appkit.SplitView `objc:"NSSplitView"`
}

func (ContentSplitView) DividerThickness() float64 {
	return 30
}

func (ContentSplitView) DividerColor() appkit.Color {
	return appkit.Color_ClearColor()
}

type MouseView struct {
	appkit.View `objc:"NSView"`
}

var (
	j      int
	inView bool
)

func (v MouseView) MouseEntered(event appkit.IEvent) {
	j++
	appkit.Cursor_ArrowCursor().Set()
	inView = true
	fmt.Printf("entered:%d=%t, ", j, inView)
}

func (v MouseView) MouseExited(event appkit.IEvent) {
	j++
	appkit.Cursor_CurrentCursor().Set()
	inView = false
	fmt.Printf("exited:%d=%t, ", j, inView)
}

// func (v MouseView) MouseDown(event appkit.IEvent) {
// 	j++
// 	fmt.Printf("down:%d=%t, ", j, inView)
// 	if inView {
// 		return
// 	}
// }

func init() {
	contentSplitViewClass = objc.NewClass[ContentSplitView](
		objc.Sel("dividerThickness"),
		objc.Sel("dividerColor"),
	)
	objc.RegisterClass(contentSplitViewClass)

	mouseViewClass = objc.NewClass[MouseView](
		objc.Sel("mouseEntered:"),
		objc.Sel("mouseExited:"),
		// objc.Sel("mouseDown:"),
	)
	objc.RegisterClass(mouseViewClass)
}

type HomeWindow struct {
	appkit.Window
	sidebar    appkit.IView
	sidebarMax appkit.LayoutConstraint
	app        appkit.Application
	editor     *Editor
	console    appkit.View

	// splitViewController content split view controller
	splitViewController appkit.SplitViewController
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

	controller := appkit.NewSplitViewController()
	controller.AddSplitViewItem(appkit.SplitViewItem_SidebarWithViewController(w.Sidebar()))
	controller.AddSplitViewItem(appkit.SplitViewItem_SplitViewItemWithViewController(nil))

	delegate := new(appkit.WindowDelegate)
	// delegate.SetWindowDidEndLiveResize(func(notification foundation.Notification) {
	// 	w.SetSidebarMaxWidth()
	// })

	toolbar := createToolbar(w.Window, controller)
	// utility.AddAppearanceObserver(func() {
	// 	w.SetBackgroundColor(utility.ColorWithAppearance(
	// 		appkit.Color_WhiteColor(),
	// 		utility.ColorHex("#292a2f"),
	// 	))
	// })
	w.Center()
	w.SetDelegate(delegate)
	w.SetToolbar(toolbar)
	// w.SetContentMinSize(homeWindowFrame)
	w.SetToolbarStyle(appkit.WindowToolbarStyleUnifiedCompact)
	w.SetTitlebarSeparatorStyle(appkit.TitlebarSeparatorStyleNone)
	w.SetTitlebarAppearsTransparent(false)
	w.SetContentViewController(controller)
	w.SetContentSize(utility.SizeOf(800, 600))
	w.MakeKeyAndOrderFront(nil)
}

func (w *HomeWindow) SetSidebarMaxWidth() {
	if !w.sidebarMax.IsNil() {
		w.sidebarMax.SetActive(false)
	}
	w.sidebarMax = w.sidebar.WidthAnchor().ConstraintLessThanOrEqualToConstant(w.Frame().Size.Width / 2)
	w.sidebarMax.SetActive(true)
}

func (w *HomeWindow) Sidebar() appkit.ViewController {
	w.sidebar = appkit.NewView()
	w.sidebar.SetTranslatesAutoresizingMaskIntoConstraints(false)
	w.SetSidebarMaxWidth()
	layout.SetMinWidth(w.sidebar, 200)
	return utility.Controller(w.sidebar)
}

func (w *HomeWindow) Content() appkit.IViewController {
	controller := appkit.NewSplitViewController()
	w.splitViewController = controller.Init()
	splitView := w.NewContentSplitView()
	w.editor = NewEditor()
	w.console = appkit.NewViewWithFrame(utility.RectOf(utility.SizeOf(400, 200)))
	w.console.SetWantsLayer(true)
	utility.AddAppearanceObserver(func() {
		w.console.Layer().SetBackgroundColor(utility.ColorWithAppearance(
			appkit.Color_WhiteColor(),
			utility.ColorHex("#2E2E2E"),
		).CGColor())
	})

	topline := utility.SeparatorLine(utility.SeparatorOption{
		Super:  w.console,
		Height: 1,
	})

	layout.AliginTop(topline, w.console)
	layout.AliginLeading(topline, w.console)
	layout.AliginTrailing(topline, w.console)

	button := appkit.NewButtonWithTitle("title string")
	action.Set(button, func(sender objc.Object) {
		// w.splitViewController.SplitViewItems()[1].SetCollapsed(true)
		w.splitViewController.ToggleSidebar(sender)
	})
	w.console.AddSubview(button)
	w.splitViewController.SetSplitView(splitView)
	w.splitViewController.SetSplitViewItems([]appkit.ISplitViewItem{
		appkit.SplitViewItem_SplitViewItemWithViewController(utility.Controller(w.editor)),
		appkit.SplitViewItem_SidebarWithViewController(utility.Controller(w.console)),
	})
	w.AddActionDivider()
	return w.splitViewController
}

func (w *HomeWindow) NewContentSplitView() appkit.SplitView {
	size := w.Frame().Size
	frame := utility.RectOf(utility.SizeOf(size.Width-w.sidebar.Frame().Size.Width, size.Height))
	splitView := contentSplitViewClass.New().InitWithFrame(frame)
	splitView.SetDividerStyle(appkit.SplitViewDividerStyleThin)
	splitView.SetVertical(false)
	return splitView
}

func (w *HomeWindow) AddActionDivider() {
	view := mouseViewClass.New().InitWithFrame(utility.RectOf(utility.SizeOf(0, 30)))
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	view.SetWantsLayer(true)
	utility.AddAppearanceObserver(func() {
		view.Layer().SetBackgroundColor(utility.ColorWithAppearance(
			appkit.Color_WhiteColor(),
			utility.ColorHex("#2C2B2C"),
		).CGColor())
	})
	super := w.splitViewController.View()
	super.AddSubview(view)

	execute := utility.SymbolButton("play.fill", view)
	fill := utility.SymbolButton("play.slash", view)
	toggleConsole := utility.SymbolButton("square.bottomthird.inset.filled", view)
	action.Set(toggleConsole, func(sender objc.Object) {
		w.splitViewController.ToggleSidebar(sender)
	})

	offset := w.splitViewController.SplitView().DividerThickness()
	thickness := w.splitViewController.SplitView().DividerThickness() - 10.5
	layout.AliginLeading(view, super)
	layout.AliginTrailing(view, super)
	layout.PinAnchorTo(view.BottomAnchor(), w.editor.BottomAnchor(), offset)
	layout.PinAnchorTo(execute.LeadingAnchor(), view.LeadingAnchor(), 11)
	layout.PinAnchorTo(execute.BottomAnchor(), w.editor.BottomAnchor(), thickness)
	layout.PinAnchorTo(fill.LeadingAnchor(), execute.TrailingAnchor(), 11)
	layout.PinAnchorTo(fill.BottomAnchor(), w.editor.BottomAnchor(), thickness)
	layout.PinAnchorTo(toggleConsole.TrailingAnchor(), view.TrailingAnchor(), -11)
	layout.PinAnchorTo(toggleConsole.BottomAnchor(), w.editor.BottomAnchor(), thickness)

	view.AddTrackingArea(appkit.NewTrackingAreaWithRectOptionsOwnerUserInfo(
		utility.RectOf(utility.SizeOf(60, 30)),
		appkit.TrackingMouseEnteredAndExited|appkit.TrackingActiveAlways,
		view,
		foundation.DictionaryFrom(nil),
	))
}
