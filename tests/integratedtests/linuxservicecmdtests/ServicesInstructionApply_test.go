package linuxservicecmdtests

import (
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/core-v9/coretests"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/tests/testwrappers/linuxservicecmdtestwrappers"
)

func Test_ServicesInstructionApply(t *testing.T) {
	coretests.SkipOnWindows(t)
	if runtime.GOOS == "darwin" {
		t.Skip("requires linux systemctl/service")
	}

	for _, testCase := range linuxservicecmdtestwrappers.ServicesTestCases {
		errCollection := errwrappers.NewCap4()
		testCase.Apply(errCollection)

		Convey(testCase.Header, t, func() {
			So(errCollection.StringWithoutHeader(), ShouldEqual, "")
		})
	}
}
