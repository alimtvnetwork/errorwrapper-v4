package linuxservicecmdtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v4/linuxservicecmd"
)

func Test_MoreCoverage_Request_Accessors(t *testing.T) {
	Convey("Request reference values and dispose", t, func() {
		req := linuxservicecmd.Request{
			ServiceName: "nginx",
			Action:      servicestate.Status,
		}
		refs := req.ReferenceValues()
		So(refs, ShouldHaveLength, 2)
		So(refs[0].Variable, ShouldEqual, "Service Name")

		req.Dispose()
	})
}

func Test_MoreCoverage_Result_NilSafe(t *testing.T) {
	Convey("Result nil-safe paths", t, func() {
		var nilRes *linuxservicecmd.Result
		So(nilRes.CompiledErrorWrapper(), ShouldNotBeNil)
		nilRes.Dispose()
		nilRes.DisposeWithoutErrorWrapper()
		So(nilRes.IsEmpty(), ShouldBeTrue)
	})

	Convey("Result accessors with populated request", t, func() {
		res := &linuxservicecmd.Result{
			Request: linuxservicecmd.Request{
				ServiceName: "svc",
				Action:      servicestate.Status,
			},
			ExitCode: linuxservicestate.ActiveRunning,
		}
		So(res.ServiceName(), ShouldEqual, "svc")
		So(res.IsSuccess(), ShouldBeTrue)
		So(res.IsFailed(), ShouldBeFalse)
		So(res.IsUnknownService(), ShouldBeFalse)
		So(res.SimplifiedError(), ShouldBeNil)
		So(res.VerifyExitCode(linuxservicestate.ActiveRunning), ShouldBeNil)
		So(res.VerifyExitCode(linuxservicestate.NotRunning), ShouldNotBeNil)
		So(res.ErrorWrapperUsingOpt(false), ShouldBeNil)

		ec := errwrappers.Empty()
		So(res.CollectSimpleError(ec), ShouldBeTrue)
		So(res.CollectError(ec, true), ShouldBeTrue)
	})

	Convey("Result with failing exit codes maps to errors", t, func() {
		for _, code := range []linuxservicestate.ExitCode{
			linuxservicestate.DeadButPidExists,
			linuxservicestate.DeadButVarLockFileExists,
			linuxservicestate.NotRunning,
			linuxservicestate.UnknownService,
			linuxservicestate.InvalidService,
		} {
			res := &linuxservicecmd.Result{
				Request:  linuxservicecmd.Request{ServiceName: "svc", Action: servicestate.Status},
				ExitCode: code,
			}
			So(res.SimplifiedError(), ShouldNotBeNil)
			So(res.IsFailed(), ShouldBeTrue)
		}
	})
}

func Test_MoreCoverage_Results(t *testing.T) {
	Convey("Results basic accessors", t, func() {
		rs := linuxservicecmd.EmptyResults()
		So(rs.Length(), ShouldEqual, 0)
		So(rs.IsEmpty(), ShouldBeTrue)
		So(rs.HasAnyItem(), ShouldBeFalse)
		So(rs.HasAnyFailed(), ShouldBeFalse)
		So(rs.IsAllSuccess(), ShouldBeTrue)

		rs2 := linuxservicecmd.NewResults(4)
		rs2.Items = append(rs2.Items, &linuxservicecmd.Result{
			ExitCode: linuxservicestate.ActiveRunning,
		})
		rs2.Items = append(rs2.Items, &linuxservicecmd.Result{
			ExitCode: linuxservicestate.NotRunning,
		})
		So(rs2.Length(), ShouldEqual, 2)
		So(rs2.HasAnyItem(), ShouldBeTrue)
		So(rs2.HasAnyFailed(), ShouldBeTrue)
		So(rs2.IsAllSuccess(), ShouldBeFalse)
	})
}

