package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
