package creationtests

import (
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coretests"
)

type TestCaseWrapper struct {
	coretests.BaseTestCase
}

func (it TestCaseWrapper) ExpectedAsStrings() []string {
	return it.Expected().([]string)
}

func (it TestCaseWrapper) ArrangeAsBaseErrorOrCollectionWrapper() []errcoreinf.BaseErrorOrCollectionWrapper {
	return it.ArrangeInput.([]errcoreinf.BaseErrorOrCollectionWrapper)
}
