package reflectinternal

import (
	"fmt"
	"reflect"
	"unsafe"
)

var (
	Uint8sType      = reflect.ValueOf([]byte{}).Type()
	BytesType       = Uint8sType
	StringsType     = reflect.ValueOf([]string{}).Type()
	StringType      = reflect.ValueOf("").Type()
	IntegersType    = reflect.ValueOf([]int{}).Type()
	IntegerType     = reflect.ValueOf(0).Type()
	BooleanType     = reflect.ValueOf(false).Type()
	Int64sType      = reflect.ValueOf([]int64{}).Type()
	Float64sType    = reflect.ValueOf([]float64{}).Type()
	BooleansType    = reflect.ValueOf([]bool{}).Type()
	AnyType         = reflect.ValueOf([]interface{}{}).Type()
	defaultMaxTries = 4
)

func IsTypeSame(type1 reflect.Type, type2 reflect.Type) bool {
	return type1 == type2
}

// GetElementType Calls GetElementTypeMaxTry with maxTry 4
//
// Examples:
//  - Cite : https://stackoverflow.com/a/39468423
//  - https://play.golang.org/p/hRQmtclUqkx
func GetElementType(any interface{}) reflect.Type {
	if any == nil {
		return nil
	}

	return GetElementTypeMaxTry(any, defaultMaxTries)
}

// GetElementTypeMaxTry Examples:
//  - Cite : https://stackoverflow.com/a/39468423
//  - https://play.golang.org/p/hRQmtclUqkx
func GetElementTypeMaxTry(any interface{}, maxTry int) reflect.Type {
	if any == nil {
		return nil
	}

	for t := reflect.TypeOf(any); maxTry >= 0; maxTry-- {
		switch t.Kind() {
		case reflect.Ptr, reflect.Slice:
			t = t.Elem()
		default:
			return t
		}
	}

	return nil
}

// GetElementTypesMaxTry Examples:
//  - Cite : https://stackoverflow.com/a/39468423
//  - https://play.golang.org/p/UswVNOUcjZi
func GetElementTypesMaxTry(any interface{}, maxTries int) (finalType reflect.Type, visitedTypes []reflect.Type) {
	if any == nil {
		return nil, nil
	}

	visitedTypes = make([]reflect.Type, 0, maxTries)

	for finalType = reflect.TypeOf(any); maxTries >= 0; maxTries-- {
		visitedTypes = append(visitedTypes, finalType)

		switch finalType.Kind() {
		case reflect.Ptr, reflect.Slice:
			finalType = finalType.Elem()
		default:
			return finalType, visitedTypes
		}
	}

	return finalType, visitedTypes
}

func GetTypeName(any interface{}) string {
	return fmt.Sprintf("%T", any)
}

func IsType(any interface{}, typeName string) bool {
	return fmt.Sprintf("%T", any) == typeName
}

func GetPointerInfo(any interface{}) PointerInfo {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		ptr := reflectValueOfAny.Pointer()
		return PointerInfo{
			IsPointer:    isPtr,
			ReflectValue: reflectValueOfAny,
			Pointer:      &ptr,
		}
	}

	return PointerInfo{
		IsPointer:    isPtr,
		ReflectValue: reflectValueOfAny,
		Pointer:      nil,
	}
}

// IsBytesOrBytesPointer Examples : https://play.golang.org/p/9XUt9Jf11WG | https://play.golang.org/p/oMAxxSzAP7F
func IsBytesOrBytesPointer(any interface{}) (isBytes bool, bytesPtr *[]byte) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == Uint8sType && indirectType.Kind() == reflect.Slice {
			bytes := (*[]byte)(unsafe.Pointer(reflectValueOfAny.Pointer()))

			return true, bytes
		}
	}

	typeOf := reflect.TypeOf(any)
	if typeOf == Uint8sType && typeOf.Kind() == reflect.Slice {
		bytes := reflectValueOfAny.Bytes()

		if bytes != nil {
			return true, &bytes
		}
	}

	return false, nil
}

// IsStringOrStringPointer Examples : https://play.golang.org/p/9XUt9Jf11WG | https://play.golang.org/p/oMAxxSzAP7F
func IsStringOrStringPointer(any interface{}) (isString bool, stringPointer *string) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == StringType {
			strPtr := (*string)(unsafe.Pointer(reflectValueOfAny.Pointer()))

			return true, strPtr
		}
	}

	typeOf := reflect.TypeOf(any)

	if typeOf == StringType {
		str, _ := any.(string)

		return true, &str
	}

	return false, nil
}

