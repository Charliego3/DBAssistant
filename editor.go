package main

import (
	"github.com/charliego3/assistant/utility"
	"fmt"
	"io/fs"
	"os"

	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/webkit"
	"github.com/progrium/macdriver/objc"
)

type Editor struct {
	webkit.WebView
}

func NewEditor() *Editor {
	var _fs fs.FS
	if Released {
		_fs, _ = fs.Sub(assetsFS, "assets")
	} else {
		_fs = os.DirFS("assets")
	}
	gofsHandler := &webkit.FileSystemURLSchemeHandler{FS: _fs}
	configuration := webkit.NewWebViewConfiguration()
	configuration.SetURLSchemeHandlerForURLScheme(gofsHandler, "gofs")
	wv := webkit.NewWebViewWithFrameConfiguration(
		foundation.Rect{},
		configuration,
	)
	wv.SetWantsLayer(true)
	wv.Layer().SetOpaque(false)
	wv.Layer().SetBackgroundColor(appkit.Color_ClearColor().CGColor())
	// wv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	wv.SetCanDrawConcurrently(true)
	wv.SetCanDrawSubviewsIntoLayer(true)
	layout.SetMinHeight(wv, 38)
	webkit.AddScriptMessageHandlerWithReply(wv, "greet", func(message objc.Object) (objc.Object, error) {
		param := message.Description()
		fmt.Println("greet handled", param)
		return foundation.NewStringWithString("hello: " + param).Object, nil
	})
	webkit.LoadURL(wv, "gofs:/index.html")

	var lineheight float64 = 1
	// view := appkit.NewView()
	// view.AddSubview(wv)
	// view.SetWantsLayer(true)
	// utility.AddAppearanceObserver(func() {
	// 	view.Layer().SetBackgroundColor(utility.ColorWithAppearance(
	// 		appkit.Color_WhiteColor(),
	// 		utility.ColorHex("#292a2f"),
	// 	).CGColor())
	// })

	// layout.AliginTop(wv, view)
	// layout.AliginLeading(wv, view)
	// layout.AliginTrailing(wv, view)
	// layout.PinAnchorTo(wv.BottomAnchor(), view.BottomAnchor(), -lineheight)

	topline := utility.SeparatorLine(utility.SeparatorOption{Super: wv, Height: lineheight})
	bottomline := utility.SeparatorLine(utility.SeparatorOption{Super: wv, Height: lineheight})
	layout.PinAnchorTo(topline.TopAnchor(), wv.TopAnchor(), 38)
	layout.AliginLeading(topline, wv)
	layout.AliginTrailing(topline, wv)

	layout.AliginLeading(bottomline, wv)
	layout.AliginTrailing(bottomline, wv)
	layout.AliginBottom(bottomline, wv)
	return &Editor{WebView: wv}
}
