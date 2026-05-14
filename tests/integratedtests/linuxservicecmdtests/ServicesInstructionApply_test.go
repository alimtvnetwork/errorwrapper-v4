package linuxservicecmdtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/tests/testwrappers/linuxservicecmdtestwrappers"
)

func Test_ServicesInstructionApply(t *testing.T) {
	coretests.SkipOnWindows(t)

	for _, testCase := range linuxservicecmdtestwrappers.ServicesTestCases {
		errCollection := errwrappers.NewCap4()
		testCase.Apply(errCollection)

		Convey(testCase.Header, t, func() {
			So(errCollection.StringWithoutHeader(), ShouldEqual, "")
		})
	}
}
