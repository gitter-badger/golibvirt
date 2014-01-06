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

//UUID
const (
	VIR_UUID_STRING_BUFLEN = C.VIR_UUID_STRING_BUFLEN
)

//virDomainUndefineFlagsValues
const (
	VIR_DOMAIN_UNDEFINE_NOFLAGS            = 0
	VIR_DOMAIN_UNDEFINE_MANAGED_SAVE       = C.VIR_DOMAIN_UNDEFINE_MANAGED_SAVE       //Also remove any managed save
	VIR_DOMAIN_UNDEFINE_SNAPSHOTS_METADATA = C.VIR_DOMAIN_UNDEFINE_SNAPSHOTS_METADATA //If last use of domain, then also remove any snapshot metadata Future undefine control flags should come here.
)

//virDomainState
const (
	VIR_DOMAIN_NOSTATE     = C.VIR_DOMAIN_NOSTATE     //no state
	VIR_DOMAIN_RUNNING     = C.VIR_DOMAIN_RUNNING     //the domain is running
	VIR_DOMAIN_BLOCKED     = C.VIR_DOMAIN_BLOCKED     //the domain is blocked on resource
	VIR_DOMAIN_PAUSED      = C.VIR_DOMAIN_PAUSED      //the domain is paused by user
	VIR_DOMAIN_SHUTDOWN    = C.VIR_DOMAIN_SHUTDOWN    //the domain is being shut down
	VIR_DOMAIN_SHUTOFF     = C.VIR_DOMAIN_SHUTOFF     //the domain is shut off
	VIR_DOMAIN_CRASHED     = C.VIR_DOMAIN_CRASHED     //the domain is crashed
	VIR_DOMAIN_PMSUSPENDED = C.VIR_DOMAIN_PMSUSPENDED //the domain is suspended by guest power management
	//VIR_DOMAIN_LAST        = C.VIR_DOMAIN_LAST        //NB: this enum value will increase over time as new events are added to the libvirt API. It reflects the last state supported by this version of the libvirt API.
)

//virDomainModificationImpact
const (
	VIR_DOMAIN_AFFECT_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT //Affect current domain state.
	VIR_DOMAIN_AFFECT_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE    //Affect running domain state.
	VIR_DOMAIN_AFFECT_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG  //Affect persistent domain state. 1 << 2 is reserved for virTypedParameterFlags
)

//virDomainMemoryModFlags
const (
	VIR_DOMAIN_MEM_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG //Additionally, these flags may be bitwise-OR'd in
	VIR_DOMAIN_MEM_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT
	VIR_DOMAIN_MEM_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE
	VIR_DOMAIN_MEM_MAXIMUM = C.VIR_DOMAIN_MEM_MAXIMUM //affect Max rather than current
)

//virDomainVcpuFlags
const (
	VIR_DOMAIN_VCPU_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG //Additionally, these flags may be bitwise-OR'd in.
	VIR_DOMAIN_VCPU_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT
	VIR_DOMAIN_VCPU_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE
	VIR_DOMAIN_VCPU_MAXIMUM = C.VIR_DOMAIN_VCPU_MAXIMUM //Max rather than current count
	VIR_DOMAIN_VCPU_GUEST   = C.VIR_DOMAIN_VCPU_GUEST   //Modify state of the cpu in the guest
)

//virDomainShutdownFlagValues
const (
	VIR_DOMAIN_SHUTDOWN_DEFAULT        = C.VIR_DOMAIN_SHUTDOWN_DEFAULT        //hypervisor choice
	VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN = C.VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN //Send ACPI event
	VIR_DOMAIN_SHUTDOWN_GUEST_AGENT    = C.VIR_DOMAIN_SHUTDOWN_GUEST_AGENT    //Use guest agent
	VIR_DOMAIN_SHUTDOWN_INITCTL        = C.VIR_DOMAIN_SHUTDOWN_INITCTL        //Use initctl
	VIR_DOMAIN_SHUTDOWN_SIGNAL         = C.VIR_DOMAIN_SHUTDOWN_SIGNAL         //Send a signal
)

