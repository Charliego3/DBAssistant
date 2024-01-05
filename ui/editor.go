package ui

import (
	"DataHarbor/enums"
	"DataHarbor/utility"
	"fmt"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
	"os"
)

type Editor struct {
	scroll appkit.ScrollableTextView
}

func NewEditor() *Editor {
	bs, _ := os.ReadFile("./ui/editor.go")

	layoutManager := appkit.NewLayoutManager()
	container := appkit.NewTextContainer()
	container.SetWidthTracksTextView(false)
	layoutManager.AddTextContainer(container)
	font := appkit.Font_FontWithNameSize("Menlo", 15)

	layoutDelegate := new(appkit.LayoutManagerDelegate)
	layoutDelegate.SetLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(func(layoutManager appkit.LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange foundation.RangePointer) map[foundation.AttributedStringKey]objc.Object {
		if !toScreen || effectiveCharRange == nil {
			return attrs
		}

		selectedRange := layoutManager.FirstTextView().SelectedRange()
		if selectedRange == (foundation.Range{}) {
			return attrs
		}

		attrs[enums.NSBackgroundColorAttributeName] = objc.ObjectFrom(appkit.Color_MagentaColor().Ptr())
		return attrs
	})
	layoutManager.SetDelegate(layoutDelegate)

	e := new(Editor)
	e.scroll = appkit.NewScrollableTextView()

	textView := e.scroll.ContentTextView()

	storage := textView.TextStorage()
	storage.AddLayoutManager(layoutManager)
	storage.SetAttributedString(foundation.NewAttributedStringWithString(string(bs)))
	storage.SetFont(font)

	textView.ReplaceTextContainer(container)
	textView.SetFont(font)
	textView.SetEditable(true)
	textView.SetSelectable(true)
	textView.SetHorizontallyResizable(true)
	textView.SetAutoresizingMask(appkit.ViewWidthSizable | appkit.ViewHeightSizable)
	textView.SetDelegate(e.getTextDelegate())
	textView.SetDelegateObject(new(appkit.TextDelegateObject))
	textView.SetMaxSize(utility.SizeOf(utility.Infinity, utility.Infinity))
	textView.SetBackgroundColor(appkit.Color_ClearColor())

	selectedStyle := appkit.NewMutableParagraphStyle()
	selectedStyle.SetLineHeightMultiple(1.3)
	textView.SetSelectedTextAttributes(map[foundation.AttributedStringKey]objc.IObject{
		enums.NSBackgroundColorAttributeName: appkit.Color_OrangeColor(),
		enums.NSParagraphStyleAttributeName:  selectedStyle,
	})
	utility.AddAppearanceObserver("setTextEditorTextColor", func() {
		textView.SetTextColor(utility.ColorWithAppearance(
			appkit.Color_BlackColor(),
			appkit.Color_WhiteColor(),
		))
	})

	e.scroll.SetHasHorizontalScroller(true)
	e.scroll.SetHasVerticalScroller(true)
	e.scroll.SetAutoresizingMask(appkit.ViewWidthSizable | appkit.ViewHeightSizable)

	line := LineNumberViewClass.New()
	line.InitWithScrollViewOrientation(e.scroll, appkit.VerticalRuler)
	line.SetClientView(textView)
	e.scroll.SetHasVerticalRuler(true)
	e.scroll.SetVerticalRulerView(line)
	return e
}

func (e *Editor) getTextDelegate() appkit.PTextViewDelegate {
	delegate := new(appkit.TextViewDelegate)
	delegate.SetTextDidChange(func(notification foundation.Notification) {
		fmt.Println("text changed")
	})
	delegate.SetTextDidBeginEditing(func(notification foundation.Notification) {
		fmt.Println("SetTextDidBeginEditing")
	})
	return delegate
}

type TextViewDelegateObj struct {
	appkit.TextViewDelegateObject
}

func (t TextViewDelegateObj) TextDidChange(n foundation.Notification) {
	fmt.Println("text did changed")
}

//func (e *Editor) getStorageDelegate() appkit.PTextStorageDelegate {
//	delegate := new(appkit.TextStorageDelegate)
//}
