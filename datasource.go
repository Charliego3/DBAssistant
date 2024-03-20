package main

import (
	_ "embed"
	"fmt"

	"github.com/charliego3/assistant/utility"

	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

const collectionIdentifier = "collectionDatasourceIdentifier"

type Creator struct {
	appkit.View
	appkit.PCollectionViewDataSource
	w appkit.IWindow
	p appkit.Panel
}

func NewCreator(w appkit.IWindow, panel appkit.Panel) Creator {
	return Creator{
		View: appkit.NewView(),
		w:    w,
		p:    panel,
	}.Init()
}

func (c Creator) Init() Creator {
	title := appkit.NewLabel("Choose a data source")
	title.SetTranslatesAutoresizingMaskIntoConstraints(false)
	c.AddSubview(title)
	layout.PinAnchorTo(title.TopAnchor(), c.TopAnchor(), 10)
	layout.PinAnchorTo(title.LeadingAnchor(), c.LeadingAnchor(), 20)

	box := appkit.NewBox()
	box.SetTranslatesAutoresizingMaskIntoConstraints(false)
	box.SetBorderColor(appkit.Color_SeparatorColor())
	box.SetBorderWidth(1)
	box.SetBoxType(appkit.BoxCustom)
	grid := c.getDataSourceView()
	box.AddSubview(grid)
	c.AddSubview(box)
	layout.AliginTop(grid, box)
	layout.AliginTrailing(grid, box)
	layout.AliginBottom(grid, box)
	layout.AliginLeading(grid, box)
	layout.PinAnchorTo(box.TopAnchor(), title.BottomAnchor(), 10)
	layout.PinAnchorTo(box.LeadingAnchor(), c.LeadingAnchor(), 20)
	layout.PinAnchorTo(box.TrailingAnchor(), c.TrailingAnchor(), -20)

	cancel := appkit.NewButtonWithTitle("Cancel")
	cancel.SetTranslatesAutoresizingMaskIntoConstraints(false)
	action.Set(cancel, func(sender objc.Object) { c.w.EndSheet(c.p) })
	c.AddSubview(cancel)
	layout.SetWidth(cancel, 100)
	layout.PinAnchorTo(cancel.TopAnchor(), box.BottomAnchor(), 20)
	layout.PinAnchorTo(cancel.LeadingAnchor(), c.LeadingAnchor(), 20)
	layout.PinAnchorTo(cancel.BottomAnchor(), c.BottomAnchor(), -20)

	next := appkit.NewButtonWithTitle("Next")
	next.SetTranslatesAutoresizingMaskIntoConstraints(false)
	next.SetBezelColor(appkit.Color_BlueColor())
	next.SetEnabled(false)
	action.Set(next, func(objc.Object) {
		title.SetStringValue("Choose options for your selected data source")
	})
	c.AddSubview(next)
	layout.SetWidth(next, 100)
	layout.PinAnchorTo(next.TrailingAnchor(), c.TrailingAnchor(), -20)
	layout.PinAnchorTo(next.BottomAnchor(), c.BottomAnchor(), -20)
	return c
}

type CollectionItem struct {
	appkit.CollectionViewItem `objc:"NSCollectionViewItem"`
}

func (CollectionItem) Identifier() appkit.UserInterfaceItemIdentifier {
	return collectionIdentifier
}

func (CollectionItem) LoadView() {
	fmt.Println(status)
}

func (CollectionItem) AwakeFromNib() {
	fmt.Println(status)
}

var itemClass = objc.NewClass[CollectionItem](
	objc.Sel("identifier"),
	objc.Sel("loadView"),
	objc.Sel("awakeFromNib"),
)

// var objects = make([]objc.IObject, 1)
var objects = []objc.IObject{}
var status bool

func (c Creator) getDataSourceView() appkit.IView {
	objc.RegisterClass(itemClass)
	collection := appkit.NewCollectionViewWithFrame(utility.RectOf(utility.SizeOf(500, 500)))
	flowLayout := appkit.NewCollectionViewFlowLayout()
	flowLayout.SetItemSize(utility.SizeOf(100, 100))
	flowLayout.SetMinimumInteritemSpacing(20)
	flowLayout.SetMinimumLineSpacing(10)
	collection.SetCollectionViewLayout(flowLayout)
	collection.SetDataSource(&c)
	collection.SetTranslatesAutoresizingMaskIntoConstraints(false)
	collection.RegisterClassForItemWithIdentifier(itemClass, collectionIdentifier)

	// var objects = []objc.IObject{}
	// status := nib.InstantiateWithOwnerTopLevelObjects(nil, objects)
	// if status == false {
	// 	panic("failed load collection item from nib")
	// }
	// collection.RegisterNibForItemWithIdentifier(nib, collectionIdentifier)

	datasource := appkit.NewCollectionViewDiffableDataSourceWithCollectionViewItemProvider(
		collection,
		func(collection appkit.CollectionView, indexPath foundation.IndexPath, identifier objc.Object) appkit.CollectionViewItem {
			item := collection.MakeItemWithIdentifierForIndexPath(
				collectionIdentifier, indexPath,
			)
			from := foundation.StringFrom(identifier.Ptr())
			fmt.Println("new collection view item:", from.String())
			if item.TextField().IsNil() {
				item.SetTextField(appkit.TextField_LabelWithString("stringValue ----"))
			} else {
				item.TextField().SetStringValue("value string")
			}
			fmt.Println("new collection done")
			return item
		},
	)
	snapshot := appkit.NewDiffableDataSourceSnapshot()
	snapshot.AppendSectionsWithIdentifiers([]objc.IObject{
		foundation.String_StringWithString(string(collectionIdentifier)),
	})
	snapshot.AppendItemsWithIdentifiers([]objc.IObject{
		foundation.String_StringWithString(string(collectionIdentifier)),
	})
	objc.Call[appkit.CollectionViewDiffableDataSource](datasource, objc.Sel("applySnapshot:animatingDifferences:"), snapshot, false)
	return collection
}

func (Creator) CollectionViewItemForRepresentedObjectAtIndexPath(
	collectionView appkit.CollectionView,
	indexPath foundation.IndexPath,
) appkit.CollectionViewItem {
	fmt.Println("into collection view method", indexPath.Section(), indexPath.Item(), indexPath.Length())
	item := collectionView.MakeItemWithIdentifierForIndexPath(
		collectionIdentifier, indexPath,
	)
	// item := appkit.NewCollectionViewItem()
	fmt.Println("new collection view item")
	if item.TextField().IsNil() {
		item.SetTextField(appkit.TextField_LabelWithString("stringValue ----"))
	} else {
		item.TextField().SetStringValue("value string")
	}
	fmt.Println("new collection done")
	return item
}

func (Creator) HasCollectionViewItemForRepresentedObjectAtIndexPath() bool {
	return true
}

func (Creator) CollectionViewNumberOfItemsInSection(collectionView appkit.CollectionView, section int) int {
	return 1
}

func (Creator) HasCollectionViewNumberOfItemsInSection() bool {
	return true
}

func (Creator) NumberOfSectionsInCollectionView(collectionView appkit.CollectionView) int {
	return 10
}

func (Creator) HasNumberOfSectionsInCollectionView() bool {
	return true
}

func (Creator) Handle(code appkit.ModalResponse) {
	fmt.Println(code, ".......")
}
