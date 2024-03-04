package enums

import (
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

const (
	AppearanceChangedNotification foundation.NotificationName    = "AppleInterfaceThemeChangedNotification"
	BackgroundColorAttributeName  foundation.AttributedStringKey = "NSBackgroundColor"
	ParagraphStyleAttributeName   foundation.AttributedStringKey = "NSParagraphStyle"

	ToolbarAddConnButtonIdentifier appkit.ToolbarItemIdentifier = "AddConnection"
	ToolbarToggleSidebarIdentifier appkit.ToolbarItemIdentifier = "ToolbarToggleSidebar"
)
