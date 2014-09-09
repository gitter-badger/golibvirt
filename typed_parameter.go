// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <string.h>
#include <stdlib.h>

// Go Helpers

virTypedParameterPtr virTypedParameterAllocate(int nparams) {
    virTypedParameterPtr params = malloc(sizeof(virTypedParameter) * nparams);
    memset(params, 0, sizeof(virTypedParameter) * nparams);
    return params;
}

*/
import "C"

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

const (
	VIR_TYPED_PARAM_INT     = C.VIR_TYPED_PARAM_INT     /* integer case */
	VIR_TYPED_PARAM_UINT    = C.VIR_TYPED_PARAM_UINT    /* unsigned integer case */
	VIR_TYPED_PARAM_LLONG   = C.VIR_TYPED_PARAM_LLONG   /* long long case */
	VIR_TYPED_PARAM_ULLONG  = C.VIR_TYPED_PARAM_ULLONG  /* unsigned long long case */
	VIR_TYPED_PARAM_DOUBLE  = C.VIR_TYPED_PARAM_DOUBLE  /* double case */
	VIR_TYPED_PARAM_BOOLEAN = C.VIR_TYPED_PARAM_BOOLEAN /* boolean(character) case */
	VIR_TYPED_PARAM_STRING  = C.VIR_TYPED_PARAM_STRING  /* string case */
)

type TypedParameter struct {
	Name  string
	Type  int
	Value interface{}
}

type TypedParameters struct {
	cptr     C.virTypedParameterPtr
	length   C.int
	capacity C.int
}

func NewTypedParameters(nparams int) *TypedParameters {
	t := new(TypedParameters)
	if nparams > 0 {
		t.cptr = C.virTypedParameterAllocate(C.int(nparams))
		t.capacity = C.int(nparams)
	}
	runtime.SetFinalizer(t, func(t *TypedParameters) {
		if t.cptr != nil {
			C.virTypedParamsFree(t.cptr, t.length)
		}
	})
	return t
}

func (t *TypedParameters) Len() int {
	return int(t.length)
}

func (t *TypedParameters) Cap() int {
	return int(t.capacity)
}

// This is pretty dirty, and will probably go away.
// For now, only use for debugging

func (t *TypedParameters) GetTypedParameters() []TypedParameter {
	var pointerSlice []*C.virTypedParameter
	result := make([]TypedParameter, t.length)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&pointerSlice))
	header.Cap = int(t.capacity)
	header.Len = int(t.length)
	header.Data = uintptr(unsafe.Pointer(t.cptr))

	for i, v := range pointerSlice {
		typedParam := &result[i]
		typedParam.Name = string(C.GoBytes(unsafe.Pointer(&v.field[0]), C.VIR_TYPED_PARAM_FIELD_LENGTH))
		typedParam.Type = int(v._type)

		uvalue := binary.LittleEndian.Uint64(v.value[0:8])
		switch v._type {
		case VIR_TYPED_PARAM_INT:
			typedParam.Value = int(uvalue)
		case VIR_TYPED_PARAM_UINT:
			typedParam.Value = uint(uvalue)
		case VIR_TYPED_PARAM_LLONG:
			typedParam.Value = int64(uvalue)
		case VIR_TYPED_PARAM_ULLONG:
			typedParam.Value = uint64(uvalue)
		case VIR_TYPED_PARAM_DOUBLE:
			typedParam.Value = float64(uvalue)
		case VIR_TYPED_PARAM_BOOLEAN:
			if uvalue != 0 {
				typedParam.Value = true
			} else {
				typedParam.Value = false
			}
		case VIR_TYPED_PARAM_STRING:
			// Since Go's type system won't let us cast a uint64 to *C.char, some pointer hackery is required
			// Get an unsafe pointer to the union value
			uvalptr := unsafe.Pointer(&uvalue)
			// Cast to a pointer to a C-string, and dereference to get the C-string.
			cstring := *(**C.char)(uvalptr)
			// Convert to Go string.
			typedParam.Value = C.GoString(cstring)
		}
	}
	return result
}

