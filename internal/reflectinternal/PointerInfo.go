package reflectinternal

import (
	"reflect"
)

type PointerInfo struct {
	IsPointer    bool
	ReflectValue reflect.Value
	Pointer      *uintptr
}
