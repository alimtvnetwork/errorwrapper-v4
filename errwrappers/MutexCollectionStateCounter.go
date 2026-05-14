package errwrappers

type MutexCollectionStateCounter struct {
	collection *MutexCollection
	Start      int
}

func NewMutexStateCount(collection *MutexCollection) MutexCollectionStateCounter {
	return MutexCollectionStateCounter{Start: collection.Length(), collection: collection}
}

func (it *MutexCollectionStateCounter) StartStateTracking(start int) int {
	it.Start = start

	return start
}

func (it *MutexCollectionStateCounter) HasChanges(currentCount int) bool {
	return it.Start != currentCount
}

func (it *MutexCollectionStateCounter) IsSameState(currentCount int) bool {
	return it.Start == currentCount
}

func (it *MutexCollectionStateCounter) HasChangesCollection() bool {
	return it.Start != it.collection.Length()
}

func (it *MutexCollectionStateCounter) IsSameStateCollection() bool {
	return it.Start == it.collection.Length()
}

func (it *MutexCollectionStateCounter) IsSuccess() bool {
	return it.Start == it.collection.Length()
}

func (it *MutexCollectionStateCounter) IsFailed() bool {
	return it.Start != it.collection.Length()
}

func (it *MutexCollectionStateCounter) IsValid() bool {
	return it.Start == it.collection.Length()
}
