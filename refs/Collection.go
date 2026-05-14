package refs

import (
	"encoding/json"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

type Collection struct {
	refs []ref.Value
}

func (it *Collection) MapStringString() map[string]string {
	return it.DynamicMap().ConvMapStringString()
}

func (it *Collection) CloneNewDefiner() errcoreinf.ReferenceCollectionDefiner {
	return it.ClonePtr()
}

func (it *Collection) ReferencesList() []errcoreinf.Referencer {
	return it.ReferencerCollection()
}

func (it *Collection) ReferencerCollection() []errcoreinf.Referencer {
	if it.IsEmpty() {
		return []errcoreinf.Referencer{}
	}

	newSlice := make([]errcoreinf.Referencer, it.Length())

	for i, value := range it.refs {
		newSlice[i] = value.AsReferencer()
	}

	return newSlice
}

func (it *Collection) Count() int {
	return it.Length()
}

func (it *Collection) AddVarVal(varName string, val interface{}) errcoreinf.ReferenceCollectionDefiner {
	it.Add(varName, val)

	return it
}

func (it *Collection) AddReferencer(ref errcoreinf.Referencer) errcoreinf.ReferenceCollectionDefiner {
	if ref == nil || it.IsEmpty() {
		return it
	}

	it.AddReferencer(ref)

	return it
}

func (it *Collection) AddReferences(
	references ...errcoreinf.Referencer,
) errcoreinf.ReferenceCollectionDefiner {
	if len(references) == 0 {
		return it
	}

	for _, reference := range references {
		if reference == nil || it.IsEmpty() {
			continue
		}

		it.AddReferencer(reference)
	}

	return it
}

func (it *Collection) MapStringAny() map[string]interface{} {
	if it.IsEmpty() {
		return map[string]interface{}{}
	}

	newMap := make(map[string]interface{}, it.Length())

	for _, value := range it.refs {
		newMap[value.KeyName()] = value.Value
	}

	return newMap
}

func (it *Collection) Serialize() ([]byte, error) {
	return it.Json().Raw()
}

func (it *Collection) SerializeMust() (jsonBytes []byte) {
	return it.JsonPtr().RawMust()
}

func (it *Collection) Compile() string {
	return it.String()
}

func (it *Collection) ReflectSetTo(
	toPointer interface{},
) error {
	return coredynamic.ReflectSetFromTo(it, toPointer)
}

func (it *Collection) Add(
	name string,
	val interface{},
) *Collection {
	it.refs = append(
		it.refs,
		ref.New(name, val))

	return it
}

func (it *Collection) Adds(
	refs ...ref.Value,
) *Collection {
	if len(refs) == 0 {
		return it
	}

	for i := range refs {
		it.refs = append(it.refs, refs[i])
	}

	return it
}

func (it *Collection) AddsIf(
	isAdd bool,
	refs ...ref.Value,
) *Collection {
	if !isAdd || len(refs) == 0 {
		return it
	}

	for i := range refs {
		it.refs = append(it.refs, refs[i])
	}

	return it
}

func (it *Collection) AddCollection(
	collection *Collection,
) *Collection {
	if collection.IsEmpty() {
		return it
	}

	return it.Adds(collection.refs...)
}

func (it *Collection) AddCollections(
	collections ...*Collection,
) *Collection {
	if len(collections) == 0 {
		return it
	}

	for _, collection := range collections {
		if collection.IsEmpty() {
			continue
		}

		it.Adds(collection.refs...)
	}

	return it
}

func (it *Collection) AddCollectionCloned(
	collections ...*Collection,
) *Collection {
	if len(collections) == 0 {
		return it
	}

	for _, collection := range collections {
		if collection.IsEmpty() {
			continue
		}

		it.AddsByCloningItems(collection.refs...)
	}

	return it
}

func (it *Collection) ConcatNew(
	isSingleItemsClone bool,
	collections ...*Collection,
) *Collection {
	return NewUsingCollection(
		isSingleItemsClone,
		it,
		collections...)
}

func (it *Collection) concatNewCloneItems(
	length int,
	collections ...*Collection,
) *Collection {
	if len(collections) == 0 {
		return it.ClonePtr()
	}

	clonedNew := New(length)

	if it != nil {
		clonedNew.AddsByCloningItems(it.refs...)
	}

	return clonedNew.
		AddCollectionCloned(collections...)
}

func (it *Collection) AddsPtrByCloningItems(
	refs ...*ref.Value,
) *Collection {
	if len(refs) == 0 {
		return it
	}

	for i := range refs {
		it.refs = append(
			it.refs,
			refs[i].Clone())
	}

	return it
}

func (it *Collection) AddsByCloningItems(
	refs ...ref.Value,
) *Collection {
	if len(refs) == 0 {
		return it
	}

	for i := range refs {
		it.refs = append(
			it.refs,
			refs[i].Clone())
	}

	return it
}

func (it *Collection) AddsPtr(
	refs ...*ref.Value,
) *Collection {
	if len(refs) == 0 {
		return it
	}

	for _, refItem := range refs {
		if refItem == nil {
			continue
		}

		it.refs = append(
			it.refs,
			*refItem)
	}

	return it
}

func (it *Collection) AddMap(
	collectionMap map[string]interface{},
) *Collection {
	if collectionMap == nil || len(collectionMap) == 0 {
		return it
	}

	for k, v := range collectionMap {
		it.refs = append(
			it.refs,
			ref.New(k, v))
	}

	return it
}

func (it *Collection) DynamicMap() enumimpl.DynamicMap {
	if it.IsEmpty() {
		return enumimpl.DynamicMap{}
	}

	return it.MapStringAny()
}

func (it *Collection) IsEqual(
	another *Collection,
) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it == another {
		return true
	}

	if it.IsEmpty() && another.IsEmpty() {
		return true
	}

	leftLen := it.Length()
	rightLen := another.Length()

	if leftLen != rightLen {
		return false
	}

	if leftLen == rightLen && leftLen == 0 {
		return true
	}

	anotherRefs := another.refs

	if &anotherRefs == &it.refs {
		return true
	}

	for i, currentRef := range it.refs {
		if !anotherRefs[i].IsEqual(currentRef) {
			return false
		}
	}

	return true
}

