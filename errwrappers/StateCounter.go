package errwrappers

import (
	"github.com/alimtvnetwork/core-v9/coreinterface"
)

type StateCounter struct {
	lengthGetter coreinterface.LengthGetter
	Start        int
}

func NewStateCount(
	collection *Collection,
) StateCounter {
	return StateCounter{
		Start:        collection.Length(),
		lengthGetter: collection,
	}
}

func NewStateCountUsingLengthGetter(
	lengthGetter coreinterface.LengthGetter,
) StateCounter {
	return StateCounter{
		Start:        lengthGetter.Length(),
		lengthGetter: lengthGetter,
	}
}

func (it StateCounter) IsSameStateUsingCount(currentCount int) bool {
	return it.Start == currentCount
}

func (it StateCounter) IsSameState() bool {
	return it.Start == it.lengthGetter.Length()
}

func (it StateCounter) HasChanges() bool {
	return it.Start != it.lengthGetter.Length()
}

func (it StateCounter) StartStateTracking(start int) int {
	it.Start = start

	return start
}

func (it StateCounter) HasChangesCollection() bool {
	return it.Start != it.lengthGetter.Length()
}

func (it StateCounter) IsSameStateCollection() bool {
	return it.Start == it.lengthGetter.Length()
}

func (it StateCounter) IsSuccess() bool {
	return it.Start == it.lengthGetter.Length()
}

func (it StateCounter) IsFailed() bool {
	return it.Start != it.lengthGetter.Length()
}

func (it StateCounter) IsValid() bool {
	return it.Start == it.lengthGetter.Length()
}

func (it StateCounter) AsCountStateTrackerBinder() coreinterface.CountStateTrackerBinder {
	return &it
}
