package refs

import "github.com/alimtvnetwork/errorwrapper-v4/ref"

type CollectionDataModel struct {
	Refs []ref.Value
}

func NewDataModel(collection *Collection) *CollectionDataModel {
	if collection == nil || collection.IsEmpty() {
		return &CollectionDataModel{}
	}

	return &CollectionDataModel{
		Refs: collection.refs,
	}
}
