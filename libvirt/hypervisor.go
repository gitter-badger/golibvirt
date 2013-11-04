package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
#include <stdio.h>

int blah(const char** xmls, uint len) {
	int i;
	printf("========================\n");
	for (i = 0; i < len; i++) {
		printf("%s\n", xmls[i]);
	}
	printf("========================\n");
	return 0;
}
*/
import "C"

import (
	"unsafe"
)

//virConnectCredentialType
const (
	VIR_CRED_USERNAME     = C.VIR_CRED_USERNAME
	VIR_CRED_AUTHNAME     = C.VIR_CRED_AUTHNAME
	VIR_CRED_LANGUAGE     = C.VIR_CRED_LANGUAGE
	VIR_CRED_CNONCE       = C.VIR_CRED_CNONCE
	VIR_CRED_PASSPHRASE   = C.VIR_CRED_PASSPHRASE
	VIR_CRED_ECHOPROMPT   = C.VIR_CRED_ECHOPROMPT
	VIR_CRED_NOECHOPROMPT = C.VIR_CRED_NOECHOPROMPT
	VIR_CRED_REALM        = C.VIR_CRED_REALM
	VIR_CRED_EXTERNAL     = C.VIR_CRED_EXTERNAL
)

//virCPUCompareResult
const (
	VIR_CPU_COMPARE_ERROR        = C.VIR_CPU_COMPARE_ERROR
	VIR_CPU_COMPARE_INCOMPATIBLE = C.VIR_CPU_COMPARE_INCOMPATIBLE
	VIR_CPU_COMPARE_IDENTICAL    = C.VIR_CPU_COMPARE_IDENTICAL
	VIR_CPU_COMPARE_SUPERSET     = C.VIR_CPU_COMPARE_SUPERSET
)

//virSecretUsageType
const (
	VIR_SECRET_USAGE_TYPE_NONE   = C.VIR_SECRET_USAGE_TYPE_NONE
	VIR_SECRET_USAGE_TYPE_VOLUME = C.VIR_SECRET_USAGE_TYPE_VOLUME
)

//virDomainEventID
const (
	VIR_DOMAIN_EVENT_ID_LIFECYCLE       = C.VIR_DOMAIN_EVENT_ID_LIFECYCLE
	VIR_DOMAIN_EVENT_ID_REBOOT          = C.VIR_DOMAIN_EVENT_ID_REBOOT
	VIR_DOMAIN_EVENT_ID_RTC_CHANGE      = C.VIR_DOMAIN_EVENT_ID_RTC_CHANGE
	VIR_DOMAIN_EVENT_ID_WATCHDOG        = C.VIR_DOMAIN_EVENT_ID_WATCHDOG
	VIR_DOMAIN_EVENT_ID_IO_ERROR        = C.VIR_DOMAIN_EVENT_ID_IO_ERROR
	VIR_DOMAIN_EVENT_ID_GRAPHICS        = C.VIR_DOMAIN_EVENT_ID_GRAPHICS
	VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON = C.VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON

//VIR_DOMAIN_EVENT_ID_LAST
)

//virConnectBaselineCPU - Only works for Qemu
const (
	VIR_CONNECT_BASELINE_CPU_EXPAND_FEATURES = C.VIR_CONNECT_BASELINE_CPU_EXPAND_FEATURES //show all the features
)

type Hypervisor struct {
	cptr C.virConnectPtr
}

func NewHypervisor(uri string) (*Hypervisor, error) {
	cUri := C.CString(uri)
	defer C.free(unsafe.Pointer(cUri))

	cptr := C.virConnectOpen(cUri)
	if cptr == nil {
		return &Hypervisor{}, GetLastError()
	}
	return &Hypervisor{cptr}, nil
}

func (h *Hypervisor) CloseConnection() error {
	result := C.virConnectClose(h.cptr)
	defer func() {
		h.cptr = nil
	}()

	if result == -1 {
		return GetLastError()
	}

	return nil
}

//virConnect functions
func (h *Hypervisor) GetBaselineCPU(cpusXml []string, flags uint8) (string, error) {
	cpus := len(cpusXml)
	cXmls := make([]*C.char, cpus)

	for i, s := range cpusXml {
		cXmls[i] = C.CString(s)
		defer C.free(unsafe.Pointer(cXmls[i]))
	}

	result := C.virConnectBaselineCPU(h.cptr, &cXmls[0], C.uint(cpus), C.uint(flags))
	//C.blah(&cXmls[0], C.uint(cpus))

	if result == nil {
		return "", GetLastError()
	}

	return C.GoString(result), nil
}

func (h *Hypervisor) CompareCPU(xmlDesc string, flags uint8) (int, error) {
	cxml := C.CString(xmlDesc)
	defer C.free(unsafe.Pointer(cxml))

	result := C.virConnectCompareCPU(h.cptr, cxml, C.uint(flags))
	if result == -1 {
		return int(result), GetLastError()
	}

	return int(result), nil
}

func (h *Hypervisor) GetCapabilities() (string, error) {
	result := C.virConnectGetCapabilities(h.cptr)
	if result == nil {
		return "", GetLastError()
	}
	capabilities := C.GoString(result)
	defer C.free(unsafe.Pointer(result))

	return capabilities, nil
}

