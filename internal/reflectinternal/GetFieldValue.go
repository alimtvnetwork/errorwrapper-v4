package reflectinternal

import "reflect"

// GetFieldValue Example
//  actualValueOf := reflect.ValueOf(*actual)
//	for i := 0; i < actualValueOf.NumField(); i++ {
//		actualFieldValue := GetFieldValue(actualValueOf.Field(i))
//		expectedFieldValue := GetFieldValue(reflect.ValueOf(expectedNewEmptyDirectoryResult).Field(i))
func GetFieldValue(field reflect.Value) interface{} {
	isPtr := field.Kind() == reflect.Ptr
	isNil := isPtr && field.IsNil()
	if isNil {
		return nil
	}

	if !field.IsValid() && !isPtr {
		return nil
	}

	if field.IsValid() && field.CanAddr() && !isPtr {
		return field.Interface()
	}

	if field.CanInterface() {
		return field.Interface()
	}

	if isPtr && !isNil {
		return field.Elem()
	}

	return field.Interface()
}
