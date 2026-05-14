package trydo

import "gitlab.com/evatix-go/errorwrapper"

func GetErrorWrapperWrappedPanic(voidFunc func()) *errorwrapper.Wrapper {
	errWrapper := errorwrapper.EmptyPtr()

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
