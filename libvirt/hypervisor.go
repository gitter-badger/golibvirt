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
	//"bytes"
	//"encoding/binary"
	// "fmt"
	"errors"
	"runtime"
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

//virConnectListAllDomainsFlags
const (
	VIR_CONNECT_LIST_DOMAINS_ACTIVE         = C.VIR_CONNECT_LIST_DOMAINS_ACTIVE
	VIR_CONNECT_LIST_DOMAINS_INACTIVE       = C.VIR_CONNECT_LIST_DOMAINS_INACTIVE
	VIR_CONNECT_LIST_DOMAINS_PERSISTENT     = C.VIR_CONNECT_LIST_DOMAINS_PERSISTENT
	VIR_CONNECT_LIST_DOMAINS_TRANSIENT      = C.VIR_CONNECT_LIST_DOMAINS_TRANSIENT
	VIR_CONNECT_LIST_DOMAINS_RUNNING        = C.VIR_CONNECT_LIST_DOMAINS_RUNNING
	VIR_CONNECT_LIST_DOMAINS_PAUSED         = C.VIR_CONNECT_LIST_DOMAINS_PAUSED
	VIR_CONNECT_LIST_DOMAINS_SHUTOFF        = C.VIR_CONNECT_LIST_DOMAINS_SHUTOFF
	VIR_CONNECT_LIST_DOMAINS_OTHER          = C.VIR_CONNECT_LIST_DOMAINS_OTHER
	VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE    = C.VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE
	VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE = C.VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE
	VIR_CONNECT_LIST_DOMAINS_AUTOSTART      = C.VIR_CONNECT_LIST_DOMAINS_AUTOSTART
	VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART   = C.VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART
	VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT   = C.VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT
	VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT    = C.VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT
)

//virConnectListAllInterfacesFlags
const (
	VIR_CONNECT_LIST_INTERFACES_INACTIVE = C.VIR_CONNECT_LIST_INTERFACES_INACTIVE
	VIR_CONNECT_LIST_INTERFACES_ACTIVE   = C.VIR_CONNECT_LIST_INTERFACES_ACTIVE
)

//virConnectListAllNetworksFlags
const (
	VIR_CONNECT_LIST_NETWORKS_INACTIVE     = C.VIR_CONNECT_LIST_NETWORKS_INACTIVE
	VIR_CONNECT_LIST_NETWORKS_ACTIVE       = C.VIR_CONNECT_LIST_NETWORKS_ACTIVE
	VIR_CONNECT_LIST_NETWORKS_PERSISTENT   = C.VIR_CONNECT_LIST_NETWORKS_PERSISTENT
	VIR_CONNECT_LIST_NETWORKS_TRANSIENT    = C.VIR_CONNECT_LIST_NETWORKS_TRANSIENT
	VIR_CONNECT_LIST_NETWORKS_AUTOSTART    = C.VIR_CONNECT_LIST_NETWORKS_AUTOSTART
	VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART = C.VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART
)

//virConnectListAllNodeDeviceFlags
const (
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SYSTEM        = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SYSTEM        //System capability
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_PCI_DEV       = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_PCI_DEV       //PCI device
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_DEV       = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_DEV       //USB device
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_INTERFACE = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_INTERFACE //USB interface
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_NET           = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_NET           //Network device
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_HOST     = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_HOST     //SCSI Host Bus Adapter
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_TARGET   = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_TARGET   //SCSI Target
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI          = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI          //SCSI device
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_STORAGE       = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_STORAGE       //Storage device
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_FC_HOST       = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_FC_HOST       //FC Host Bus Adapter
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_VPORTS        = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_VPORTS        //Capable of vport
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_GENERIC  = C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_GENERIC  //Capable of scsi_generic
)

//virConnectListAllSecretsFlags
const (
	VIR_CONNECT_LIST_SECRETS_EPHEMERAL    = C.VIR_CONNECT_LIST_SECRETS_EPHEMERAL    //kept in memory, never stored persistently
	VIR_CONNECT_LIST_SECRETS_NO_EPHEMERAL = C.VIR_CONNECT_LIST_SECRETS_NO_EPHEMERAL //Stored persistenly
	VIR_CONNECT_LIST_SECRETS_PRIVATE      = C.VIR_CONNECT_LIST_SECRETS_PRIVATE      //not revealed to any caller of libvirt, nor to any other node
	VIR_CONNECT_LIST_SECRETS_NO_PRIVATE   = C.VIR_CONNECT_LIST_SECRETS_NO_PRIVATE
)