//virKeycodeSet
const (
	VIR_KEYCODE_SET_LINUX  = C.VIR_KEYCODE_SET_LINUX
	VIR_KEYCODE_SET_XT     = C.VIR_KEYCODE_SET_XT
	VIR_KEYCODE_SET_ATSET1 = C.VIR_KEYCODE_SET_ATSET1
	VIR_KEYCODE_SET_ATSET2 = C.VIR_KEYCODE_SET_ATSET2
	VIR_KEYCODE_SET_ATSET3 = C.VIR_KEYCODE_SET_ATSET3
	VIR_KEYCODE_SET_OSX    = C.VIR_KEYCODE_SET_OSX
	VIR_KEYCODE_SET_XT_KBD = C.VIR_KEYCODE_SET_XT_KBD
	VIR_KEYCODE_SET_USB    = C.VIR_KEYCODE_SET_USB
	VIR_KEYCODE_SET_WIN32  = C.VIR_KEYCODE_SET_WIN32
	VIR_KEYCODE_SET_RFB    = C.VIR_KEYCODE_SET_RFB
)

//virDomainDeviceModifyFlags
const (
	VIR_DOMAIN_DEVICE_MODIFY_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG
	VIR_DOMAIN_DEVICE_MODIFY_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT
	VIR_DOMAIN_DEVICE_MODIFY_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE
	VIR_DOMAIN_DEVICE_MODIFY_FORCE   = C.VIR_DOMAIN_DEVICE_MODIFY_FORCE //Forcibly modify device (ex. force eject a cdrom)
)

//virDomainXMLFlags
const (
	VIR_DOMAIN_XML_SECURE     = C.VIR_DOMAIN_XML_SECURE     //dump security sensitive information too
	VIR_DOMAIN_XML_INACTIVE   = C.VIR_DOMAIN_XML_INACTIVE   //dump inactive domain information
	VIR_DOMAIN_XML_UPDATE_CPU = C.VIR_DOMAIN_XML_UPDATE_CPU //update guest CPU requirements according to host CPU
	VIR_DOMAIN_XML_MIGRATABLE = C.VIR_DOMAIN_XML_MIGRATABLE //dump XML suitable for migration
)

type Domain struct {
	cptr C.virDomainPtr
}

type DomainInfo struct {
	State     uint   //the running state, one of virDomainState
	MaxMemory uint64 //the maximum memory in KBytes allowed
	Memory    uint64 //the memory in KBytes used by the domain
	Vcpus     uint8  //the number of virtual CPUs for the domain
	CpuTime   uint64 //the CPU time used in nanoseconds
}

type DomainJobInfo struct {
	JobType       int //Time is measured in milliseconds
	TimeElapsed   uint64
	TimeRemaining uint64
	DataTotal     uint64
	DataProcessed uint64
	DataRemaining uint64
	MemTotal      uint64
	MemProcessed  uint64
	MemRemaining  uint64
	FileTotal     uint64
	FileProcessed uint64
	FileRemaining uint64
}

func cleanupDomain(domain *Domain) {
	if domain.cptr != nil {
		C.virDomainFree(domain.cptr)
	}
}

func newDomain(cptr C.virDomainPtr) *Domain {
	domain := &Domain{cptr}
	runtime.SetFinalizer(domain, cleanupDomain)
	return domain
}

