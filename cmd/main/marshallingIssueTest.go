package main

import (
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errbool"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

func marshallingIssueTest() {
	rs := errbool.Empty.Result()
	rs.Value = true
	rs.ErrorWrapper = errnew.Type.Create(errtype.CmdOnceFailed)
	fmt.Println(rs.ErrorWrapper.Error())
	json := rs.JsonPtr()
	rs2 := errbool.Result{}
	rs2.JsonParseSelfInject(json)
	fmt.Println(rs2.ErrorWrapper.ErrorString())
	json2 := rs2.Json()
	fmt.Println(json.JsonString())
	fmt.Println(json2.JsonString())
}
