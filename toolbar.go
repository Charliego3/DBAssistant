package main

import (
	"DataForge/enums"
	"DataForge/utility"
	"fmt"

	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

type Toolbar struct {
	appkit.Toolbar
	w appkit.Window

	splitViewController appkit.SplitViewController
}

func getToolbar(w appkit.Window, controller appkit.SplitViewController) *Toolbar {
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
		// appkit.ToolbarToggleSidebarItemIdentifier,
		enums.ToolbarToggleSidebarIdentifier,
		appkit.ToolbarFlexibleSpaceItemIdentifier,
		enums.ToolbarAddConnButtonIdentifier,
		appkit.ToolbarSidebarTrackingSeparatorItemIdentifier,
		appkit.ToolbarShowFontsItemIdentifier,
		appkit.ToolbarShowColorsItemIdentifier,
	}
}

func (t Toolbar) createItem(identifier appkit.ToolbarItemIdentifier, symbol string) appkit.ToolbarItem {
	cfg := appkit.ImageSymbolConfiguration_ConfigurationWithScale(appkit.ImageSymbolScaleLarge)
	item := appkit.NewToolbarItemWithItemIdentifier(identifier)
	item.SetBordered(true)
	item.SetImage(utility.SymbolImage(symbol, cfg))
	return item
}

func (t Toolbar) removeFocusRingType() {
	for _, item := range t.Items() {
		if item.View().IsNil() {
			continue
		}

		fmt.Println(item.ItemIdentifier())
		item.View().SetFocusRingType(appkit.FocusRingTypeNone)
		item.SetNavigational(false)
		item.SetImage(item.Image().ImageWithSymbolConfiguration(
			appkit.ImageSymbolConfiguration_ConfigurationWithScale(appkit.ImageSymbolScaleLarge),
		))
	}
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
			item := t.createItem(identifier, "sidebar.leading")
			action.Set(item, func(sender objc.Object) {
				t.splitViewController.ToggleSidebar(sender)
			})
			return item
		case enums.ToolbarAddConnButtonIdentifier:
			item := t.createItem(identifier, "plus")
			action.Set(item, func(sender objc.Object) {
				fmt.Println("clicked add button")
				t.OpenNewPanelSheet()
			})
			return item
		}
		return appkit.ToolbarItem{}
	})
	return delegate
}

func (t Toolbar) OpenNewPanelSheet() {
	panel := appkit.NewPanelWithContentRectStyleMaskBackingDefer(
		utility.RectOf(utility.SizeOf(300, 300)),
		appkit.WindowStyleMaskFullSizeContentView,
		appkit.BackingStoreBuffered,
		false,
	)

	content := appkit.NewButtonWithTitle("Close")
	action.Set(content, func(sender objc.Object) {
		t.w.EndSheet(panel)
	})
	panel.SetContentView(content)
	t.w.BeginSheetCompletionHandler(panel, func(returnCode appkit.ModalResponse) {
		fmt.Println(returnCode)
	})
}
