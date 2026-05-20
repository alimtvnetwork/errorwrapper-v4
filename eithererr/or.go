package eithererr

import (
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

func Or(
	isCondition bool,
	trueWrapper,
	falseWrapper *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if isCondition {
		return trueWrapper
	}

	return falseWrapper
}

func OrEmpty(
	isCondition bool,
	trueWrapper *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if isCondition {
		return trueWrapper
	}

	return nil
}

func OrCollectionPtr(
	isCondition bool,
	trueWrappers,
	falseWrappers *errwrappers.Collection,
) *errwrappers.Collection {
	if isCondition {
		return trueWrappers
	}

	return falseWrappers
}

func OrEmptyCollectionPtr(
	isCondition bool,
	trueWrappers *errwrappers.Collection,
) *errwrappers.Collection {
	if isCondition {
		return trueWrappers
	}

	return errwrappers.Empty()
}

// AnyFirstOrEmpty
//
//  OrNull Return which first wrapper which has error
//  Return null based
func AnyFirstOrEmpty(
	wrappers ...*errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if len(wrappers) == 0 {
		return nil
	}

	for _, wrapper := range wrappers {
		if wrapper.HasError() {
			return wrapper
		}
	}

	return nil
}

// AnyErrInfFirstOrEmpty
//
//  OrNull Return which first wrapper which has error
//  Return null based
func AnyErrInfFirstOrEmpty(
	errInterfaces ...errcoreinf.BaseErrorOrCollectionWrapper,
) *errorwrapper.Wrapper {
	if len(errInterfaces) == 0 {
		return nil
	}

	for _, errInf := range errInterfaces {
		if errInf == nil || errInf.IsEmpty() {
			continue
		}

		if errInf.HasAnyError() {
			return errorwrapper.CastInterfaceToErrorWrapper(
				errInf)
		}
	}

	return nil
}

// AnyBasicErrFirstOrEmpty
//
//  OrNull Return which first wrapper which has error
//  Return null based
func AnyBasicErrFirstOrEmpty(
	basicErrors ...errcoreinf.BasicErrWrapper,
) *errorwrapper.Wrapper {
	if len(basicErrors) == 0 {
		return nil
	}

	for _, basicErr := range basicErrors {
		if basicErr == nil || basicErr.IsEmpty() {
			continue
		}

		if basicErr.HasAnyError() {
			return errorwrapper.NewUsingBasicErr(
				basicErr)
		}
	}

	return nil
}

// ExecAnyFunc
//
// either one of the function returns error first
// will block the other to execute and returns the first one
func ExecAnyFunc(
	functions ...func() *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if len(functions) == 0 {
		return nil
	}

	for _, currentFunc := range functions {
		errW := currentFunc()
		if errW != nil && errW.HasError() {
			return errW
		}
	}

	return nil
}

// ExecAllFunc runs all and returns the final error
func ExecAllFunc(
	functions ...func() *errorwrapper.Wrapper,
) []*errorwrapper.Wrapper {
	if len(functions) == 0 {
		return nil
	}

	var errorCollection []*errorwrapper.Wrapper

	for _, currentFunc := range functions {
		errW := currentFunc()
		if errW != nil && errW.HasError() {
			errorCollection = append(
				errorCollection,
				errW)
		}
	}

	return errorCollection
}

func WrapperOrFunc(
	wrapper *errorwrapper.Wrapper,
	errWrapperFunc func() *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if wrapper != nil {
		return wrapper
	}

	return errWrapperFunc()
}
