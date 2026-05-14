package errcmd

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/enum/osmixtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
