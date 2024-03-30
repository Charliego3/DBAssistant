package main

import (
	"github.com/charliego3/assistant/utility"
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ContentViewController struct {
	appkit.IViewController
	view     appkit.IView
	backward appkit.Button
	forward  appkit.Button
	tabView  appkit.TabView
}

func NewContentViewController() *ContentViewController {
	controller := new(ContentViewController)
	controller.IViewController = appkit.NewViewController()
	controller.view = appkit.NewView()
	controller.IViewController.SetView(controller.view)
	controller.Init()
	return controller
}

func (c *ContentViewController) Init() {
	header := c.AddHeader()
	separator := utility.SeparatorLine(utility.SeparatorOption{Super: c.view, Height: 1})
	layout.AliginAnchors(separator.TopAnchor(), header.BottomAnchor())
	layout.AliginLeading(separator, c.view)
	layout.AliginTrailing(separator, c.view)

	c.AddTabsView(separator)
}

func (c *ContentViewController) AddTabsView(top appkit.Box) {
	headers := appkit.NewView()
	headers.SetTranslatesAutoresizingMaskIntoConstraints(false)
	headers.SetWantsLayer(true)
	c.view.AddSubview(headers)
	layout.AliginAnchors(headers.TopAnchor(), top.BottomAnchor())
	layout.AliginLeading(headers, c.view)
	layout.AliginTrailing(headers, c.view)
	layout.SetHeight(headers, 27)

	recentMenu := appkit.NewMenu()
	recentMenu.AddItemWithTitleActionKeyEquivalent("A", objc.Sel("toggleSidebar:"), "")
	recent := utility.SymbolButton("square.grid.2x2", headers)
	action.Set(recent, func(sender objc.Object) {
		origin := foundation.Point{X: -8, Y: 25}
		recentMenu.PopUpMenuPositioningItemAtLocationInView(nil, origin, recent)
	})
	recent.SetMenu(recentMenu)
	layout.PinAnchorTo(recent.LeadingAnchor(), headers.LeadingAnchor(), 7)
	layout.AliginCenterY(recent, headers)

	separator := utility.SeparatorLine(utility.SeparatorOption{
		Super: headers,
		Width: 1,
		Color: appkit.Color_SeparatorColor(),
	})
	layout.PinAnchorTo(separator.LeadingAnchor(), recent.TrailingAnchor(), 7)
	layout.PinAnchorTo(separator.TopAnchor(), headers.TopAnchor(), 8)
	layout.PinAnchorTo(separator.BottomAnchor(), headers.BottomAnchor(), -8)
	layout.AliginCenterY(recent, headers)

	c.tabView = appkit.NewTabView()
	c.tabView.SetTabViewType(appkit.NoTabsNoBorder)
	c.tabView.SetTabViewItems([]appkit.ITabViewItem{
		appkit.TabViewItem_TabViewItemWithViewController(utility.Controller(appkit.NewLabel("item1"))),
		appkit.TabViewItem_TabViewItemWithViewController(utility.Controller(appkit.NewLabel("item2"))),
		appkit.TabViewItem_TabViewItemWithViewController(utility.Controller(appkit.NewLabel("item3"))),
	})

	c.backward = utility.SymbolButton("chevron.backward", headers)
	c.backward.SetEnabled(false)
	action.Set(c.backward, func(sender objc.Object) {
		c.tabView.SelectPreviousTabViewItem(sender)
		firstState := c.tabView.TabViewItems()[0].TabState()
		c.forward.SetEnabled(true)
		c.backward.SetEnabled(firstState != appkit.SelectedTab)
	})
	layout.PinAnchorTo(c.backward.LeadingAnchor(), separator.TrailingAnchor(), 7)
	layout.AliginCenterY(c.backward, headers)

	c.forward = utility.SymbolButton("chevron.right", headers)
	c.forward.SetEnabled(len(c.tabView.TabViewItems()) > 1)
	action.Set(c.forward, func(sender objc.Object) {
		c.tabView.SelectNextTabViewItem(sender)
		lastState := c.tabView.TabViewItems()[c.tabView.NumberOfTabViewItems()-1].TabState()
		c.backward.SetEnabled(true)
		c.forward.SetEnabled(lastState != appkit.SelectedTab)
	})
	layout.PinAnchorTo(c.forward.LeadingAnchor(), c.backward.TrailingAnchor(), 7)
	layout.AliginCenterY(c.forward, headers)

	tabs := appkit.NewStackView()
	tabs.SetDistribution(appkit.StackViewDistributionFill)
	tabs.SetOrientation(appkit.UserInterfaceLayoutOrientationHorizontal)
	tabs.SetAlignment(appkit.LayoutAttributeCenterY)
	tabs.SetSpacing(8)
	//tabs.SetTranslatesAutoresizingMaskIntoConstraints(false)
	//c.view.AddSubview(tabs)
	//layout.AliginAnchors(tabs.TopAnchor(), headers.TopAnchor())
	//layout.AliginAnchors(tabs.BottomAnchor(), headers.BottomAnchor())
	//layout.PinAnchorTo(tabs.LeadingAnchor(), forward.TrailingAnchor(), 7)
	//layout.AliginTrailing(tabs, c.view)

	testView := appkit.NewView()
	testView.SetWantsLayer(true)
	testView.Layer().SetBackgroundColor(appkit.Color_OrangeColor().ColorWithAlphaComponent(0.2).CGColor())
	l0 := appkit.NewLabel("test View 0")
	l0.SetTranslatesAutoresizingMaskIntoConstraints(false)

	testView.AddSubview(l0)
	layout.PinAnchorTo(l0.LeadingAnchor(), testView.LeadingAnchor(), 10)
	layout.PinAnchorTo(l0.TrailingAnchor(), testView.TrailingAnchor(), -10)
	layout.AliginCenterY(l0, testView)
	//tabs.AddArrangedSubview(testView)
	tabs.AddViewInGravity(testView, appkit.StackViewGravityLeading)
	testView1 := appkit.NewView()
	testView1.SetWantsLayer(true)
	testView1.Layer().SetBackgroundColor(appkit.Color_RedColor().ColorWithAlphaComponent(0.2).CGColor())
	l1 := appkit.NewLabel("test View 1")
	l1.SetTranslatesAutoresizingMaskIntoConstraints(false)
	testView1.AddSubview(l1)
	layout.PinAnchorTo(l1.LeadingAnchor(), testView1.LeadingAnchor(), 10)
	layout.PinAnchorTo(l1.TrailingAnchor(), testView1.TrailingAnchor(), -10)
	layout.AliginCenterY(l1, testView1)
	//tabs.AddArrangedSubview(testView1)
	tabs.AddViewInGravity(testView1, appkit.StackViewGravityLeading)
	testView1 = appkit.NewView()
	testView1.SetWantsLayer(true)
	testView1.Layer().SetBackgroundColor(appkit.Color_YellowColor().ColorWithAlphaComponent(0.2).CGColor())
	l1 = appkit.NewLabel("test View 2")
	l1.SetTranslatesAutoresizingMaskIntoConstraints(false)
	testView1.AddSubview(l1)
	layout.PinAnchorTo(l1.LeadingAnchor(), testView1.LeadingAnchor(), 10)
	layout.PinAnchorTo(l1.TrailingAnchor(), testView1.TrailingAnchor(), -10)
	layout.AliginCenterY(l1, testView1)
	//tabs.AddArrangedSubview(testView1)
	tabs.AddViewInGravity(testView1, appkit.StackViewGravityLeading)

	scroll := appkit.NewScrollView()
	scroll.SetHasHorizontalScroller(true)
	scroll.SetHasVerticalScroller(false)
	scroll.SetBorderType(appkit.NoBorder)
	scroll.SetTranslatesAutoresizingMaskIntoConstraints(false)
	scroll.SetBackgroundColor(utility.ColorHex("#292A2F"))
	//scroll.SetDrawsBackground(true)
	//scroll.SetDocumentView(tabs)
	tabs.SetTranslatesAutoresizingMaskIntoConstraints(false)
	//c.view.AddSubview(tabs)
	//layout.AliginAnchors(tabs.TopAnchor(), headers.TopAnchor())
	//layout.AliginAnchors(tabs.BottomAnchor(), headers.BottomAnchor())
	//layout.PinAnchorTo(tabs.LeadingAnchor(), c.forward.TrailingAnchor(), 7)
	//layout.AliginTrailing(tabs, c.view)

	separator = utility.SeparatorLine(utility.SeparatorOption{
		Super:  c.view,
		Height: 1,
		Light:  utility.ColorHex("#E3E3E3"),
		Dark:   utility.ColorHex("#464445"),
	})
	layout.AliginAnchors(separator.TopAnchor(), headers.BottomAnchor())
	layout.AliginLeading(separator, c.view)
	layout.AliginTrailing(separator, c.view)

	placeholder := appkit.NewView()
	placeholder.SetTranslatesAutoresizingMaskIntoConstraints(false)
	placeholder.SetWantsLayer(true)
	c.view.AddSubview(placeholder)
	layout.AliginLeading(placeholder, c.view)
	layout.AliginTrailing(placeholder, c.view)
	layout.AliginAnchors(placeholder.TopAnchor(), separator.BottomAnchor())
	layout.AliginBottom(placeholder, c.view)
	c.tabView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	placeholder.AddSubview(c.tabView)
	layout.AliginTop(c.tabView, placeholder)
	layout.AliginBottom(c.tabView, placeholder)
	layout.AliginTrailing(c.tabView, placeholder)
	layout.AliginLeading(c.tabView, placeholder)

	//tips := appkit.NewLabel("No open datasource or editor")
	//tips.SetTranslatesAutoresizingMaskIntoConstraints(false)
	//tips.SetFont(appkit.Font_SystemFontOfSize(20))
	//tips.SetTextColor(appkit.Color_SecondaryLabelColor())
	//placeholder.AddSubview(tips)
	//layout.AliginCenterY(tips, placeholder)
	//layout.AliginCenterX(tips, placeholder)
	//utility.AddAppearanceObserver(func() {
	//	placeholder.Layer().SetBackgroundColor(utility.ColorWithAppearance(
	//		appkit.Color_WhiteColor(),
	//		utility.ColorHex("#313030"),
	//	).CGColor())
	//	headers.Layer().SetBackgroundColor(utility.ColorWithAppearance(
	//		appkit.Color_WhiteColor(),
	//		utility.ColorHex("#313030"),
	//	).CGColor())
	//})
}

func (c *ContentViewController) AddHeader() appkit.IView {
	header := appkit.NewView()
	header.SetTranslatesAutoresizingMaskIntoConstraints(false)
	header.SetWantsLayer(true)
	utility.AddAppearanceObserver(func() {
		header.Layer().SetBackgroundColor(utility.ColorWithAppearance(
			utility.ColorHex("#FAF9F9"),
			utility.ColorHex("#403F3F"),
		).CGColor())
	})
	c.view.AddSubview(header)
	layout.AliginTop(header, c.view)
	layout.AliginTrailing(header, c.view)
	layout.AliginLeading(header, c.view)
	layout.SetHeight(header, 38)
	return header
}