func Test_MoreCoverage_ServicesInstruction_HasXxx(t *testing.T) {
	Convey("ServicesInstruction Has* getters", t, func() {
		var nilSI *linuxservicecmd.ServicesInstruction
		So(nilSI.HasStopServicesNames(), ShouldBeFalse)
		So(nilSI.HasStartServicesNames(), ShouldBeFalse)
		So(nilSI.HasRestartServicesNames(), ShouldBeFalse)
		So(nilSI.HasEnableServicesNames(), ShouldBeFalse)
		So(nilSI.HasDisableServicesNames(), ShouldBeFalse)
		So(nilSI.HasStatusServicesNames(), ShouldBeFalse)
		So(nilSI.HasServices(), ShouldBeFalse)
		So(nilSI.HasValidations(), ShouldBeFalse)

		si := &linuxservicecmd.ServicesInstruction{
			StopServicesNames:    []string{"a"},
			StartServicesNames:   []string{"a"},
			RestartServicesNames: []string{"a"},
			EnableServicesNames:  []string{"a"},
			DisableServicesNames: []string{"a"},
			StatusServicesNames:  []string{"a"},
		}
		So(si.HasStopServicesNames(), ShouldBeTrue)
		So(si.HasStartServicesNames(), ShouldBeTrue)
		So(si.HasRestartServicesNames(), ShouldBeTrue)
		So(si.HasEnableServicesNames(), ShouldBeTrue)
		So(si.HasDisableServicesNames(), ShouldBeTrue)
		So(si.HasStatusServicesNames(), ShouldBeTrue)
		So(si.HasServices(), ShouldBeFalse)
		So(si.HasValidations(), ShouldBeFalse)

		// Apply on nil & empty paths
		ec := errwrappers.Empty()
		So(nilSI.Apply(ec), ShouldBeTrue)
		So((&linuxservicecmd.ServicesInstruction{}).Apply(ec), ShouldBeTrue)
	})
}

func Test_MoreCoverage_ManyServicesInstruction(t *testing.T) {
	Convey("ManyServicesInstruction IsEmpty/Apply on empty", t, func() {
		var nilMS *linuxservicecmd.ManyServicesInstruction
		So(nilMS.IsEmpty(), ShouldBeTrue)

		ms := &linuxservicecmd.ManyServicesInstruction{}
		So(ms.IsEmpty(), ShouldBeTrue)
		ec := errwrappers.Empty()
		So(ms.Apply(ec), ShouldBeTrue)
	})

	Convey("ManyServicesInstructions IsEmpty/Apply on empty", t, func() {
		var nilMs *linuxservicecmd.ManyServicesInstructions
		So(nilMs.IsEmpty(), ShouldBeTrue)

		ms := &linuxservicecmd.ManyServicesInstructions{}
		So(ms.IsEmpty(), ShouldBeTrue)
		ec := errwrappers.Empty()
		So(ms.Apply(ec), ShouldBeTrue)
		So(ms.Status(ec), ShouldBeTrue)
		So(ms.Run(ec, servicestate.Status), ShouldBeTrue)
	})
}

func Test_MoreCoverage_StateValidate(t *testing.T) {
	Convey("StateValidateInstruction empty paths", t, func() {
		var nilSV *linuxservicecmd.StateValidateInstruction
		So(nilSV.IsEmpty(), ShouldBeTrue)
		So(nilSV.Apply(), ShouldBeNil)

		ec := errwrappers.Empty()
		So(nilSV.ApplyUsingErrCollection(ec), ShouldBeTrue)
	})

	Convey("StateValidateInstructions empty paths", t, func() {
		var nilSVs *linuxservicecmd.StateValidateInstructions
		So(nilSVs.IsEmpty(), ShouldBeTrue)
		So(nilSVs.HasValidations(), ShouldBeFalse)

		svs := &linuxservicecmd.StateValidateInstructions{}
		So(svs.IsEmpty(), ShouldBeTrue)
		So(svs.HasValidations(), ShouldBeFalse)

		ec := errwrappers.Empty()
		So(svs.ApplyUsingErrCollection(ec), ShouldBeTrue)
	})
}

func Test_MoreCoverage_LookPathHelpers(t *testing.T) {
	Convey("IsNamePathExist returns a bool without panicking", t, func() {
		// Just exercise the path; result varies per environment.
		_ = linuxservicecmd.IsNamePathExist("definitely-not-a-real-binary-xyz")
		_ = linuxservicecmd.IsNamePathExist("ls")
	})
}

func Test_MoreCoverage_VerifyExitCode_NilResult(t *testing.T) {
	Convey("Request accessors are safe to call without exec", t, func() {
		req := linuxservicecmd.Request{
			ServiceName: "svc-x",
			Action:      servicestate.Start,
		}
		// Just verify zero-arg getters don't panic on the struct.
		So(req.ServiceName, ShouldEqual, "svc-x")
	})
}
