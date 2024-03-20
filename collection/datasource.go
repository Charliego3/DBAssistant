package main

import (
	"github.com/charliego3/assistant/utility"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func usingDiffableDatasource(w appkit.Window) appkit.IView {
	w.SetTitle("CollectionView DiffableDatasource")
	collection := appkit.NewCollectionViewWithFrame(utility.RectOf(utility.SizeOf(500, 500)))
	flowLayout := appkit.NewCollectionViewFlowLayout()
	flowLayout.SetItemSize(utility.SizeOf(100, 100))
	flowLayout.SetMinimumInteritemSpacing(20)
	flowLayout.SetMinimumLineSpacing(10)
	collection.SetCollectionViewLayout(flowLayout)
	collection.SetTranslatesAutoresizingMaskIntoConstraints(false)
	collection.RegisterClassForItemWithIdentifier(appkit.CollectionViewItemClass, collectionIdentifier)

	datasource := appkit.NewCollectionViewDiffableDataSourceWithCollectionViewItemProvider(
		collection,
		func(collection appkit.CollectionView, indexPath foundation.IndexPath, identifier objc.Object) appkit.CollectionViewItem {
			item := collection.MakeItemWithIdentifierForIndexPath(
				collectionIdentifier, indexPath,
			)
			id := foundation.StringFrom(identifier.Ptr())
			if item.TextField().IsNil() {
				item.SetTextField(appkit.TextField_LabelWithString(id.String() + "--"))
			} else {
				item.TextField().SetStringValue(id.String() + "++")
			}
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

func usingDatasource(w appkit.Window) appkit.IView {
	collection := appkit.NewCollectionViewWithFrame(utility.RectOf(utility.SizeOf(500, 500)))
	flowLayout := appkit.NewCollectionViewFlowLayout()
	flowLayout.SetItemSize(utility.SizeOf(100, 100))
	flowLayout.SetMinimumInteritemSpacing(20)
	flowLayout.SetMinimumLineSpacing(10)
	collection.SetCollectionViewLayout(flowLayout)
	collection.RegisterClassForItemWithIdentifier(appkit.CollectionViewItemClass, collectionIdentifier)
	collection.SetDataSource(new(CollectionViewDatasource))
	return collection
}

type CollectionViewDatasource struct {
	appkit.CollectionViewDataSourceObject
}

func (CollectionViewDatasource) CollectionViewItemForRepresentedObjectAtIndexPath(
	collection appkit.CollectionView,
	indexPath foundation.IndexPath,
) appkit.CollectionViewItem {
	item := collection.MakeItemWithIdentifierForIndexPath(
		collectionIdentifier, indexPath,
	)
	value := "from datasource"
	if item.TextField().IsNil() {
		item.SetTextField(appkit.TextField_LabelWithString(value + "--"))
	} else {
		item.TextField().SetStringValue(value + "++")
	}
	return item
}

func (CollectionViewDatasource) HasCollectionViewItemForRepresentedObjectAtIndexPath() bool {
	return true
}

func (CollectionViewDatasource) CollectionViewNumberOfItemsInSection(collectionView appkit.CollectionView, section int) int {
	return 1
}

func (CollectionViewDatasource) HasCollectionViewNumberOfItemsInSection() bool {
	return true
}

func (CollectionViewDatasource) NumberOfSectionsInCollectionView(collectionView appkit.CollectionView) int {
	return 10
}

func (CollectionViewDatasource) HasNumberOfSectionsInCollectionView() bool {
	return true
}
