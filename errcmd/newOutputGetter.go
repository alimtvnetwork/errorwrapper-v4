package errcmd

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newOutputGetter struct{}

func (it newOutputGetter) CmdOutput(
	process string,
	arguments ...string,
) (allBytes []byte, err error) {
	allBytes, errWp := it.OutputErrWrapper(
		process,
		arguments...)

	return allBytes, errWp.CompiledErrorWithStackTraces()
}

func (it newOutputGetter) OutputErrWrapper(
	process string,
	arguments ...string,
) (allBytes []byte, errWp *errorwrapper.Wrapper) {
	cmd := exec.Command(
		process, arguments...)

	wholeCommand := ProcessArgsJoinAppend(
		process, arguments...)

	if cmd == nil {
		return nil, errnew.Message.Default(
			errtype.NotFoundProcess,
			wholeCommand)
	}

	buffer := &bytes.Buffer{}
	cmd.Stderr = buffer
	rawErrCollection := errcore.RawErrCollection{}

	allBytes, err := cmd.Output()

	rawErrCollection.AddError(err)
	if buffer.Len() > 0 {
		rawErrCollection.AddString(buffer.String())
	}

	if rawErrCollection.HasError() {
		return allBytes, errnew.Error.TypeMsg(
			errtype.FailedProcess,
			rawErrCollection.CompiledError(),
			wholeCommand)
	}

	return allBytes, nil
}

func (it newOutputGetter) OutputStringErrWrapper(
	process string,
	arguments ...string,
) (output string, errWp *errorwrapper.Wrapper) {
	allBytes, errWp := it.OutputErrWrapper(
		process,
		arguments...)

	if len(allBytes) == 0 {
		return "", errWp
	}

	return string(allBytes), errWp
}

func (it newOutputGetter) OutputStringsErrWrapper(
	process string,
	arguments ...string,
) (outputLines []string, errWp *errorwrapper.Wrapper) {
	output, errWp := it.OutputStringErrWrapper(
		process,
		arguments...)

	if output == "" {
		return []string{}, errWp
	}

	return strings.Split(output, constants.DefaultLine), errWp
}
