package refs

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"

	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

func New(capacity int) *Collection {
	refs := make([]ref.Value, 0, capacity)

	return &Collection{
		refs: refs,
	}
}

func New2() *Collection {
	refs := make(
		[]ref.Value,
		0,
		constants.Capacity2)

	return &Collection{
		refs: refs,
	}
}

func New4() *Collection {
	refs := make(
		[]ref.Value,
		0,
		constants.Capacity4)

	return &Collection{
		refs: refs,
	}
}

func NewUsingCollection(
	isCloneSingleItems bool,
	first *Collection,
	items ...*Collection,
) *Collection {
	length := AllIndividualItemsCount(
		first,
		items...) + constants.Capacity4
	hasItems := len(items) > 0
	hasFirst := first != nil
	collection := New(length)

	if hasFirst && isCloneSingleItems {
		collection.AddCollectionCloned(first)
	} else if hasFirst && !isCloneSingleItems {
		collection.AddCollections(first)
	}

	if isCloneSingleItems && hasItems {
		collection.AddCollectionCloned(items...)
	} else if hasItems {
		collection.AddCollections(items...)
	}

	return collection
}

func NewUsingBasicErrWrap(
	basicErrWrap errcoreinf.BasicErrWrapper,
) *Collection {
	if basicErrWrap == nil || basicErrWrap.IsEmpty() || basicErrWrap.HasReferences() {
		return EmptyPtr()
	}

	referencesList := basicErrWrap.ReferencesList()

	return NewUsingReferencers(referencesList...)
}

func NewUsingReferencers(
	references ...errcoreinf.Referencer,
) *Collection {
	length := len(references)
	sliceOfReferences := make([]ref.Value, length)

	if length == 0 {
		return &Collection{
			refs: sliceOfReferences,
		}
	}

	for i, referencer := range references {
		sliceOfReferences[i] = ref.New(
			referencer.VarName(),
			referencer.ValueDynamic())
	}

	return &Collection{
		refs: sliceOfReferences,
	}
}

func NewUsingRefs(
	isClone bool,
	items ...ref.Value,
) *Collection {
	length := len(items)
	sliceOfReferences := make([]ref.Value, length)

	if length == 0 {
		return &Collection{
			refs: sliceOfReferences,
		}
	}

	if !isClone {
		return &Collection{
			refs: items,
		}
	}

	for i, item := range items {
		sliceOfReferences[i] = item.Clone()
	}

	return &Collection{
		refs: sliceOfReferences,
	}
}

func NewUsingRefsOrNil(
	items ...ref.Value,
) *Collection {
	if len(items) == 0 {
		return nil
	}

	return NewUsingRefs(false, items...)
}

func NewUsingValues(
	items ...ref.Value,
) *Collection {
	if len(items) == 0 {
		return nil
	}

	return NewUsingRefs(false, items...)
}

func NewUsingMap(
	itemsMap map[string]interface{},
) *Collection {
	if len(itemsMap) == 0 {
		return EmptyPtr()
	}

	collection := New(len(itemsMap) + 1)
	for variable, value := range itemsMap {
		collection.Add(variable, value)
	}

	return collection
}

func NewClone(
	items ...ref.Value,
) *Collection {
	if len(items) == 0 {
		return nil
	}

	return NewUsingRefs(true, items...)
}

func NewUsingMany(
	manyCollections ...[]ref.Value,
) *Collection {
	length := LengthOfEachItems(manyCollections)
	refs := make([]ref.Value, 0, length)

	if length > 0 {
		for _, collection := range manyCollections {
			if len(collection) == 0 {
				continue
			}

			refs = append(refs, collection...)
		}
	}

	return &Collection{
		refs: refs,
	}
}

func NewExistingCollectionPlusAddition(
	existingReferences *Collection,
	newReferences ...ref.Value,
) *Collection {
	existingLength := constants.Zero

	if existingReferences != nil && existingReferences.Length() > 0 {
		existingLength = existingReferences.Length()
	}

	newReferencesLength := len(newReferences)

	refs := make([]ref.Value, 0, existingLength+newReferencesLength)

	if existingLength > 0 {
		refs = append(refs, existingReferences.refs...)
	}

	if newReferencesLength > 0 {
		refs = append(refs, newReferences...)
	}

	return &Collection{
		refs: refs,
	}
}

func NewExistingPlusAddition(
	existingReferences []ref.Value,
	newReferences ...ref.Value,
) *Collection {
	return &Collection{
		refs: MergeReferences(
			existingReferences,
			newReferences...),
	}
}

func MergeReferences(
	existingReferences []ref.Value,
	newReferences ...ref.Value,
) []ref.Value {
	existingLength := len(existingReferences)
	newReferencesLength := len(newReferences)

	mergedReferences := make(
		[]ref.Value,
		0,
		existingLength+newReferencesLength)

	if existingLength > 0 {
		mergedReferences = append(mergedReferences, existingReferences...)
	}

	if newReferencesLength > 0 {
		mergedReferences = append(mergedReferences, newReferences...)
	}

	return mergedReferences
}

func PrependReferences(
	isClone bool,
	appendingItems []ref.Value,
	prependingItems ...ref.Value,
) []ref.Value {
	if !isClone && len(appendingItems) == 0 {
		return prependingItems
	}

	if !isClone && len(prependingItems) == 0 {
		return appendingItems
	}

	return MergeReferences(prependingItems, appendingItems...)
}

func NewWithItem(capacity int, varName string, val interface{}) *Collection {
	collection := New(capacity)

	return collection.Add(varName, val)
}

func NewDirectItem(varName string, val interface{}) *Collection {
	collection := New(constants.N1)

	return collection.Add(varName, val)
}

func Empty() Collection {
	return *New(0)
}

func EmptyPtr() *Collection {
	return New(0)
}

func NewFromDataModelPtr(model *CollectionDataModel) *Collection {
	if len(model.Refs) == 0 {
		return EmptyPtr()
	}

	return &Collection{
		refs: model.Refs,
	}
}
