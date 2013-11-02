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

type Hypervisor struct {
	conn C.virConnectPtr
}

func (h *Hypervisor) BaseLineCPU() (string, *LibvirtError) {
	return "", nil
}