func (d *Domain) Undefine(flags uint) error {
	result := C.virDomainUndefineFlags(d.cptr, C.uint(flags))
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

func (d *Domain) GetUUID() (string, error) {
	uuid := make([]C.char, VIR_UUID_STRING_BUFLEN)
	result := C.virDomainGetUUIDString(d.cptr, &uuid[0])
	if result == -1 {
		return "", GetLastError()
	}

	return C.GoString(&uuid[0]), nil
}

func (d *Domain) GetId() (uint, error) {
	id := C.virDomainGetID(d.cptr)
	if int(id) == -1 {
		return 0, GetLastError()
	}
	return uint(id), nil
}

func (d *Domain) GetInfo() (DomainInfo, error) {
	var cDomainInfo C.virDomainInfo
	result := C.virDomainGetInfo(d.cptr, &cDomainInfo)
	if result == -1 {
		return DomainInfo{}, GetLastError()
	}

	return DomainInfo{
		State:     uint(cDomainInfo.state),
		MaxMemory: uint64(cDomainInfo.maxMem),
		Memory:    uint64(cDomainInfo.memory),
		Vcpus:     uint8(cDomainInfo.nrVirtCpu),
		CpuTime:   uint64(cDomainInfo.cpuTime),
	}, nil
}

func (d *Domain) GetAutostart() (bool, error) {
	var autostart C.int
	result := C.virDomainGetAutostart(d.cptr, &autostart)
	if result == -1 {
		return false, GetLastError()
	}

	if autostart == 0 {
		return true, nil
	}

	return false, nil
}

func (d *Domain) SetAutostart(autostart bool) error {
	var cautostart int
	if autostart {
		cautostart = 0
	} else {
		cautostart = 1
	}

	result := C.virDomainSetAutostart(d.cptr, C.int(cautostart))
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) GetOsType() (string, error) {
	result := C.virDomainGetOSType(d.cptr)
	if result == nil {
		return "", GetLastError()
	}

	osType := C.GoString(result)
	defer C.free(unsafe.Pointer(result))

	return osType, nil
}

func (d *Domain) GetMaxMemory() (uint64, error) {
	result := C.virDomainGetMaxMemory(d.cptr)
	if result == 0 {
		return uint64(0), GetLastError()
	}

	return uint64(result), nil
}

func (d *Domain) SetMaxMemory(memory uint64) error {
	result := C.virDomainSetMaxMemory(d.cptr, C.ulong(memory))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *Domain) SetMemory(memory uint64) error {
	result := C.virDomainSetMemory(d.cptr, C.ulong(memory))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *Domain) GetMaxVcpus() (int, error) {
	result := C.virDomainGetMaxVcpus(d.cptr)
	if result == -1 {
		return 0, GetLastError()
	}

	return int(result), nil
}

func (d *Domain) IsActive() (bool, error) {
	result := C.virDomainIsActive(d.cptr)
	if result == -1 {
		return false, GetLastError()
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

func (d *Domain) IsPersistent() (bool, error) {
	result := C.virDomainIsPersistent(d.cptr)
	if result == -1 {
		return false, GetLastError()
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

func (d *Domain) IsUpdated() (bool, error) {
	result := C.virDomainIsUpdated(d.cptr)
	if result == -1 {
		return false, GetLastError()
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

func (d *Domain) Reboot(flags uint) error {
	result := C.virDomainReboot(d.cptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) Reset() error {
	result := C.virDomainReset(d.cptr, 0)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) Save(filepath string) error {
	cfilepath := C.CString(filepath)
	defer C.free(unsafe.Pointer(cfilepath))

	result := C.virDomainSave(d.cptr, cfilepath)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) Suspend() error {
	result := C.virDomainSuspend(d.cptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) Resume() error {
	result := C.virDomainResume(d.cptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) Shutdown() error {
	result := C.virDomainShutdown(d.cptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) Start() error {
	result := C.virDomainCreate(d.cptr)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) SendKey(codeset uint, holdtime uint, keycodes []uint, flags uint) error {
	length := len(keycodes)
	ckeycodes := make([]C.uint, length)

	for i, kc := range keycodes {
		ckeycodes[i] = C.uint(kc)
	}

	result := C.virDomainSendKey(
		d.cptr,
		C.uint(codeset),
		C.uint(holdtime),
		&ckeycodes[0],
		C.int(length),
		C.uint(flags),
	)

	if result == -1 {
		return GetLastError()
	}
	return nil
}

//TODO
func (d *Domain) GetVcpus() {}

func (d *Domain) SetVcpus(vcpus uint8, flags uint16) error {
	result := C.virDomainSetVcpusFlags(d.cptr, C.uint(vcpus), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}

	return nil
}

//TODO
func (d *Domain) Migrate()                 {}
func (d *Domain) SetMigrationMaxDowntime() {}
func (d *Domain) PinVcpu()                 {}

func (d *Domain) AttachDevice(xml string) error {
	cxml := C.CString(xml)
	defer C.free(unsafe.Pointer(cxml))

	result := C.virDomainAttachDevice(d.cptr, cxml)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) DetachDevice(xml string) error {
	cxml := C.CString(xml)
	defer C.free(unsafe.Pointer(cxml))

	result := C.virDomainDetachDevice(d.cptr, cxml)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) UpdateDevice(xml string, flags uint16) error {
	cxml := C.CString(xml)
	defer C.free(unsafe.Pointer(cxml))

	result := C.virDomainUpdateDeviceFlags(d.cptr, cxml, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (d *Domain) ToXml(flags uint16) (string, error) {
	result := C.virDomainGetXMLDesc(d.cptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}

	xml := C.GoString(result)
	defer C.free(unsafe.Pointer(result))

	return xml, nil
}

func (d *Domain) GetState(flags uint16) (int, int, error) {
	var cstate, creason C.int
	result := C.virDomainGetState(d.cptr, &cstate, &creason, C.uint(flags))
	if result == -1 {
		return 0, 0, GetLastError()
	}

	return int(cstate), int(creason), nil
}

func (d *Domain) GetJobInfo() (*DomainJobInfo, error) {
	var cJobInfo C.virDomainJobInfo
	result := C.virDomainGetJobInfo(d.cptr, &cJobInfo)
	if result == -1 {
		return nil, GetLastError()
	}

	return &DomainJobInfo{
		int(cJobInfo._type),
		uint64(cJobInfo.timeElapsed),
		uint64(cJobInfo.timeRemaining),
		uint64(cJobInfo.dataTotal),
		uint64(cJobInfo.dataProcessed),
		uint64(cJobInfo.dataRemaining),
		uint64(cJobInfo.memTotal),
		uint64(cJobInfo.memProcessed),
		uint64(cJobInfo.memRemaining),
		uint64(cJobInfo.memTotal),
		uint64(cJobInfo.memProcessed),
		uint64(cJobInfo.memRemaining),
	}, nil

}

func (d *Domain) GetJobStats() (*TypedParameters, error) {
	params := new(TypedParameters)
	var jobType C.int
	result := C.virDomainGetJobStats(d.cptr, &jobType, &params.cptr, &params.length, 0)

	if result == -1 {
		return nil, GetLastError()
	}
	return params, nil
}

func (d *Domain) AbortCurrentJob() error {
	result := C.virDomainAbortJob(d.cptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

//TODO
func (d *Domain) GetSchedType()       {}
func (d *Domain) GetSchedParams()     {}
func (d *Domain) SetSchedParams()     {}
func (d *Domain) GetSecurityLabel()   {}
func (d *Domain) SaveManagedImage()   {}
func (d *Domain) RemoveManagedImage() {}
func (d *Domain) HasManagedImage()    {}
func (d *Domain) MemoryPeek()         {}
func (d *Domain) GetMemoryStats()     {}
func (d *Domain) BlockPeek()          {}
func (d *Domain) GetBlockStats()      {}
func (d *Domain) GetBlockInfo()       {}
func (d *Domain) CoreDump()           {}
func (d *Domain) GetInterfaceStats()  {}

func (d *Domain) HasCurrentSnapshot(flags uint16) (bool, error) {
	result := C.virDomainHasCurrentSnapshot(d.cptr, C.uint(flags))
	if result == -1 {
		return false, GetLastError()
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

//TODO
func (d *Domain) RevertToSnapshot()           {}
func (d *Domain) TakeSnapshot()               {}
func (d *Domain) GetCurrentSnapshot()         {}
func (d *Domain) DeleteSnapshot()             {}
func (d *Domain) LookupSnapshotByName()       {}
func (d *Domain) GetSnapshots()               {}
func (d *Domain) BlockCommit()                {}
func (d *Domain) BlockJobAbort()              {}
func (d *Domain) BlockJobSetSpeed()           {}
func (d *Domain) BlockPull()                  {}
func (d *Domain) BlockRebase()                {}
func (d *Domain) BlockResize()                {}
func (d *Domain) FSTrim()                     {}
func (d *Domain) GetBlockIoTune()             {}
func (d *Domain) GetBlockJobInfo()            {}
func (d *Domain) GetCpuStats()                {}
func (d *Domain) GetControlInfo()             {}
func (d *Domain) GetDiskErrors()              {}
func (d *Domain) GetEmulatorPinInfo()         {}
func (d *Domain) GetHostname()                {}
func (d *Domain) GetMetadata()                {}
func (d *Domain) SetMetadata()                {}
func (d *Domain) InjectNmi()                  {}
func (d *Domain) ManagedSave()                {}
func (d *Domain) ManagedSaveRemove()          {}
func (d *Domain) MigrateGetCompressionCache() {}
func (d *Domain) MigrateSetCompressionCache() {}
func (d *Domain) MigrateGetMaxSpeed()         {}
func (d *Domain) MigrateSetMaxSpeed()         {}
func (d *Domain) OpenChannel()                {}
func (d *Domain) OpenConsole()                {}
func (d *Domain) OpenGraphics()               {}
func (d *Domain) PMSuspendForDuration()       {}
func (d *Domain) PMWakUp()                    {}
func (d *Domain) PinEmulator()                {}
func (d *Domain) Screenshot()                 {}
func (d *Domain) SendProcessSignal()          {}
func (d *Domain) SetNumaParameters()          {}
