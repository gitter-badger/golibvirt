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

type Domain struct {
	cptr C.virDomainPtr
}

func cleanup(domain *Domain) {
	C.virDomainFree(domain.cptr)
}

func newDomain(cptr C.virDomainPtr) *Domain {
	domain := &Domain{cptr}
	runtime.SetFinalizer(domain, cleanup)
	return domain
}

func (d *Domain) GetName() (string, error) {
	result := C.virDomainGetName(d.cptr)
	if result == nil {
		return "", GetLastError()
	}

	name := C.GoString(result)
	return name, nil
}