func (h *Hypervisor) GetHostname() (string, error) {
	result := C.virConnectGetHostname(h.cptr)
	if result == nil {
		return "", GetLastError()
	}

	hostname := C.GoString(result)
	defer C.free(unsafe.Pointer(result))

	return hostname, nil
}

func (h *Hypervisor) GetSysInfo(flags uint8) (string, error) {
	result := C.virConnectGetSysinfo(h.cptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}

	sysinfo := C.GoString(result)
	defer C.free(unsafe.Pointer(result))

	return sysinfo, nil
}

func (h *Hypervisor) GetType() (string, error) {
	result := C.virConnectGetType(h.cptr)
	if result == nil {
		return "", GetLastError()
	}

	htype := C.GoString(result)

	return htype, nil
}

func (h *Hypervisor) GetConnectionUri() (string, error) {
	result := C.virConnectGetURI(h.cptr)
	if result == nil {
		return "", GetLastError()
	}

	uri := C.GoString(result)
	defer C.free(unsafe.Pointer(result))

	return uri, nil
}

func (h *Hypervisor) GetVersion() (map[string]int, error) {
	version := map[string]int{
		"major": 0,
		"minor": 0,
		"rel":   0,
	}

	var hver int
	result := C.virConnectGetVersion(h.cptr, (*C.ulong)(unsafe.Pointer(&hver)))
	if result == -1 {
		return version, GetLastError()
	}

	version["major"] = hver / 1000000
	hver %= 1000000
	version["minor"] = hver / 1000
	version["rel"] = hver % 1000

	return version, nil
}

func (h *Hypervisor) GetLibVirtVersion() (map[string]int, error) {
	version := map[string]int{
		"major": 0,
		"minor": 0,
		"rel":   0,
	}

	var libver int
	result := C.virConnectGetLibVersion(h.cptr, (*C.ulong)(unsafe.Pointer(&libver)))
	if result == -1 {
		return version, GetLastError()
	}

	version["major"] = libver / 1000000
	libver %= 1000000
	version["minor"] = libver / 1000
	version["rel"] = libver % 1000

	return version, nil
}

func (h *Hypervisor) GetMaxVcpus(dtype string) (uint8, error) {
	result := C.virConnectGetMaxVcpus(h.cptr, C.CString(dtype))
	if result == -1 {
		return uint8(result), GetLastError()
	}

	return uint8(result), nil
}

func (h *Hypervisor) IsConnectionEncrypted() (bool, error) {
	result := C.virConnectIsEncrypted(h.cptr)

	if result == -1 {
		return false, GetLastError()
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

func (h *Hypervisor) IsConnectionSecure() (bool, error) {
	result := C.virConnectIsSecure(h.cptr)

	if result == -1 {
		return false, GetLastError()
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

func (h *Hypervisor) IsConnectionAlive() (bool, error) {
	result := C.virConnectIsAlive(h.cptr)

	if result == -1 {
		return false, GetLastError()
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

//virConnectList functions
func (h *Hypervisor) GetDefinedDomains()      {}
func (h *Hypervisor) GetDefinedInterfaces()   {}
func (h *Hypervisor) GetDefinedNetworks()     {}
func (h *Hypervisor) GetDefinedStoragePools() {}
func (h *Hypervisor) GetActiveDomains()       {}
func (h *Hypervisor) GetActiveInterfaces()    {}
func (h *Hypervisor) GetNetworkFilters()      {}
func (h *Hypervisor) GetActiveNetworks()      {}
func (h *Hypervisor) GetSecrets()             {}
func (h *Hypervisor) GetActiveStoragePools()  {}

//virConnectNumOf functions
func (h *Hypervisor) GetNumberOfDefinedDomains()      {}
func (h *Hypervisor) GetNumberOfDefinedInterfaces()   {}
func (h *Hypervisor) GetNumberOfDefinedNetworks()     {}
func (h *Hypervisor) GetNumberOfDefinedStoragePools() {}
func (h *Hypervisor) GetNumberOfActiveDomains()       {}
func (h *Hypervisor) GetNumberOfActiveInterfaces()    {}
func (h *Hypervisor) GetNumberOfActiveNetworks()      {}
func (h *Hypervisor) GetNumberOfNetworkFilters()      {}
func (h *Hypervisor) GetNumberOfSecrets()             {}
func (h *Hypervisor) GetNumberOfActiveStoragePools()  {}

//Node functions
func (h *Hypervisor) GetNodeFreeMemoryInNumaCells() {}
func (h *Hypervisor) GetNodeFreeMemory()            {}
func (h *Hypervisor) GetNodeInfo()                  {}
func (h *Hypervisor) GetNodeDevicesNames()          {}
func (h *Hypervisor) GetNodeSecurityModel()         {}

//Event functions
func (h *Hypervisor) RegisterDomainEvent()   {}
func (h *Hypervisor) UnregisterDomainEvent() {}

//Misc functions
func (h *Hypervisor) FindStoragePoolSources() {}
