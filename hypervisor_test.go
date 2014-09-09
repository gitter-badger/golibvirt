// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package libvirt

import (
	"os"
	"reflect"
	"testing"
)

func TestNewHypervisor(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestCloseConnection(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = h.CloseConnection()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = h.CloseConnection()
	if err == nil {
		t.Errorf("Hypervisor connection should be "+
			"closed and an error should be returned. Got: %#v\n", err)
	}
}

//Only supported by Qemu: http://libvirt.org/hvsupport.html
// func TestGetBaselineCPU(t *testing.T) {
// 	h, err := NewHypervisor("test:///default")

// 	if err != nil {
// 		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
// 	}

// 	want := `
// <cpu match='exact'>
//     <model>qemu32</model>
//     <feature policy='require' name='xtpr'/>
//     <feature policy='require' name='tm2'/>
//     <feature policy='require' name='est'/>
//     <feature policy='require' name='vmx'/>
//     <feature policy='require' name='monitor'/>
//     <feature policy='require' name='pbe'/>
//     <feature policy='require' name='tm'/>
//     <feature policy='require' name='ht'/>
//     <feature policy='require' name='ss'/>
//     <feature policy='require' name='acpi'/>
//     <feature policy='require' name='ds'/>
//     <feature policy='require' name='clflush'/>
//     <feature policy='require' name='mca'/>
//     <feature policy='require' name='mtrr'/>
//     <feature policy='require' name='vme'/>
// </cpu>
// `

// 	baseline, err := h.GetBaselineCPU([]string{
// 		`<cpu>
// 		    <arch>i686</arch>
// 		    <model>coreduo</model>
// 		    <topology sockets='1' cores='2' threads='1'/>
// 		    <feature name='xtpr'/>
// 		    <feature name='tm2'/>
// 		    <feature name='est'/>
// 		    <feature name='vmx'/>
// 		    <feature name='pbe'/>
// 		    <feature name='tm'/>
// 		    <feature name='ht'/>
// 		    <feature name='ss'/>
// 		    <feature name='acpi'/>
// 		    <feature name='ds'/>
// 		</cpu>`,
// 		`<cpu>
// 		    <arch>i686</arch>
// 		    <model>pentium3</model>
// 		    <topology sockets='1' cores='2' threads='1'/>
// 		    <feature name='lahf_lm'/>
// 		    <feature name='lm'/>
// 		    <feature name='xtpr'/>
// 		    <feature name='cx16'/>
// 		    <feature name='ssse3'/>
// 		    <feature name='tm2'/>
// 		    <feature name='est'/>
// 		    <feature name='vmx'/>
// 		    <feature name='ds_cpl'/>
// 		    <feature name='monitor'/>
// 		    <feature name='pni'/>
// 		    <feature name='pbe'/>
// 		    <feature name='tm'/>
// 		    <feature name='ht'/>
// 		    <feature name='ss'/>
// 		    <feature name='sse2'/>
// 		    <feature name='acpi'/>
// 		    <feature name='ds'/>
// 		    <feature name='clflush'/>
// 		    <feature name='apic'/>
// 		</cpu>
// `}, 0)

// 	if err != nil {
// 		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, want)
// 	}

// 	if baseline != want {
// 		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", baseline, want)
// 	}
// }

//TODO: Add flags to run tests based on the drivers
//present in the local libvirt installation
func TestCompareCPU(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()
	result, err := h.CompareCPU(`<cpu match='exact'>
	<model>qemu32</model>
	<feature policy='require' name='xtpr'/>
	<feature policy='require' name='tm2'/>
	<feature policy='require' name='est'/>
	<feature policy='require' name='vmx'/>
	<feature policy='require' name='monitor'/>
	<feature policy='require' name='pbe'/>
	<feature policy='require' name='tm'/>
	<feature policy='require' name='ht'/>
	<feature policy='require' name='ss'/>
	<feature policy='require' name='acpi'/>
	<feature policy='require' name='ds'/>
	<feature policy='require' name='clflush'/>
	<feature policy='require' name='mca'/>
	<feature policy='require' name='mtrr'/>
	<feature policy='require' name='vme'/>
</cpu>`, 0)

	if result != VIR_CPU_COMPARE_ERROR {
		t.Errorf("incorrect result\ngot:  %#d\nwant: %#d", result, VIR_CPU_COMPARE_ERROR)
	}

	want := "this function is not supported by the connection driver: virConnectCompareCPU"

	virtError, _ := err.(*LibvirtError)
	if virtError.Message != want {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", virtError.Message, want)
	}
}

func TestGetCapabilities(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	_, err = h.GetCapabilities()
	if err != nil {
		virtError, _ := err.(*LibvirtError)
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", virtError, nil)
	}
}

func TestGetHostname(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	hostname, err := h.GetHostname()
	if err != nil {
		virtError, _ := err.(*LibvirtError)
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", virtError, nil)
	}

	osHostname, _ := os.Hostname()

	if hostname != osHostname {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", hostname, osHostname)
	}
}

func TestGetSysInfo(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	_, err = h.GetSysInfo(0)

	want := "this function is not supported by the connection driver: virConnectGetSysinfo"

	virtError, _ := err.(*LibvirtError)
	if virtError.Message != want {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", virtError.Message, want)
	}
}

func TestGetType(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()
	htype, err := h.GetType()
	want := "Test"
	if htype != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", htype, want)
	}

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestGetConnectionUri(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()
	uri, err := h.GetConnectionUri()

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := "test:///default"
	if uri != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", uri, want)
	}
}

func TestGetVersion(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	version, err := h.GetVersion()

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if version["rel"] != 2 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", version["rel"], 2)
	}
}

