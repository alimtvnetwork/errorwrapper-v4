package creationtests

import (
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/coretests"
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
