package ui

import (
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
	"os"
)

type Editor struct {
	appkit.ScrollView
}

func NewEditor() *Editor {
	bs, _ := os.ReadFile("/Users/charlie/dev/go/DataHarbor/ui/editor.go")
	//textView := appkit.NewTextView()
	//textView.SetString(string(bs))
	//textView.SetEditable(true)
	//textView.SetSelectable(true)
	//textView.SetBackgroundColor(appkit.Color_TextBackgroundColor())
	//textView.SetFont(appkit.Font_SystemFontOfSize(15))
	//textView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	//
	//scroll := appkit.NewScrollView()
	//scroll.SetDocumentView(textView)
	//scroll.SetHasHorizontalScroller(true)
	//scroll.SetHasVerticalScroller(true)
	//scroll.SetAutoresizingMask(appkit.ViewHeightSizable | appkit.ViewWidthSizable)
	//
	//layout.AliginTop(textView, scroll)
	//layout.AliginLeading(textView, scroll)
	//layout.AliginBottom(textView, scroll)
	//layout.AliginTrailing(textView, scroll)

	scroll := appkit.TextView_ScrollableTextView()
	scroll.SetHasHorizontalScroller(true)
	scroll.SetHasVerticalScroller(true)
	scroll.SetAutoresizingMask(appkit.ViewHeightSizable | appkit.ViewWidthSizable)

	tv := appkit.TextViewFrom(scroll.DocumentView().Ptr())
	tv.SetAllowsUndo(true)
	tv.SetString(string(bs))
	tv.SetEditable(true)
	tv.SetSelectable(true)
	tv.SetBackgroundColor(appkit.Color_TextBackgroundColor())
	tv.SetFont(appkit.Font_FontWithNameSize("MonoLisa Nerd Font", 15))
	tv.SetHorizontallyResizable(true)
	tv.SetHorizontalContentSizeConstraintActive(false)
	tv.TextContainer().SetWidthTracksTextView(false)
	tv.SetSelectedTextAttributes(map[foundation.AttributedStringKey]objc.IObject{
		"NSBackgroundColor": appkit.Color_OrangeColor(),
	})
	return &Editor{ScrollView: scroll}
}
