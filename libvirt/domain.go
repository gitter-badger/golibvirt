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

type Domain struct {
	domain C.virDomainPtr
}
