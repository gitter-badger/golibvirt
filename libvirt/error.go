package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

// import (
// 	"unsafe"
// )

const (
	VIR_ERR_OK             = C.VIR_ERR_OK
	VIR_ERR_INTERNAL_ERROR = C.VIR_ERR_INTERNAL_ERROR
)

type LibvirtError struct {
	ptr     C.virErrorPtr
	Code    int
	Domain  int
	Message string
	Level   C.virErrorLevel
	Str1    string
	Str2    string
	Str3    string
	Int1    int
	Int2    int
}
