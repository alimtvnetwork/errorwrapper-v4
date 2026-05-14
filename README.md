# `errorwrapper` Intro

![Errorwrapper package](https://github.com/alimtvnetwork/errorwrapper-v3/uploads/431c560f98827024257c4ac542623e36/image.png)

[`errorwrapper`](https://github.com/alimtvnetwork/errorwrapper-v3) is wrapper that helps error handling smartly in go.

## Git Clone

`git clone https://github.com/alimtvnetwork/errorwrapper-v3.git`

### 2FA enabled, for linux

`git clone https://[YourGitLabUserName]:[YourGitlabAcessTokenGenerateFromGitlabsTokens]@github.com/alimtvnetwork/errorwrapper-v3.git`

### Prerequisites

- Update git to latest 2.29
- Update or install the latest of Go 1.15.2
- Either add your ssh key to your gitlab account
- Or, use your access token to clone it.

## Installation

`go get github.com/alimtvnetwork/errorwrapper-v3`

### Go get issue for private package

- Update git to 2.29
- Enable go modules. (Windows : `go env -w GO111MODULE=on`, Unix : `export GO111MODULE=on`)
- Add `gitlab.com/evatix-go` to go env private

To set for Windows:

`go env -w GOPRIVATE=[AddExistingOnes;]gitlab.com/evatix-go`

To set for Unix:

`expoort GOPRIVATE=[AddExistingOnes;]gitlab.com/evatix-go`

## Why `errorwrapper?`

It is to avoid the `if-else` part for checking errors not nil or handle error based if exits.

## Training Videos

- [Basics of ErrorWrapper](https://drive.google.com/drive/folders/1hprKXbgVd8swV0zJ-Ew0mAPohLGksuEc?usp=sharing)

## [Examples](https://github.com/alimtvnetwork/errorwrapper-v3/-/issues/24)

### Example 1

```go
errtype.OutOfRangeValue.Panic("something wrong", "alim", []int{1, 2})
```

![image](/uploads/186d7580df85394683028e74fdc6b56d/image.png)

### Example 2

```go
errtype.NotSupportInWindows.PanicNoRefs("not support in windows")
```

![image](/uploads/84ef65256bc60b29adc3621c0ab782f0/image.png)

### Example 3

```go
err2 := errnew.Messages.Many(errtype.Conversion, "I am not ready", "Convert failed.")
fmt.Println(1, err2.IsErrorEquals(err2.Error()))
fmt.Println(2, err2.GetTypeVariantStruct().Variant.IsConversion())
fmt.Println(3, err2.GetTypeVariantStruct().Variant.Is(errtype.Conversion))
fmt.Println(4, err2.GetTypeString())
fmt.Println(5, err2.GetTypeWithCodeString())
fmt.Println(6, err2.GetTypeVariantStruct().Variant.Is(errtype.NotSupportInWindows))
fmt.Println(7, err2.FullString())
fmt.Println(8, err2.IsErrorMessage(err2.FullString(), false))
fmt.Println(9, err2.IsErrorMessage(err2.Error().Error(), false))
```

![image](/uploads/088abb947c05deb102ce24d60cd700fe/image.png)

```
1 true
2 true
3 true
4 Conversion (Code - 69) : Conversion related error, cannot process the request.
5 (Code - #69) : Conversion
6 false
7 [Error (Conversion - #69): Conversion related error, cannot process the request.
"I am not ready, Convert failed."]
8 false
9 true

```

### Example 4

```go
err2 := errnew.Messages.Many(errtype.Conversion, "I am not ready", "Convert failed.")
fmt.Println(1, err2.FullString())
err2.HandleErrorWithRefs("hello", "var", []int{1, 2})
```

![image](/uploads/82ca80e8c7dbaa62bb5b07084942cbad/image.png)

```
1 [Error (Conversion - #69): Conversion related error, cannot process the request.
"I am not ready, Convert failed."]
panic: hello
[Error (Conversion - #69): Conversion related error, cannot process the request.
"I am not ready, Convert failed."] 
Reference { var([]int): [1 2] }
```

### Example 5 : JSON Parsing

```go
err := errnew.Refs(errtype.AnalysisFailed, "varName", []int{1, 20}, ",", "msg1")
collection := errwrappers.New(5)
collection.AddWrapperPtr(&err)
fmt.Println(collection.String())
jsonResult := collection.Json()

errorsCollection := errwrappers.New(1)
jsonResult.Unmarshal(&errorsCollection)

fmt.Println(errorsCollection.String())
```

![Serializing ErrorWrapper](https://github.com/alimtvnetwork/errorwrapper-v3/uploads/3f5b90840de73bff5d178ea8e055b75b/image.png)

```
# Error Wrappers - Collection - Length[1]

- [Error (AnalysisFailed - #138): Analysis is failed. Please consult with appropriate personnel to get the solution.
"msg1" Reference(s) {varName([]int) : [1 20]}.]
# Error Wrappers - Collection - Length[1]

- [Error (AnalysisFailed - #138): Analysis is failed. Please consult with appropriate personnel to get the solution.
"msg1" Reference(s) {varName(string) : [1 20]}.]

Process finished with exit code 0
```

## Acknowledgement

Any other packages used

## Links

## Issues

- [Create your issues](https://github.com/alimtvnetwork/errorwrapper-v3/-/issues)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)
