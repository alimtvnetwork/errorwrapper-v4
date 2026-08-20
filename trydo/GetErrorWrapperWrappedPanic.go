package trydo

import "github.com/alimtvnetwork/errorwrapper-v4"

func GetErrorWrapperWrappedPanic(voidFunc func()) *errorwrapper.Wrapper {
	errWrapper := errorwrapper.StaticEmptyPtr

	Block{
		Try: func() {
			voidFunc()
		},
		Catch: func(e Exception) {
			errWp2, isOkay := e.(*errorwrapper.Wrapper)

			if isOkay {
				errWrapper = errWp2
			}
		},
		Finally: nil,
	}.Do()

	return errWrapper
}