func TestGetLibVirtVersion(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	_, err = h.GetLibVirtVersion()

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestGetMaxVcpus(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	vcpus, err := h.GetMaxVcpus("")

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if vcpus != 32 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", vcpus, 32)
	}
}

func TestIsConnectionEncrypted(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	encrypted, err := h.IsConnectionEncrypted()

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if encrypted {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", encrypted, false)
	}
}

func TestIsConnectionSecure(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	secure, err := h.IsConnectionSecure()

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if !secure {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", secure, true)
	}
}

func TestIsConnectionAlive(t *testing.T) {
	h, err := NewHypervisor("test:///default")
	defer h.CloseConnection()

	alive, err := h.IsConnectionAlive()

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if !alive {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", alive, true)
	}
}

func TestListDomains(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domains, _ := h.ListDomains(0)
	length := len(domains)
	if length != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 1)
	}
}

func TestGetDefinedDomains(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	dnames, _ := h.GetDefinedDomains()
	length := len(dnames)
	if length > 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetActiveDomains(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	ids, _ := h.GetActiveDomains()
	length := len(ids)
	if length != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 1)
	}

	if ids[0] != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", ids[0], 1)
	}
}

func TestListInterfaces(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	interfaces, _ := h.ListInterfaces(0)
	length := len(interfaces)
	if length != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetDefinedInterfaces(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	dnames, _ := h.GetDefinedInterfaces()
	length := len(dnames)
	if length > 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetActiveInterfaces(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	names, _ := h.GetActiveInterfaces()
	length := len(names)
	if length != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 1)
	}
	want := "eth1"

	if names[0] != want {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", names[0], want)
	}
}

func TestListNetworks(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	networks, _ := h.ListNetworks(0)
	length := len(networks)
	if length != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 1)
	}
}

func TestGetDefinedNetworks(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	dnames, _ := h.GetDefinedNetworks()
	length := len(dnames)
	if length > 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetActiveNetworks(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	names, _ := h.GetActiveNetworks()
	length := len(names)
	if length != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 1)
	}
	want := "default"

	if names[0] != want {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", names[0], want)
	}
}

func TestGetNetworkFilters(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	names, _ := h.GetNetworkFilters()
	length := len(names)
	if length > 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestListNetworkFilters(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	filters, _ := h.ListNetworkFilters(0)
	length := len(filters)
	if length != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetSecrets(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	secrets, _ := h.GetSecrets()
	length := len(secrets)
	if length > 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetDefinedStoragePools(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	names, _ := h.GetDefinedStoragePools()
	length := len(names)
	if length > 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetActivateStoragePools(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	names, _ := h.GetActiveStoragePools()
	length := len(names)
	if length != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 1)
	}
	want := "default-pool"

	if names[0] != want {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", names[0], want)
	}
}

func TestListStoragePools(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	pools, _ := h.ListStoragePools(0)
	length := len(pools)
	if length != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 1)
	}
}

func TestGetNumberOfDefinedDomains(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfDefinedDomains()
	if n != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 0)
	}
}

func TestGetNumberOfActiveDomains(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfActiveDomains()
	if n != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 1)
	}
}

func TestGetNumberOfDefinedInterfaces(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfDefinedInterfaces()
	if n != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 0)
	}
}

func TestGetNumberOfDefinedNetworks(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfDefinedNetworks()
	if n != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 0)
	}
}

func TestGetNumberOfDefinedStoragePools(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfDefinedStoragePools()
	if n != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 0)
	}
}

func TestGetNumberOfActiveInterfaces(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfActiveInterfaces()
	if n != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 1)
	}
}

func TestGetNumberOfActiveNetworks(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfActiveNetworks()
	if n != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 1)
	}
}

func TestGetNumberOfNetworkFilters(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfNetworkFilters()
	if n != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 1)
	}
}

func TestGetNumberOfSecrets(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfSecrets()
	if n != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 1)
	}
}

func TestListSecrets(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	secrets, _ := h.ListSecrets(0)
	length := len(secrets)
	if length != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetNumberOfActiveStoragePools(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	n, _ := h.GetNumberOfActiveStoragePools()
	if n != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", n, 1)
	}
}

func TestGetNodeFreeMemoryInNumaCells(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	cells, _ := h.GetNodeFreeMemoryInNumaCells(0, 30)
	length := len(cells)
	if length != 2 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 2)
	}

	cells, err := h.GetNodeFreeMemoryInNumaCells(0, 0)
	want := "GetNodeFreeMemoryInNumaCells: Inconsistent cell bounds"
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err, want)
	}
}

func TestGetNodeFreeMemory(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	memory, _ := h.GetNodeFreeMemory()
	if memory != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", memory, 0)
	}
}

func TestGetNodeInfo(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	nodeInfo, _ := h.GetNodeInfo()

	want := NodeInfo{"i686", 3145728, 16, 1400, 2, 2, 2, 2}

	if !reflect.DeepEqual(nodeInfo, want) {
		t.Errorf("Incorrect result\ngot:  %#v\nwant: %#v", nodeInfo, want)
	}
}

func TestGetNodeDevicesNames(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	names, _ := h.GetNodeDevicesNames("", 0)
	length := len(names)
	if length != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestListNodeDevices(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	devices, _ := h.ListNodeDevices(0)
	length := len(devices)
	if length != 0 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 0)
	}
}

func TestGetNodeSecurityModel(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	secModel, err := h.GetNodeSecurityModel()

	want := "this function is not supported by the connection driver: virNodeGetSecurityModel"
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err, want)
	}

	if secModel["model"] != "" {
		t.Errorf("Incorrect result\ngot:  %#v\nwant: %#v", secModel["model"], nil)
	}

	if secModel["doi"] != "" {
		t.Errorf("Incorrect result\ngot:  %#v\nwant: %#v", secModel["doi"], nil)
	}
}
