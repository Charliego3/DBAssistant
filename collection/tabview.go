package main

import (
	"fmt"

	"github.com/charliego3/assistant/utility"
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

func tabviewActions(w appkit.Window) appkit.IView {
	w.SetTitle("TabViewController actions")
	view := appkit.NewView()

	tabView := appkit.NewTabView()
	tabView.SetTabViewBorderType(appkit.TabViewBorderTypeBezel)
	tabView.SetTabViewType(appkit.TopTabsBezelBorder)
	tabView.SetTabViewItems([]appkit.ITabViewItem{
		appkit.TabViewItem_TabViewItemWithViewController(utility.Controller(appkit.NewLabel("item1"))),
		appkit.TabViewItem_TabViewItemWithViewController(utility.Controller(appkit.NewLabel("item2"))),
		appkit.TabViewItem_TabViewItemWithViewController(utility.Controller(appkit.NewLabel("item3"))),
	})

	for idx, item := range tabView.TabViewItems() {
		item.SetLabel(fmt.Sprintf("Item %d", idx+1))
	}

	//controller := appkit.NewTabViewController()
	//controller.SetTabView(tabView)

	backward := appkit.NewButtonWithTitle("backward")
	backward.SetFocusRingType(appkit.FocusRingTypeNone)
	action.Set(backward, func(sender objc.Object) {
		fmt.Println(len(tabView.TabViewItems()))
		tabView.SelectPreviousTabViewItem(sender)
		for i, item := range tabView.TabViewItems() {
			fmt.Println(i, item.TabState())
		}
	})

	forward := appkit.NewButtonWithTitle("forward")
	action.Set(forward, func(sender objc.Object) {
		fmt.Println(len(tabView.TabViewItems()))
		tabView.SelectNextTabViewItem(sender)
	})

	tabView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	backward.SetTranslatesAutoresizingMaskIntoConstraints(false)
	forward.SetTranslatesAutoresizingMaskIntoConstraints(false)
	view.AddSubview(backward)
	view.AddSubview(forward)
	view.AddSubview(tabView)

	layout.PinAnchorTo(backward.LeadingAnchor(), view.LeadingAnchor(), 10)
	layout.PinAnchorTo(backward.TopAnchor(), view.TopAnchor(), 10)
	layout.PinAnchorTo(forward.LeadingAnchor(), backward.TrailingAnchor(), 20)
	layout.PinAnchorTo(forward.TopAnchor(), view.TopAnchor(), 10)
	layout.PinAnchorTo(tabView.TopAnchor(), backward.BottomAnchor(), 20)
	layout.PinAnchorTo(tabView.LeadingAnchor(), view.LeadingAnchor(), 10)
	layout.PinAnchorTo(tabView.TrailingAnchor(), view.TrailingAnchor(), -10)
	layout.PinAnchorTo(tabView.BottomAnchor(), view.BottomAnchor(), -10)
	return view
}
