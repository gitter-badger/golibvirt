package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
)

type NodeDevice struct {
	cptr C.virNodeDevicePtr
}

func cleanupNodeDevice(device *NodeDevice) {
	C.virNodeDeviceFree(device.cptr)
}

func newNodeDevice(cptr C.virNodeDevicePtr) *NodeDevice {
	device := &NodeDevice{cptr}
	runtime.SetFinalizer(device, cleanupNodeDevice)
	return device
}

func (n *NodeDevice) GetName() (string, error) {
	result := C.virNodeDeviceGetName(n.cptr)
	if result == nil {
		return "", GetLastError()
	}

	name := C.GoString(result)
	return name, nil
}
