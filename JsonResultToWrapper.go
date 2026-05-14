package errorwrapper

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
