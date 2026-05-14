package errcmd

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	osmixtype "github.com/alimtvnetwork/enum-v10/osdetect"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func CurrentOsDetails() (*osmixtype.OperatingSystemDetail, *errorwrapper.Wrapper) {
	osDetails, err := osmixtype.GetCurrentOsDetail()

	if err != nil {
		return nil, errorwrapper.NewMsgDisplayErrorNoReference(
			codestack.Skip1,
			errtype.OperatingSystemRelated,
			err.Error())
	}

	return osDetails, nil
}
