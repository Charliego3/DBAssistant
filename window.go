package main

import (
	"fmt"

	"github.com/charliego3/assistant/db"
	"github.com/charliego3/assistant/utility"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var (
	// homeWindowFrame the main window initialize frame size
	windowFrame = utility.SizeOf(700, 500)
)

func showMainWindow(_ foundation.Notification) {
	w := appkit.NewWindowWithSizeAndStyle(
		windowFrame.Width, windowFrame.Height,
		appkit.WindowStyleMaskTitled|
			appkit.WindowStyleMaskClosable|
			appkit.WindowStyleMaskMiniaturizable|
			appkit.WindowStyleMaskResizable|
			appkit.WindowStyleMaskFullSizeContentView,
	)
	objc.Retain(&w)

	db.InitializeDB()
	sidebarController := NewSidebarController(w)
	controller := appkit.NewSplitViewController()
	controller.SetSplitViewItems([]appkit.ISplitViewItem{
		appkit.SplitViewItem_SidebarWithViewController(sidebarController),
		appkit.SplitViewItem_SplitViewItemWithViewController(NewContentViewController()),
	})

	delegate := new(appkit.WindowDelegate)
	delegate.SetWindowDidEndLiveResize(func(notification foundation.Notification) {
		sidebarController.SetSidebarMaxWidth()
	})
	w.Center()
	w.SetDelegate(delegate)
	w.SetToolbar(createToolbar(w, controller))
	w.SetContentMinSize(homeWindowFrame)
	w.SetToolbarStyle(appkit.WindowToolbarStyleUnifiedCompact)
	w.SetTitlebarSeparatorStyle(appkit.TitlebarSeparatorStyleAutomatic)
	w.SetTitlebarAppearsTransparent(true)
	w.SetContentViewController(controller)
	w.SetContentSize(utility.SizeOf(800, 600))
	w.MakeKeyAndOrderFront(nil)
	//time.Sleep(time.Millisecond * 100)
	//OpenNewPanelSheet(w)

	key := "testKey"
	defaults := appkit.UserDefaultsController_SharedUserDefaultsController().Defaults()
	value := defaults.StringForKey(key)
	if value == "" {
		fmt.Println("defaults is empty")
		defaults.SetObjectForKey(foundation.String_StringWithString("testValue"), key)
	} else {
		fmt.Println("defaults value is", value)
	}
}