func (t *TypedParameters) TypedParamsAddBool(val bool, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var tmp uint8
	if val {
		tmp = 1
	} else {
		tmp = 0
	}
	result := C.virTypedParamsAddBoolean(&t.cptr, &t.length, &t.capacity, cname, C.int(tmp))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (t *TypedParameters) TypedParamsAddFloat64(val float64, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.virTypedParamsAddDouble(&t.cptr, &t.length, &t.capacity, cname, C.double(val))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (t *TypedParameters) TypedParamsAddInt32(val int32, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.virTypedParamsAddInt(&t.cptr, &t.length, &t.capacity, cname, C.int(val))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (t *TypedParameters) TypedParamsAddInt64(val int64, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.virTypedParamsAddLLong(&t.cptr, &t.length, &t.capacity, cname, C.longlong(val))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (t *TypedParameters) TypedParamsAddString(val string, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	tmp := C.CString(val)
	result := C.virTypedParamsAddString(&t.cptr, &t.length, &t.capacity, cname, tmp)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (t *TypedParameters) TypedParamsAddUInt32(val uint32, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.virTypedParamsAddUInt(&t.cptr, &t.length, &t.capacity, cname, C.uint(val))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (t *TypedParameters) TypedParamsAddUInt64(val uint64, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	result := C.virTypedParamsAddULLong(&t.cptr, &t.length, &t.capacity, cname, C.ulonglong(val))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (t *TypedParameters) TypedParamsGetBool(name string) (bool, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cval C.int
	result := C.virTypedParamsGetBoolean(t.cptr, t.length, cname, &cval)

	if result > 0 {
		return int(cval) != 0, nil
	}
	if result == -1 {
		return false, GetLastError()
	}
	return false, fmt.Errorf("Can't locate boolean parameter: %s", name)

}

func (t *TypedParameters) TypedParamsGetFloat64(name string) (float64, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cval C.double
	result := C.virTypedParamsGetDouble(t.cptr, t.length, cname, &cval)

	if result > 0 {
		return float64(cval), nil
	}
	if result == -1 {
		return 0.0, GetLastError()
	}
	return 0.0, fmt.Errorf("Can't locate float64 parameter: %s", name)
}

func (t *TypedParameters) TypedParamsGetInt32(name string) (int32, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cval C.int
	result := C.virTypedParamsGetInt(t.cptr, t.length, cname, &cval)

	if result > 0 {
		return int32(cval), nil
	}
	if result == -1 {
		return 0, GetLastError()
	}
	return 0, fmt.Errorf("Can't locate int32 parameter: %s", name)
}

func (t *TypedParameters) TypedParamsGetInt64(name string) (int64, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cval C.longlong
	result := C.virTypedParamsGetLLong(t.cptr, t.length, cname, &cval)

	if result > 0 {
		return int64(cval), nil
	}
	if result == -1 {
		return 0, GetLastError()
	}
	return 0, fmt.Errorf("Can't locate int64 parameter: %s", name)
}

func (t *TypedParameters) TypedParamsGetUInt32(name string) (uint32, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cval C.uint
	result := C.virTypedParamsGetUInt(t.cptr, t.length, cname, &cval)

	if result > 0 {
		return uint32(cval), nil
	}
	if result == -1 {
		return 0, GetLastError()
	}
	return 0, fmt.Errorf("Can't locate uint32 parameter: %s", name)
}

func (t *TypedParameters) TypedParamsGetUInt64(name string) (uint64, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var cval C.ulonglong
	result := C.virTypedParamsGetULLong(t.cptr, t.length, cname, &cval)

	if result > 0 {
		return uint64(cval), nil
	}
	if result == -1 {
		return 0, GetLastError()
	}
	return 0, fmt.Errorf("Can't locate uint64 parameter: %s", name)
}
