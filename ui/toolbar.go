package ui

import (
	"github.com/progrium/macdriver/macos/appkit"
)

var toolbarIdentifiers = []appkit.ToolbarItemIdentifier{
	appkit.ToolbarToggleSidebarItemIdentifier,
	appkit.ToolbarSidebarTrackingSeparatorItemIdentifier,
	appkit.ToolbarShowColorsItemIdentifier,
}

func addToolbar(w appkit.Window) appkit.IToolbar {
	toolbar := appkit.NewToolbar()
	toolbar.SetDisplayMode(appkit.ToolbarDisplayModeIconOnly)
	toolbar.SetShowsBaselineSeparator(true)
	toolbar.SetDelegate(getToolbarDelegate())
	toolbar.SetAllowsExtensionItems(true)
	w.SetToolbar(toolbar)
	for _, item := range toolbar.Items() {
		if item.ItemIdentifier() == appkit.ToolbarToggleSidebarItemIdentifier {
			item.View().SetFocusRingType(appkit.FocusRingTypeNone)
			item.SetNavigational(false)
			item.SetImage(item.Image().ImageWithSymbolConfiguration(
				appkit.ImageSymbolConfiguration_ConfigurationWithScale(appkit.ImageSymbolScaleLarge),
			))
			break
		}
	}
	return toolbar
}

func toolbarItemIdentifiers(appkit.Toolbar) []appkit.ToolbarItemIdentifier {
	return toolbarIdentifiers
}

func getToolbarDelegate() *appkit.ToolbarDelegate {
	toolbarDelegate := &appkit.ToolbarDelegate{}
	toolbarDelegate.SetToolbarAllowedItemIdentifiers(toolbarItemIdentifiers)
	toolbarDelegate.SetToolbarDefaultItemIdentifiers(toolbarItemIdentifiers)
	toolbarDelegate.SetToolbarItemForItemIdentifierWillBeInsertedIntoToolbar(func(
		_ appkit.Toolbar,
		identifier appkit.ToolbarItemIdentifier,
		_ bool,
	) appkit.ToolbarItem {
		return appkit.ToolbarItem{}
	})
	return toolbarDelegate
}