//virConnectListAllStoragePoolsFlags
const (
	VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE     = C.VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE
	VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE       = C.VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE
	VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT   = C.VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT
	VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT    = C.VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT
	VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART    = C.VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART
	VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART = C.VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART //List pools by type
	VIR_CONNECT_LIST_STORAGE_POOLS_DIR          = C.VIR_CONNECT_LIST_STORAGE_POOLS_DIR
	VIR_CONNECT_LIST_STORAGE_POOLS_FS           = C.VIR_CONNECT_LIST_STORAGE_POOLS_FS
	VIR_CONNECT_LIST_STORAGE_POOLS_NETFS        = C.VIR_CONNECT_LIST_STORAGE_POOLS_NETFS
	VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL      = C.VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL
	VIR_CONNECT_LIST_STORAGE_POOLS_DISK         = C.VIR_CONNECT_LIST_STORAGE_POOLS_DISK
	VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI        = C.VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI
	VIR_CONNECT_LIST_STORAGE_POOLS_SCSI         = C.VIR_CONNECT_LIST_STORAGE_POOLS_SCSI
	VIR_CONNECT_LIST_STORAGE_POOLS_MPATH        = C.VIR_CONNECT_LIST_STORAGE_POOLS_MPATH
	VIR_CONNECT_LIST_STORAGE_POOLS_RBD          = C.VIR_CONNECT_LIST_STORAGE_POOLS_RBD
	VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG     = C.VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG
)

//virDomainDestroyFlagsValues
const (
	VIR_DOMAIN_DESTROY_DEFAULT  = C.VIR_DOMAIN_DESTROY_DEFAULT  //Default behavior - could lead to data loss!!
	VIR_DOMAIN_DESTROY_GRACEFUL = C.VIR_DOMAIN_DESTROY_GRACEFUL //only SIGTERM, no SIGKILL
)

type Hypervisor struct {
	cptr C.virConnectPtr
}

