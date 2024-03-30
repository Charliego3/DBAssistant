package main

import (
	"github.com/charliego3/assistant/db"
	"github.com/charliego3/assistant/images"
	"github.com/charliego3/assistant/lib"
	"github.com/charliego3/assistant/utility"
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

const sidebarIdentifier = "sidebarDatasourceIdentifier"

var items []objc.Object

func init() {
	for i := 65; i < 65+26; i++ {
		items = append(items, foundation.String_StringWithString(string(rune(i))).Object)
	}
}

type Sidebar struct {
	appkit.IViewController
	w       appkit.IWindow
	view    appkit.IView
	outline appkit.OutlineView
	max     appkit.LayoutConstraint
}

func NewSidebarController(w appkit.IWindow) *Sidebar {
	sidebar := new(Sidebar)
	sidebar.w = w
	sidebar.IViewController = appkit.NewViewController()
	sidebar.Init()
	return sidebar
}

func (s *Sidebar) Init() {
	s.outline = appkit.NewOutlineView()
	s.outline.SetColumnAutoresizingStyle(appkit.TableViewSequentialColumnAutoresizingStyle)
	s.outline.SetUsesAlternatingRowBackgroundColors(false)
	s.outline.SetStyle(appkit.TableViewStyleSourceList)
	s.outline.SetSelectionHighlightStyle(appkit.TableViewSelectionHighlightStyleSourceList)
	s.outline.SetUsesSingleLineMode(true)
	s.outline.SetAllowsColumnSelection(false)
	s.outline.SetAutoresizingMask(appkit.ViewWidthSizable)
	s.outline.SetHeaderView(nil)
	s.outline.AddTableColumn(utility.TableColumn(sidebarIdentifier, ""))
	s.setDelegate()
	s.setDatasource()
	scrollView := appkit.NewScrollView()
	clipView := appkit.ClipViewFrom(scrollView.ContentView().Ptr())
	clipView.SetDocumentView(s.outline)
	clipView.SetAutomaticallyAdjustsContentInsets(false)
	clipView.SetContentInsets(foundation.EdgeInsets{Top: 10})

	s.outline.SelectRowIndexesByExtendingSelection(foundation.NewIndexSetWithIndex(0), true)
	scrollView.SetBorderType(appkit.NoBorder)
	scrollView.SetDrawsBackground(false)
	scrollView.SetAutohidesScrollers(true)
	scrollView.SetHasVerticalScroller(true)
	scrollView.ContentView().ScrollToPoint(foundation.Point{Y: -10})

	s.view = scrollView
	s.IViewController.SetView(s.view)
	s.view.SetFrameSize(utility.SizeOf(260, 0))
	layout.SetMinWidth(s.view, 200)
	s.SetSidebarMaxWidth()
}

func (s *Sidebar) setDelegate() {
	delegate := new(appkit.OutlineViewDelegate)
	delegate.SetOutlineViewViewForTableColumnItem(s.createColumnItem)
	delegate.SetControlTextDidBeginEditing(func(obj foundation.Notification) {})
	// delegate.SetOutlineViewIsGroupItem(func(outlineView appkit.OutlineView, item objc.Object) bool {
	// 	s := foundation.StringFrom(item.Ptr())
	// 	return s.String() == "A -- 65"
	// })
	//delegate.SetOutlineViewHeightOfRowByItem(func(outlineView appkit.OutlineView, item objc.Object) float64 {
	//	s := foundation.StringFrom(item.Ptr())
	//	if s.String() == "A -- 65" {
	//		return 50
	//	}
	//	return 30
	//})
	po0 := objc.WrapAsProtocol("NSOutlineViewDelegate", appkit.POutlineViewDelegate(delegate))
	objc.SetAssociatedObject(s.outline, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s.outline, objc.Sel("setDelegate:"), po0)
}

func (s *Sidebar) setDatasource() {
	datasource := new(lib.OutlineViewDatasource)
	datasource.SetOutlineViewChildOfItem(func(outlineView appkit.OutlineView, index int, item objc.Object) objc.Object {
		if item.IsNil() {
			return foundation.Number_NumberWithInt(db.TopConnection(index)).Object
		}

		num := foundation.NumberFrom(item.Ptr())
		conn := db.Connections[db.Childrens[num.IntegerValue()][index]]
		return foundation.Number_NumberWithInt(conn.Id).Object
	})
	datasource.SetOutlineViewIsItemExpandable(func(outlineView appkit.OutlineView, item objc.Object) bool {
		num := foundation.NumberFrom(item.Ptr())
		return db.FetchChildrenLength(num.IntegerValue()) > 0
	})
	datasource.SetOutlineViewNumberOfChildrenOfItem(func(_ appkit.OutlineView, item objc.Object) int {
		if item.IsNil() {
			return db.TopConnLength()
		}

		num := foundation.NumberFrom(item.Ptr())
		return len(db.Childrens[num.IntegerValue()])
	})
	appkit.NewTableViewRowAction()
	po1 := objc.WrapAsProtocol("NSOutlineViewDataSource", appkit.POutlineViewDataSource(datasource))
	objc.SetAssociatedObject(s.outline, objc.AssociationKey("setDataSource"), po1, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s.outline, objc.Sel("setDataSource:"), po1)
}

func (s *Sidebar) createColumnItem(_ appkit.OutlineView, tableColumn appkit.TableColumn, item objc.Object) appkit.View {
	num := foundation.NumberFrom(item.Ptr())
	conn := db.Connections[num.IntegerValue()]

	image := appkit.NewImageView()
	image.SetTranslatesAutoresizingMaskIntoConstraints(false)
	// image.SetImage(utility.SymbolImage("folder.fill"))
	icon := appkit.NewImageWithData(images.RedisData)
	icon.SetSize(utility.SizeOf(16, 16))
	image.SetImage(icon)

	text := appkit.NewTextField()
	text.SetBordered(false)
	text.SetBezelStyle(appkit.TextFieldSquareBezel)
	text.SetEditable(true)
	text.SetDrawsBackground(false)
	text.SetTranslatesAutoresizingMaskIntoConstraints(false)
	text.SetStringValue(conn.Name + " = " + conn.Type.String())

	rowView := appkit.NewTableRowView()
	rowView.AddSubview(image)
	rowView.AddSubview(text)

	image.LeadingAnchor().ConstraintEqualToAnchor(rowView.LeadingAnchor()).SetActive(true)
	image.CenterYAnchor().ConstraintEqualToAnchor(rowView.CenterYAnchor()).SetActive(true)
	text.LeadingAnchor().ConstraintEqualToAnchorConstant(image.TrailingAnchor(), 3).SetActive(true)
	text.TrailingAnchor().ConstraintEqualToAnchor(rowView.TrailingAnchor()).SetActive(true)
	text.CenterYAnchor().ConstraintEqualToAnchor(rowView.CenterYAnchor()).SetActive(true)
	return rowView.View
}

func (s *Sidebar) SetSidebarMaxWidth() {
	if !s.max.IsNil() {
		s.max.SetActive(false)
	}
	s.max = s.view.WidthAnchor().ConstraintLessThanOrEqualToConstant(s.w.Frame().Size.Width / 2)
	s.max.SetActive(true)
}