func (it *Collection) IsNull() bool {
	return it == nil
}

func (it *Collection) IsEmpty() bool {
	return it == nil ||
		len(it.refs) == 0
}

func (it *Collection) Collection() []ref.Value {
	return it.refs
}

func (it *Collection) List() []ref.Value {
	return it.refs
}

func (it *Collection) Items() []ref.Value {
	return it.refs
}

func (it *Collection) Length() int {
	if it == nil || it.refs == nil {
		return 0
	}

	return len(it.refs)
}

func (it *Collection) Strings() []string {
	stringsCollection := make([]string, it.Length())

	for i, v := range it.refs {
		stringsCollection[i] = v.FullString()
	}

	return stringsCollection
}

func (it Collection) String() string {
	return strings.Join(
		it.Strings(),
		constants.CommaSpace)
}

func (it *Collection) Dispose() {
	if it == nil {
		return
	}

	it.refs = nil
}

func (it *Collection) ClonePtr() *Collection {
	if it == nil {
		return nil
	}

	if it.refs != nil {
		refs := make([]ref.Value, it.Length())

		for i, value := range it.refs {
			refs[i] = value.Clone()
		}

		return &Collection{
			refs: refs,
		}
	}

	return EmptyPtr()
}

func (it Collection) Clone() Collection {
	cloned := it.ClonePtr()

	if cloned != nil {
		return *cloned
	}

	return Empty()
}

func (it *Collection) JsonModel() []ref.Value {
	if it == nil {
		return nil
	}

	return it.refs
}

func (it *Collection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModelAny())
}

func (it *Collection) UnmarshalJSON(rawJsonBytes []byte) error {
	var referencesDataModel []ref.Value

	err := corejson.Deserialize.UsingBytes(
		rawJsonBytes,
		&referencesDataModel)

	if err == nil {
		it.refs = referencesDataModel
	}

	return err
}

func (it Collection) Json() corejson.Result {
	return corejson.New(it)
}

func (it Collection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *Collection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Collection, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return it, defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(
		jsonResult.Bytes,
		&it)

	if err != nil {
		return it, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Collection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Collection {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Collection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Collection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Collection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Collection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *Collection) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *Collection) HasAnyItem() bool {
	return it.Length() > 0
}
