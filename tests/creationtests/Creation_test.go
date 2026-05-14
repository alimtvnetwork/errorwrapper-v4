package creationtests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/corevalidator"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/core/errcore"
)

func Test_Creation(t *testing.T) {
	for index, testCase := range testCases {
		// Arrange
		slice := corestr.New.SimpleSlice.Cap(constants.ArbitraryCapacity100)
		baseErrors := testCase.ArrangeAsBaseErrorOrCollectionWrapper()

		for errorIndex, baseError := range baseErrors {
			serializedWithoutTraces, err := baseError.SerializeWithoutTraces()
			errcore.HandleErr(err)

			slice.AppendFmt(
				"Case Index : %d, Error Index: %d, Full string: %s, Json : %s",
				index,
				errorIndex,
				baseError.FullString(),
				string(serializedWithoutTraces))
		}

		sliceValidator := corevalidator.SliceValidator{
			ActualLines:   slice.Items,
			ExpectedLines: testCase.ExpectedAsStrings(),
			CompareAs:     stringcompareas.Equal,
		}

		validatorParamBase := corevalidator.ValidatorParamsBase{
			CaseIndex:          index,
			IsAttachUserInputs: true,
			IsCaseSensitive:    true,
		}

		// Act
		validationFinalError := sliceValidator.AllVerifyError(
			&validatorParamBase)
		isValid := validationFinalError == nil

		// Assert
		convey.Convey(testCase.Title, t, func() {
			errcore.ErrPrintWithTestIndex(index, validationFinalError)
			convey.So(isValid, convey.ShouldBeTrue)
		})
	}
}
