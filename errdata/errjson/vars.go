package errjson

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/coredynamic"
)

var (
	writerLock = sync.Mutex{}
	New        = &newCreator{
		Result:            &newResultCreator{},
		ResultsCollection: &newResultsCollectionCreator{},
	}
	Empty      = &emptyCreator{}
	resultType = coredynamic.TypeName(Result{})
)