// IsStringsOrStringsPointer Examples : https://play.golang.org/p/9XUt9Jf11WG | https://play.golang.org/p/oMAxxSzAP7F
func IsStringsOrStringsPointer(any interface{}) (isStrings bool, stringsPtr *[]string) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == StringsType && indirectType.Kind() == reflect.Slice {
			strings := (*[]string)(unsafe.Pointer(reflectValueOfAny.Pointer()))

			return true, strings
		}
	}

	typeOf := reflect.TypeOf(any)

	if typeOf == StringsType && typeOf.Kind() == reflect.Slice {
		strings, _ := any.([]string)
		isStrings = true

		if strings != nil {
			return isStrings, &strings
		}

		return isStrings, nil
	}

	return false, nil
}

// IsIntegersOrIntegersPointer Examples : https://play.golang.org/p/9XUt9Jf11WG | https://play.golang.org/p/oMAxxSzAP7F
func IsIntegersOrIntegersPointer(any interface{}) (isIntegers bool, integersResults *[]int) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == IntegersType && indirectType.Kind() == reflect.Slice {
			integers := (*[]int)(unsafe.Pointer(reflectValueOfAny.Pointer()))

			return true, integers
		}
	}

	typeOf := reflect.TypeOf(any)

	if typeOf == IntegersType && typeOf.Kind() == reflect.Slice {
		integers, _ := any.([]int)
		isIntegers = true
		if integers != nil {
			return isIntegers, &integers
		}

		return isIntegers, nil
	}

	return false, nil
}

// IsIntegerOrIntegerPointer Examples : https://play.golang.org/p/9XUt9Jf11WG | https://play.golang.org/p/oMAxxSzAP7F
func IsIntegerOrIntegerPointer(any interface{}) (isInteger bool, intPtr *int) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == IntegerType {
			integer := (*int)(unsafe.Pointer(reflectValueOfAny.Pointer()))

			return true, integer
		}
	}

	typeOf := reflect.TypeOf(any)

	if typeOf == IntegerType {
		integer, _ := any.(int)

		return true, &integer
	}

	return false, nil
}

// IsInteger Examples : https://play.golang.org/p/9XUt9Jf11WG | https://play.golang.org/p/oMAxxSzAP7F
func IsInteger(any interface{}) (isInteger bool, intValue int) {
	isIntegerResult, intPtr := IsIntegerOrIntegerPointer(any)

	if intPtr == nil {
		return isIntegerResult, 0
	}

	return isIntegerResult, *intPtr
}

func IsByte(any interface{}) (isByte bool, result byte) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == IntegerType {
			currentByte := *(*byte)(unsafe.Pointer(reflectValueOfAny.Pointer()))

			return true, currentByte
		}
	}

	typeOf := reflect.TypeOf(any)

	if typeOf == IntegerType {
		currentByte, _ := any.(byte)

		return true, currentByte
	}

	return false, 0
}

func IsBoolean(any interface{}) (isBool bool, isResult bool) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == BooleanType {
			currentBoolPtr := (*bool)(unsafe.Pointer(reflectValueOfAny.Pointer()))

			if currentBoolPtr != nil {
				isResult = *currentBoolPtr
			}

			// pointer is nil so found = true, result = false.
			return true, isResult
		}
	}

	typeOf := reflect.TypeOf(any)

	if typeOf == BooleanType {
		currentBool, _ := any.(bool)

		return true, currentBool
	}

	return false, false
}

func IsBooleanPointer(any interface{}) (isBool bool, isResult *bool) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == BooleanType {
			return true, (*bool)(unsafe.Pointer(reflectValueOfAny.Pointer()))
		}
	}

	typeOf := reflect.TypeOf(any)

	if typeOf == BooleanType {
		currentBool, _ := any.(bool)

		return true, &currentBool
	}

	return false, nil
}

func IsFloat64sOrFloat64sPointer(any interface{}) (isFloat64 bool, floats *[]float64) {
	reflectValueOfAny := reflect.ValueOf(any)
	isPtr := reflectValueOfAny.Kind() == reflect.Ptr

	if isPtr {
		indirectType := reflect.Indirect(reflect.ValueOf(any)).Type()

		if indirectType == Float64sType {
			return true, (*[]float64)(unsafe.Pointer(reflectValueOfAny.Pointer()))
		}
	}

	typeOf := reflect.TypeOf(any)

	if typeOf == Float64sType {
		float64s, _ := any.([]float64)
		isFloat64 = true

		if float64s != nil {
			return isFloat64, &float64s
		}

		return isFloat64, nil
	}

	return false, nil
}

// IsString Examples : https://play.golang.org/p/9XUt9Jf11WG | https://play.golang.org/p/oMAxxSzAP7F
func IsString(any interface{}) (isStrings bool, str string) {
	isStr, strPtr := IsStringOrStringPointer(any)

	if strPtr != nil {
		return isStr, *strPtr
	}

	return isStr, ""
}
