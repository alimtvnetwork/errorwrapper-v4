package errorwrapper

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

func JsonResultToWrapper(
	jsonResult *corejson.Result,
) (*Wrapper, error) {
	if jsonResult.IsEmpty() {
		return nil, nil
	}

	emptyErr := NewPtr(errtype.NoError)
	deserializedErr := jsonResult.Deserialize(emptyErr)

	return emptyErr, deserializedErr
}
