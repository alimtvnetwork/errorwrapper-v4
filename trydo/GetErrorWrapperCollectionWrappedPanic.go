package trydo

import "github.com/alimtvnetwork/errorwrapper-v4/errwrappers"

func GetErrorWrapperCollectionWrappedPanic(voidFunc func()) *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	Block{
		Try: func() {
			voidFunc()
		},
		Catch: func(e Exception) {
			errCollection2, isOkay := e.(*errwrappers.Collection)

			if isOkay {
				errCollection = errCollection2
			}
		},
		Finally: nil,
	}.Do()

	return errCollection
}
