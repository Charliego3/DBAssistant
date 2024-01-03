package main

import (
	"DataHarbor/ui"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
	"runtime"
)

var Version string

func main() {
	if Version == "" {
		Version = "1.0.0"
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	app := appkit.Application_SharedApplication()
	delegate := &appkit.ApplicationDelegate{}
	delegate.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		//ui.ActiveWelcomeWindow(app)
		ui.ActiveHomeWindow(app)
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	delegate.SetApplicationWillFinishLaunching(func(foundation.Notification) {
		setMainMenu(app)
	})
	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(delegate)
	app.Run()
}

func setMainMenu(app appkit.Application) {
	menu := appkit.NewMenuWithTitle("main")
	app.SetMainMenu(menu)

	mainMenuItem := appkit.NewMenuItemWithSelector("", "", objc.Selector{})
	mainMenuMenu := appkit.NewMenuWithTitle("App")
	mainMenuMenu.AddItem(appkit.NewMenuItemWithAction("Hide", "h", func(_ objc.Object) { app.Hide(nil) }))
	mainMenuMenu.AddItem(appkit.NewMenuItemWithAction("Quit", "q", func(_ objc.Object) { app.Terminate(nil) }))
	mainMenuItem.SetSubmenu(mainMenuMenu)
	menu.AddItem(mainMenuItem)

	testMenuItem := appkit.NewMenuItemWithSelector("", "", objc.Selector{})
	testMenu := appkit.NewMenuWithTitle("Edit")
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Select All", "a", objc.Sel("selectAll:")))
	testMenu.AddItem(appkit.MenuItem_SeparatorItem())
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Copy", "c", objc.Sel("copy:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Paste", "v", objc.Sel("paste:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Cut", "x", objc.Sel("cut:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Undo", "z", objc.Sel("undo:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Redo", "Z", objc.Sel("redo:")))
	testMenuItem.SetSubmenu(testMenu)
	menu.AddItem(testMenuItem)
}