type NodeInfo struct {
	Model   string //string indicating the CPU model
	Memory  uint64 //memory size in kilobytes
	Cpus    uint   //the number of active CPUs
	Mhz     uint   //expected CPU frequency
	Nodes   uint   //the number of NUMA cell, 1 for unusual NUMA topologies or uniform memory access; check capabilities XML for the actual NUMA topology
	Sockets uint   //number of CPU sockets per node if nodes > 1, 1 in case of unusual NUMA topology
	Cores   uint   //number of cores per socket, total number of processors in case of unusual NUMA topology
	Threads uint   //number of threads per core, 1 in case of unusual numa topology
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
	domainType := C.CString(dtype)
	defer C.free(unsafe.Pointer(domainType))

	result := C.virConnectGetMaxVcpus(h.cptr, domainType)
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
func (h *Hypervisor) ListDomains(flags uint) ([]*Domain, error) {
	var cdomains *C.virDomainPtr
	result := C.virConnectListAllDomains(h.cptr, &cdomains, C.uint(flags))
	if result == -1 {
		return nil, GetLastError()
	}

	var domains = make([]*Domain, result)
	p := (*[1 << 30]C.virDomainPtr)(unsafe.Pointer(cdomains))

	for i := 0; i < int(result); i++ {
		domains[i] = newDomain(p[i])
	}
	defer C.free(unsafe.Pointer(cdomains))

	return domains, nil
}

func (h *Hypervisor) GetDefinedDomains() ([]string, error) {
	number := C.virConnectNumOfDefinedDomains(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virConnectListDefinedDomains(h.cptr, &cnames[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) GetActiveDomains() ([]int, error) {
	number := C.virConnectNumOfDomains(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	ids := make([]int, number)

	if number == 0 {
		return ids, nil
	}

	cids := make([]C.int, number)
	result := C.virConnectListDomains(h.cptr, &cids[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cids {
		ids[i] = int(v)
	}

	return ids, nil
}

func (h *Hypervisor) ListInterfaces(flags uint) ([]*Interface, error) {
	var cinterfaces *C.virInterfacePtr
	result := C.virConnectListAllInterfaces(h.cptr, &cinterfaces, C.uint(flags))
	if result == -1 {
		return nil, GetLastError()
	}

	var interfaces = make([]*Interface, result)
	p := (*[1 << 30]C.virInterfacePtr)(unsafe.Pointer(cinterfaces))

	for i := 0; i < int(result); i++ {
		interfaces[i] = newInterface(p[i])
	}
	defer C.free(unsafe.Pointer(cinterfaces))

	return interfaces, nil
}

func (h *Hypervisor) GetDefinedInterfaces() ([]string, error) {
	number := C.virConnectNumOfDefinedInterfaces(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virConnectListDefinedInterfaces(h.cptr, &cnames[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) GetActiveInterfaces() ([]string, error) {
	number := C.virConnectNumOfInterfaces(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virConnectListInterfaces(h.cptr, &cnames[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) ListNetworks(flags uint) ([]*Network, error) {
	var cnetworks *C.virNetworkPtr
	result := C.virConnectListAllNetworks(h.cptr, &cnetworks, C.uint(flags))
	if result == -1 {
		return nil, GetLastError()
	}

	var networks = make([]*Network, result)
	p := (*[1 << 30]C.virNetworkPtr)(unsafe.Pointer(cnetworks))

	for i := 0; i < int(result); i++ {
		networks[i] = newNetwork(p[i])
	}
	defer C.free(unsafe.Pointer(cnetworks))

	return networks, nil
}

func (h *Hypervisor) GetDefinedNetworks() ([]string, error) {
	number := C.virConnectNumOfDefinedNetworks(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virConnectListDefinedNetworks(h.cptr, &cnames[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) GetActiveNetworks() ([]string, error) {
	number := C.virConnectNumOfNetworks(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virConnectListNetworks(h.cptr, &cnames[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) GetNetworkFilters() ([]string, error) {
	number := C.virConnectNumOfNWFilters(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virConnectListNWFilters(h.cptr, &cnames[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) ListNetworkFilters(flags uint) ([]*NetworkFilter, error) {
	var cfilters *C.virNWFilterPtr
	result := C.virConnectListAllNWFilters(h.cptr, &cfilters, C.uint(flags))
	if result == -1 {
		return nil, GetLastError()
	}

	var filters = make([]*NetworkFilter, result)
	p := (*[1 << 30]C.virNWFilterPtr)(unsafe.Pointer(cfilters))

	for i := 0; i < int(result); i++ {
		filters[i] = newNetworkFilter(p[i])
	}
	defer C.free(unsafe.Pointer(cfilters))

	return filters, nil
}

func (h *Hypervisor) GetSecrets() ([][]byte, error) {
	number := C.virConnectNumOfSecrets(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	uuids := make([][]byte, number)

	if number == 0 {
		return uuids, nil
	}

	cuuids := make([]*C.char, number)
	result := C.virConnectListSecrets(h.cptr, &cuuids[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cuuids {
		uuids[i] = C.GoBytes(unsafe.Pointer(v), 16)
		defer C.free(unsafe.Pointer(v))
	}

	return uuids, nil
}

func (h *Hypervisor) GetDefinedStoragePools() ([]string, error) {
	number := C.virConnectNumOfDefinedStoragePools(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virConnectListDefinedStoragePools(h.cptr, &cnames[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) GetActiveStoragePools() ([]string, error) {
	number := C.virConnectNumOfStoragePools(h.cptr)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virConnectListStoragePools(h.cptr, &cnames[0], number)
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) ListStoragePools(flags uint) ([]*StoragePool, error) {
	var cpools *C.virStoragePoolPtr
	result := C.virConnectListAllStoragePools(h.cptr, &cpools, C.uint(flags))
	if result == -1 {
		return nil, GetLastError()
	}

	var pools = make([]*StoragePool, result)
	p := (*[1 << 30]C.virStoragePoolPtr)(unsafe.Pointer(cpools))

	for i := 0; i < int(result); i++ {
		pools[i] = newStoragePool(p[i])
	}
	defer C.free(unsafe.Pointer(cpools))

	return pools, nil
}

//virConnectNumOf functions
func (h *Hypervisor) GetNumberOfDefinedDomains() (int, error) {
	number := C.virConnectNumOfDefinedDomains(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) GetNumberOfActiveDomains() (int, error) {
	number := C.virConnectNumOfDomains(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) GetNumberOfDefinedInterfaces() (int, error) {
	number := C.virConnectNumOfDefinedInterfaces(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) GetNumberOfDefinedNetworks() (int, error) {
	number := C.virConnectNumOfDefinedNetworks(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) GetNumberOfDefinedStoragePools() (int, error) {
	number := C.virConnectNumOfDefinedStoragePools(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) GetNumberOfActiveInterfaces() (int, error) {
	number := C.virConnectNumOfInterfaces(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) GetNumberOfActiveNetworks() (int, error) {
	number := C.virConnectNumOfNetworks(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) GetNumberOfNetworkFilters() (int, error) {
	number := C.virConnectNumOfNWFilters(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) GetNumberOfSecrets() (int, error) {
	number := C.virConnectNumOfSecrets(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

func (h *Hypervisor) ListSecrets(flags uint) ([]*Secret, error) {
	var csecrets *C.virSecretPtr
	result := C.virConnectListAllSecrets(h.cptr, &csecrets, C.uint(flags))
	if result == -1 {
		return nil, GetLastError()
	}

	var secrets = make([]*Secret, result)
	p := (*[1 << 30]C.virSecretPtr)(unsafe.Pointer(csecrets))

	for i := 0; i < int(result); i++ {
		secrets[i] = newSecret(p[i])
	}
	defer C.free(unsafe.Pointer(csecrets))

	return secrets, nil
}

func (h *Hypervisor) GetNumberOfActiveStoragePools() (int, error) {
	number := C.virConnectNumOfStoragePools(h.cptr)
	if number == -1 {
		return 0, GetLastError()
	}

	return int(number), nil
}

//Node functions
func (h *Hypervisor) GetNodeFreeMemoryInNumaCells(startCell int, maxCells int) ([]uint64, error) {
	if startCell < 0 || maxCells <= 0 || (startCell+maxCells) > 10000 {
		return nil, errors.New("GetNodeFreeMemoryInNumaCells: Inconsistent cell bounds")
	}

	freeMemory := make([]C.ulonglong, maxCells)

	result := C.virNodeGetCellsFreeMemory(h.cptr, &freeMemory[0], C.int(startCell), C.int(maxCells))
	if result == -1 {
		return nil, GetLastError()
	}

	cells := make([]uint64, result)
	if result == 0 {
		return cells, nil
	}

	for i := 0; i < int(result); i++ {
		cells[i] = uint64(freeMemory[i])
	}

	return cells, nil
}

func (h *Hypervisor) GetNodeFreeMemory() (uint64, error) {
	result := C.virNodeGetFreeMemory(h.cptr)
	if result == 0 {
		return 0, GetLastError()
	}

	return uint64(result), nil
}

func (h *Hypervisor) GetNodeInfo() (NodeInfo, error) {
	var cNodeInfo C.virNodeInfo
	result := C.virNodeGetInfo(h.cptr, &cNodeInfo)
	if result == -1 {
		return NodeInfo{}, GetLastError()
	}

	return NodeInfo{
		Model:   C.GoString(&cNodeInfo.model[0]),
		Memory:  uint64(cNodeInfo.memory),
		Cpus:    uint(cNodeInfo.cpus),
		Mhz:     uint(cNodeInfo.mhz),
		Nodes:   uint(cNodeInfo.nodes),
		Sockets: uint(cNodeInfo.sockets),
		Cores:   uint(cNodeInfo.cores),
		Threads: uint(cNodeInfo.threads),
	}, nil
}

func (h *Hypervisor) GetNodeDevicesNames(capability string, flags uint) ([]string, error) {
	var cCapability *C.char

	if capability == "" {
		cCapability = C.CString(capability)
		defer C.free(unsafe.Pointer(cCapability))
	}

	number := C.virNodeNumOfDevices(h.cptr, cCapability, 0)
	if number == -1 {
		return nil, GetLastError()
	}

	names := make([]string, number)

	if number == 0 {
		return names, nil
	}

	cnames := make([]*C.char, number)
	result := C.virNodeListDevices(h.cptr, cCapability, &cnames[0], number, C.uint(flags))
	if result == -1 {
		return nil, GetLastError()
	}

	for i, v := range cnames {
		names[i] = C.GoString(v)
		defer C.free(unsafe.Pointer(v))
	}

	return names, nil
}

func (h *Hypervisor) ListNodeDevices(flags uint) ([]*NodeDevice, error) {
	var cdevices *C.virNodeDevicePtr
	result := C.virConnectListAllNodeDevices(h.cptr, &cdevices, C.uint(flags))
	if result == -1 {
		return nil, GetLastError()
	}

	var devices = make([]*NodeDevice, result)
	p := (*[1 << 30]C.virNodeDevicePtr)(unsafe.Pointer(cdevices))

	for i := 0; i < int(result); i++ {
		devices[i] = newNodeDevice(p[i])
	}
	defer C.free(unsafe.Pointer(cdevices))

	return devices, nil
}

func (h *Hypervisor) GetNodeSecurityModel() (map[string]string, error) {
	var cSecurityModel C.virSecurityModel
	result := C.virNodeGetSecurityModel(h.cptr, &cSecurityModel)
	if result == -1 {
		return nil, GetLastError()
	}

	return map[string]string{
		"model": C.GoString(&cSecurityModel.model[0]),
		"doi":   C.GoString(&cSecurityModel.doi[0]),
	}, nil
}

//Event functions
func (h *Hypervisor) RegisterDomainEvent()   {}
func (h *Hypervisor) UnregisterDomainEvent() {}

//Misc functions
func (h *Hypervisor) FindStoragePoolSources() {}

/**
 * Some functions that make more sense
 * being executed from the Hypervisor perspective
 * and using factory pattern
 */

//Domain factories
func (h *Hypervisor) CreateDomain(xml string, flags uint) (*Domain, error) {
	cxml := C.CString(xml)
	defer C.free(unsafe.Pointer(cxml))

	cdomain := C.virDomainCreateXML(h.cptr, cxml, C.uint(flags))
	if cdomain == nil {
		return nil, GetLastError()
	}

	domain := &Domain{cdomain}
	runtime.SetFinalizer(domain, cleanupDomain)

	return domain, nil
}

func (h *Hypervisor) DestroyDomain(domain *Domain, flags uint) error {
	result := C.virDomainDestroyFlags(domain.cptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}

	domain = nil

	return nil
}

func (h *Hypervisor) DefineDomain(xml string) (*Domain, error) {
	cxml := C.CString(xml)
	defer C.free(unsafe.Pointer(cxml))

	cdomain := C.virDomainDefineXML(h.cptr, cxml)
	if cdomain == nil {
		return nil, GetLastError()
	}

	domain := &Domain{cdomain}
	runtime.SetFinalizer(domain, cleanupDomain)

	return domain, nil
}

func (h *Hypervisor) RestoreDomain(filepath string) error {
	cfilepath := C.CString(filepath)
	defer C.free(unsafe.Pointer(cfilepath))

	result := C.virDomainRestore(h.cptr, cfilepath)
	if result == -1 {
		return GetLastError()
	}

	return nil
}

func (h *Hypervisor) LookupDomainById(id int) (*Domain, error) {
	cdomain := C.virDomainLookupByID(h.cptr, C.int(id))
	if cdomain == nil {
		return nil, GetLastError()
	}

	domain := &Domain{cdomain}
	runtime.SetFinalizer(domain, cleanupDomain)

	return domain, nil
}

func (h *Hypervisor) LookupDomainByName(name string) (*Domain, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cdomain := C.virDomainLookupByName(h.cptr, cname)
	if cdomain == nil {
		return nil, GetLastError()
	}

	domain := &Domain{cdomain}
	runtime.SetFinalizer(domain, cleanupDomain)

	return domain, nil
}

func (h *Hypervisor) LookupDomainByUUID(uuid string) (*Domain, error) {
	cuuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cuuid))

	cdomain := C.virDomainLookupByUUIDString(h.cptr, cuuid)
	if cdomain == nil {
		return nil, GetLastError()
	}

	domain := &Domain{cdomain}
	runtime.SetFinalizer(domain, cleanupDomain)

	return domain, nil
}

//NodeDevice
func (h *Hypervisor) CreateNodeDevice()       {}
func (h *Hypervisor) LookupNodeDeviceByName() {}

//Network
func (h *Hypervisor) CreateNetwork()       {}
func (h *Hypervisor) DefineNetwork()       {}
func (h *Hypervisor) LookupNetworkByName() {}
func (h *Hypervisor) LookupNetworkByUUID() {}

//NetworkFilter
func (h *Hypervisor) DefineNetworkFilter()       {}
func (h *Hypervisor) LookupNetworkFilterByName() {}
func (h *Hypervisor) LookupNetworkFilterByUUID() {}

//Interface
func (h *Hypervisor) DefineInterface()             {}
func (h *Hypervisor) LookupInterfaceByName()       {}
func (h *Hypervisor) LookupInterfaceByMacAddress() {}

//Secret
func (h *Hypervisor) DefineSecret()        {}
func (h *Hypervisor) LookupSecretByUsage() {}
func (h *Hypervisor) LookupSecretByUUID()  {}

//StoragePool
func (h *Hypervisor) CreateStoragePool()         {}
func (h *Hypervisor) DefineStoragePool()         {}
func (h *Hypervisor) LookupStoragePoolByName()   {}
func (h *Hypervisor) LookupStoragePoolByUUID()   {}
func (h *Hypervisor) LookupStoragePoolByVolume() {}

//StorageVolume
func (h *Hypervisor) LookupStorageVolumeByKey()  {}
func (h *Hypervisor) LookupStorageVolumeByPath() {}
