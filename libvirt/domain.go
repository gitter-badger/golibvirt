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

func defineDomain(conn C.virConnectPtr, xml string) (*Domain, error) {
	cxml := C.CString(xml)
	defer C.free(unsafe.Pointer(cxml))

	cdomain := C.virDomainDefineXML(conn, cxml)
	if cdomain == nil {
		return nil, GetLastError()
	}

	domain := &Domain{cdomain, conn}
	runtime.SetFinalizer(domain, cleanupDomain)

	return domain, nil
}

func restoreDomain(conn C.virConnectPtr, filepath string) error {
	cfilepath := C.CString(filepath)
	defer C.free(unsafe.Pointer(cfilepath))

	result := C.virDomainRestore(conn, cfilepath)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) GetName() (string, error) {
	result := C.virDomainGetName(d.cptr)
	if result == nil {
		return "", GetLastError()
	}

	name := C.GoString(result)
	return name, nil
}

func (d *Domain) lookupById()              {}
func (d *Domain) lookupByName()            {}
func (d *Domain) lookupByUUID()            {}
func (d *Domain) Undefine()                {}
func (d *Domain) GetId()                   {}
func (d *Domain) GetInfo()                 {}
func (d *Domain) GetUUID()                 {}
func (d *Domain) GetAutostart()            {}
func (d *Domain) SetAutostart()            {}
func (d *Domain) GetOsType()               {}
func (d *Domain) GetMaxMemory()            {}
func (d *Domain) SetMaxMemory()            {}
func (d *Domain) SetMemory()               {}
func (d *Domain) GetMaxVcpus()             {}
func (d *Domain) IsActive()                {}
func (d *Domain) IsPersistent()            {}
func (d *Domain) IsUpdated()               {}
func (d *Domain) Reboot()                  {}
func (d *Domain) Reset()                   {}
func (d *Domain) Save()                    {}
func (d *Domain) Restore()                 {}
func (d *Domain) Suspend()                 {}
func (d *Domain) Resume()                  {}
func (d *Domain) Shutdown()                {}
func (d *Domain) Start()                   {}
func (d *Domain) Destroy()                 {}
func (d *Domain) SendKey()                 {}
func (d *Domain) GetVcpus()                {}
func (d *Domain) SetVcpus()                {}
func (d *Domain) Migrate()                 {}
func (d *Domain) SetMigrationMaxDowntime() {}
func (d *Domain) PinVcpu()                 {}
func (d *Domain) AttachDevice()            {}
func (d *Domain) DetachDevice()            {}
func (d *Domain) UpdateDevice()            {}
func (d *Domain) ToXml()                   {}
func (d *Domain) GetJobInfo()              {}
func (d *Domain) AbortCurrentJob()         {}
func (d *Domain) GetSchedType()            {}
func (d *Domain) GetSchedParams()          {}
func (d *Domain) SetSchedParams()          {}
func (d *Domain) GetSecurityLabel()        {}
func (d *Domain) SaveManagedImage()        {}
func (d *Domain) RemoveManagedImage()      {}
func (d *Domain) HasManagedImage()         {}
func (d *Domain) MemoryPeek()              {}
func (d *Domain) GetMemoryStats()          {}
func (d *Domain) BlockPeek()               {}
func (d *Domain) GetBlockStats()           {}
func (d *Domain) GetBlockInfo()            {}
func (d *Domain) CoreDump()                {}
func (d *Domain) GetInterfaceStats()       {}
func (d *Domain) HasCurrentSnapshot()      {}
func (d *Domain) RevertToSnapshot()        {}
func (d *Domain) TakeSnapshot()            {}
func (d *Domain) GetCurrentSnapshot()      {}
func (d *Domain) DeleteSnapshot()          {}
func (d *Domain) LookupSnapshotByName()    {}
func (d *Domain) GetSnapshots()            {}
