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

type StoragePool struct {
	cptr C.virStoragePoolPtr
}

func cleanupStoragePool(pool *StoragePool) {
	C.virStoragePoolFree(pool.cptr)
}

func newStoragePool(cptr C.virStoragePoolPtr) *StoragePool {
	pool := &StoragePool{cptr}
	runtime.SetFinalizer(pool, cleanupStoragePool)
	return pool
}

func (n *StoragePool) GetName() (string, error) {
	result := C.virStoragePoolGetName(n.cptr)
	if result == nil {
		return "", GetLastError()
	}

	name := C.GoString(result)
	return name, nil
}
