package main

import (
	"fmt"

	"github.com/charliego3/assistant/db"
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
	db.InitializeDB()
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
	delegate.SetOutlineViewViewForTableColumnItem(func(_ appkit.OutlineView, tableColumn appkit.TableColumn, item objc.Object) appkit.View {
		fmt.Println("into view for table column item")
		conn := foundation.StringFrom(item.Ptr())
		text := appkit.NewTextField()
		text.SetBordered(false)
		text.SetBezelStyle(appkit.TextFieldSquareBezel)
		text.SetEditable(false)
		text.SetDrawsBackground(false)
		text.SetTranslatesAutoresizingMaskIntoConstraints(false)
		text.SetStringValue(conn.CapitalizedString())

		rowView := appkit.NewTableRowView()
		rowView.AddSubview(text)

		text.LeadingAnchor().ConstraintEqualToAnchor(rowView.LeadingAnchor()).SetActive(true)
		text.TrailingAnchor().ConstraintEqualToAnchor(rowView.TrailingAnchor()).SetActive(true)
		text.CenterYAnchor().ConstraintEqualToAnchor(rowView.CenterYAnchor()).SetActive(true)
		return rowView.View
	})
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
		fmt.Println("into child of item")
		defer func() {
			fmt.Println("out child of item")
		}()
		if item.IsNil() {
			return foundation.String_StringWithString(db.Conns()[index].Name).Object
		}

		conn := (*db.Connection)(item.Ptr())
		// return objc.ObjectFrom(unsafe.Pointer(conn.Children[index]))
		// return items[index]
		return foundation.String_StringWithString(conn.Children[index].Name).Object
	})
	datasource.SetOutlineViewIsItemExpandable(func(outlineView appkit.OutlineView, item objc.Object) bool {
		fmt.Println("into view is expandable")
		defer func() {
			fmt.Println("out view is expandable")
		}()
		if item.IsNil() {
			return false
		}

		// conn := (*db.Connection)(item.Ptr())
		// return conn.Type == enums.DataSourceTypeFolder
		return false
	})
	// this is test text
	datasource.SetOutlineViewNumberOfChildrenOfItem(func(_ appkit.OutlineView, item objc.Object) int {
		fmt.Println("into number of children of item")
		defer func() {
			fmt.Println("out number of children of item")
		}()
		// if item.IsNil() {
		// 	return len(db.Conns())
		// }

		// conn := (*db.Connection)(item.Ptr())
		// return len(conn.Children)
		return len(db.Conns())
	})
	po1 := objc.WrapAsProtocol("NSOutlineViewDataSource", appkit.POutlineViewDataSource(datasource))
	objc.SetAssociatedObject(s.outline, objc.AssociationKey("setDataSource"), po1, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s.outline, objc.Sel("setDataSource:"), po1)
}

func (s *Sidebar) SetSidebarMaxWidth() {
	if !s.max.IsNil() {
		s.max.SetActive(false)
	}
	s.max = s.view.WidthAnchor().ConstraintLessThanOrEqualToConstant(s.w.Frame().Size.Width / 2)
	s.max.SetActive(true)
}
