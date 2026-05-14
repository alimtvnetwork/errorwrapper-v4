package trydo

func WrapPanic(voidFunc func()) Exception {
	var exception Exception

	Block{
		Try: func() {
			voidFunc()
		},
		Catch: func(e Exception) {
			exception = e
		},
		Finally: nil,
	}.Do()

	return exception
}
