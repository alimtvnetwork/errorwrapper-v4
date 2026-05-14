package errjson

import (
	"sync"

	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
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
