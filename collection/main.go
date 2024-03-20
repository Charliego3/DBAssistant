package main

import (
	"runtime"

	"github.com/charliego3/assistant/utility"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

const collectionIdentifier = appkit.UserInterfaceItemIdentifier("collectionItemIdentifier")

func main() {
	viewer := tabviewActions

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	app := appkit.Application_SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	delegate := new(appkit.ApplicationDelegate)
	delegate.SetApplicationDidFinishLaunching(launched(viewer))
	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(delegate)
	app.Run()
}

func launched(fn func(w appkit.Window) appkit.IView) func(_ foundation.Notification) {
	return func(_ foundation.Notification) {
		w := appkit.NewWindowWithSizeAndStyle(
			700, 500,
			appkit.WindowStyleMaskTitled|
				appkit.WindowStyleMaskClosable|
				appkit.WindowStyleMaskMiniaturizable|
				appkit.WindowStyleMaskResizable,
		)
		objc.Retain(&w)

		w.Center()
		w.SetTitlebarAppearsTransparent(false)
		w.SetContentView(fn(w))
		w.SetContentSize(utility.SizeOf(800, 600))
		w.MakeKeyAndOrderFront(nil)
	}
}
