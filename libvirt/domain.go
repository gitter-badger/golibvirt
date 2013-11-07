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
	"unsafe"
)

//virDomainCreateFlags
const (
	VIR_DOMAIN_NONE               = C.VIR_DOMAIN_NONE               //Default behavior
	VIR_DOMAIN_START_PAUSED       = C.VIR_DOMAIN_START_PAUSED       //Launch guest in paused state
	VIR_DOMAIN_START_AUTODESTROY  = C.VIR_DOMAIN_START_AUTODESTROY  //Automatically kill guest when hypervisor.CloseConnection() is called
	VIR_DOMAIN_START_BYPASS_CACHE = C.VIR_DOMAIN_START_BYPASS_CACHE //Avoid file system cache pollution
	VIR_DOMAIN_START_FORCE_BOOT   = C.VIR_DOMAIN_START_FORCE_BOOT   //Boot, discarding any managed save
)

type Domain struct {
	cptr C.virDomainPtr
	conn C.virConnectPtr
}

func cleanupDomain(domain *Domain) {
	C.virDomainFree(domain.cptr)
}

func newDomain(cptr C.virDomainPtr, conn C.virConnectPtr) *Domain {
	domain := &Domain{cptr, conn}
	runtime.SetFinalizer(domain, cleanupDomain)
	return domain
}

func createDomain(conn C.virConnectPtr, xml string, flags uint) (*Domain, error) {
	cxml := C.CString(xml)
	defer C.free(unsafe.Pointer(cxml))

	cdomain := C.virDomainCreateXML(conn, cxml, C.uint(flags))
	if cdomain == nil {
		return nil, GetLastError()
	}

	domain := &Domain{cdomain, conn}
	runtime.SetFinalizer(domain, cleanupDomain)

	return domain, nil
}

func (d *Domain) GetName() (string, error) {
	result := C.virDomainGetName(d.cptr)
	if result == nil {
		return "", GetLastError()
	}

	name := C.GoString(result)
	return name, nil
}
