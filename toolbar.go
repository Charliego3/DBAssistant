package main

import (
	"github.com/charliego3/assistant/enums"
	"github.com/charliego3/assistant/utility"

	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

type Toolbar struct {
	appkit.Toolbar
	w appkit.Window

	splitViewController appkit.SplitViewController
}

func createToolbar(w appkit.Window, controller appkit.SplitViewController) *Toolbar {
	toolbar := new(Toolbar)
	toolbar.w = w
	toolbar.splitViewController = controller
	toolbar.Toolbar = appkit.NewToolbar()
	toolbar.SetDisplayMode(appkit.ToolbarDisplayModeIconOnly)
	toolbar.SetShowsBaselineSeparator(true)
	toolbar.SetDelegate(toolbar.getToolbarDelegate())
	toolbar.SetAllowsExtensionItems(true)
	return toolbar
}

func (t Toolbar) identifiers(appkit.Toolbar) []appkit.ToolbarItemIdentifier {
	return []appkit.ToolbarItemIdentifier{
		enums.ToolbarToggleSidebarIdentifier,
		appkit.ToolbarFlexibleSpaceItemIdentifier,
		enums.ToolbarAddConnButtonIdentifier,
		appkit.ToolbarSidebarTrackingSeparatorItemIdentifier,
		appkit.ToolbarFlexibleSpaceItemIdentifier,
		enums.ToolbarConnectionIdentifier,
		appkit.ToolbarFlexibleSpaceItemIdentifier,
	}
}

func (t Toolbar) createItem(identifier appkit.ToolbarItemIdentifier, symbol string, handler action.Handler) appkit.ToolbarItem {
	cfg := appkit.ImageSymbolConfiguration_ConfigurationWithScale(appkit.ImageSymbolScaleLarge)
	item := appkit.NewToolbarItemWithItemIdentifier(identifier)
	button := appkit.NewButton()
	button.SetImage(utility.SymbolImage(symbol, cfg))
	button.SetButtonType(appkit.ButtonTypeMomentaryPushIn)
	button.SetBezelStyle(appkit.BezelStyleTexturedRounded)
	button.SetFocusRingType(appkit.FocusRingTypeNone)
	action.Set(button, handler)
	item.SetView(button)
	return item
}

func (t Toolbar) getToolbarDelegate() *appkit.ToolbarDelegate {
	delegate := new(appkit.ToolbarDelegate)
	delegate.SetToolbarAllowedItemIdentifiers(t.identifiers)
	delegate.SetToolbarDefaultItemIdentifiers(t.identifiers)
	delegate.SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(func(
		_ appkit.Toolbar,
		identifier appkit.ToolbarItemIdentifier,
		_ bool,
	) appkit.ToolbarItem {
		switch identifier {
		case enums.ToolbarToggleSidebarIdentifier:
			return t.createItem(identifier, "sidebar.leading", func(sender objc.Object) {
				t.splitViewController.ToggleSidebar(nil)
			})
		case enums.ToolbarAddConnButtonIdentifier:
			return t.createItem(identifier, "plus", func(_ objc.Object) {
				OpenNewPanelSheet(t.w)
			})
		case enums.ToolbarConnectionIdentifier:
			view := appkit.NewBox()
			view.SetBoxType(appkit.BoxCustom)
			view.SetBorderWidth(0)
			view.SetContentViewMargins(utility.SizeOf(0, 0))
			view.SetCornerRadius(5)
			view.SetFrameSize(utility.SizeOf(200, 25))
			view.SetAutoresizingMask(appkit.ViewWidthSizable)
			utility.AddAppearanceObserver(func() {
				view.SetFillColor(utility.ColorWithAppearance(
					utility.ColorHex("#EDECEC"),
					utility.ColorHex("#ffffff").ColorWithAlphaComponent(0.05),
				))
			})
			item := appkit.NewToolbarItemWithItemIdentifier(identifier)
			item.SetView(view)
			item.SetNavigational(true)
			return item
		}
		return appkit.ToolbarItem{}
	})
	return delegate
}

func OpenNewPanelSheet(w appkit.IWindow) {
	panel := appkit.NewPanelWithContentRectStyleMaskBackingDefer(
		utility.RectOf(utility.SizeOf(600, 500)),
		appkit.WindowStyleMaskFullSizeContentView|appkit.ResizableWindowMask,
		appkit.BackingStoreBuffered,
		false,
	)

	panel.SetMinSize(panel.Frame().Size)
	creator := NewCreator(w, panel)
	panel.SetContentView(creator)
	w.BeginSheetCompletionHandler(panel, creator.Handle)
}
