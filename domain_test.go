package libvirt

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Printf // For debugging; delete when done.

func TestCreateDomain(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	xml := `
<domain type='test'>
  <name>test1</name>
  <uuid>4dab22b31d52d8f32516782e98ab3fa1</uuid>

  <os>
    <type>hvm</type>
    <boot dev='cdrom'/>
    <boot dev='hd'/>
    <boot dev='network'/>
  </os>

  <memory unit='KiB'>654321</memory>
  <vcpu>1</vcpu>

  <features>
    <pae/>
    <acpi/>
    <apic/>
  </features>
</domain>
`

	_, err := h.CreateDomain(xml, VIR_DOMAIN_NONE)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestDefineDomain(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	xml := `
<domain type='test'>
  <name>test1</name>
  <uuid>4dab22b31d52d8f32516782e98ab3fa1</uuid>

  <os>
    <type>hvm</type>
    <boot dev='cdrom'/>
    <boot dev='hd'/>
    <boot dev='network'/>
  </os>

  <memory unit='KiB'>654321</memory>
  <vcpu>1</vcpu>

  <features>
    <pae/>
    <acpi/>
    <apic/>
  </features>
</domain>
`

	_, err := h.DefineDomain(xml)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestLookupDomainById(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	_, err := h.LookupDomainById(1)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestLookupDomainByName(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	_, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestLookupDomainByUUID(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	_, err := h.LookupDomainByUUID("4dab22b3-1d52-d8f3-2516-782e98ab3fa1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestGetName(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	name, err := domain.GetName()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := "test1"
	if name != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", name, want)
	}
}

func TestGetUUID(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	uuid, err := domain.GetUUID()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := "4dab22b3-1d52-d8f3-2516-782e98ab3fa1"
	if uuid != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", uuid, want)
	}
}

func TestGetId(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	id, err := domain.GetId()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := uint(1)
	if id != want {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", id, want)
	}
}

func TestGetInfo(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	domainInfo, err := domain.GetInfo()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := DomainInfo{VIR_DOMAIN_RUNNING, 654321, 654321, 1, 1383851248042553000}

	//CputTime changes upon every execution so lets set it to the same
	//value returned by GetInfo() in order to make the test pass
	want.CpuTime = domainInfo.CpuTime

	if !reflect.DeepEqual(domainInfo, want) {
		t.Errorf("Incorrect result\ngot:  %v\nwant: %v", domainInfo, want)
	}
}

func TestGetAutostart(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	autostart, err := domain.GetAutostart()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if !autostart {
		t.Errorf("incorrect result\ngot:  %t\nwant: %t", autostart, true)
	}
}

func TestSetAutostart(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	autostart, _ := domain.GetAutostart()
	if !autostart {
		t.Errorf("incorrect result\ngot:  %t\nwant: %t", autostart, true)
	}

	err = domain.SetAutostart(false)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	autostart, err = domain.GetAutostart()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if autostart {
		t.Errorf("incorrect result\ngot:  %t\nwant: %t", autostart, false)
	}
}

func TestGetOsType(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := "linux"
	ostype, err := domain.GetOsType()
	if ostype != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", ostype, want)
	}
}

func TestGetMaxMemory(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := uint64(654321)
	memory, err := domain.GetMaxMemory()
	if memory != want {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", memory, want)
	}
}



func TestSetMemory(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := uint64(262144)

	err = domain.SetMemory(want)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	domainInfo, err := domain.GetInfo()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if domainInfo.Memory != want {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", domainInfo.Memory, want)
	}
}

func TestSetMaxMemory(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := uint64(503316)

	err = domain.SetMaxMemory(want)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	memory, err := domain.GetMaxMemory()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if memory != want {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", memory, want)
	}
}

func TestMaxVcpus(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	vcpus, err := domain.GetMaxVcpus()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := 2
	if vcpus != want {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", vcpus, want)
	}
}

func TestIsActive(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	isActive, err := domain.IsActive()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if !isActive {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", isActive, true)
	}
}

func TestIsPersistent(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	isPersistent, err := domain.IsPersistent()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if !isPersistent {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", isPersistent, true)
	}
}

func TestIsUpdated(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	isUpdated, err := domain.IsUpdated()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	if isUpdated {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", isUpdated, false)
	}
}

func TestReboot(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = domain.Reboot(VIR_DOMAIN_SHUTDOWN_DEFAULT)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestReset(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := "this function is not supported by the connection driver: virDomainReset"

	err = domain.Reset()
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err.Error(), want)
	}
}

func TestSave(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = domain.Save("/tmp/golibvirt-test")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err, nil)
	}
	err = h.RestoreDomain("/tmp/golibvirt-test")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %v\nwant: %v", err, nil)
	}
}

func TestRestoreDomain(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
		return
	}

	err = domain.Save("/tmp/golibvirt-test")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = h.RestoreDomain("/tmp/golibvirt-test")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestSuspendResume(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = domain.Suspend()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = domain.Resume()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestShutdownAndStart(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = domain.Shutdown()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	domain, err = h.LookupDomainByName("test1")
	err = domain.Start()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestDestroyDomain(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := "this function is not supported by the connection driver: virDomainDestroyFlags"
	err = h.DestroyDomain(domain, VIR_DOMAIN_DESTROY_GRACEFUL)
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err.Error(), want)
	}
}

func TestSendKeys(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := "this function is not supported by the connection driver: virDomainSendKey"

	err = domain.SendKey(VIR_KEYCODE_SET_LINUX, 150, []uint{12, 2, 3}, 0)
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err.Error(), want)
	}
}

func TestGetVcpus(t *testing.T) {}

func TestSetVcpus(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	err = domain.SetVcpus(1, VIR_DOMAIN_AFFECT_LIVE)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	domainInfo, _ := domain.GetInfo()
	if domainInfo.Vcpus != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", domainInfo.Vcpus, 1)
	}
}

func TestMigrate(t *testing.T)                 {}
func TestSetMigrationMaxDowntime(t *testing.T) {}
func TestPinVcpu(t *testing.T)                 {}

func TestAttachDetachDevice(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	xml := `
<devices>
  <disk type='file' device='cdrom'>
    <source file='/tmp/boot.iso'/>
    <target dev='hdc'/>
    <readonly/>
  </disk>
</devices>
`
	want := "this function is not supported by the connection driver: virDomainAttachDevice"
	err = domain.AttachDevice(xml)
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err.Error(), want)
	}

	want = "this function is not supported by the connection driver: virDomainDetachDevice"
	err = domain.DetachDevice(xml)
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err.Error(), want)
	}
}

func TestUpdateDevice(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	xml := `
<devices>
  <disk type='file' device='cdrom'>
    <source file='/tmp/boot.iso'/>
    <target dev='hdc'/>
  </disk>
</devices>
`
	want := "this function is not supported by the connection driver: virDomainUpdateDeviceFlags"
	err = domain.UpdateDevice(xml, VIR_DOMAIN_DEVICE_MODIFY_LIVE)
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, want)
	}
}

func TestToXml(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	_, err = domain.ToXml(VIR_DOMAIN_XML_SECURE)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestGetState(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	_, _, err = domain.GetState(0)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

//TODO
func TestGetJobInfo(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	want := "this function is not supported by the connection driver: virDomainGetJobInfo"

	_, err = domain.GetJobInfo()
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err.Error(), want)
	}

}

func TestGetJobStats(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
	want := "this function is not supported by the connection driver: virDomainGetJobStats"

	_, err = domain.GetJobStats()
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err.Error(), want)
	}

}

func TestAbortCurrentJob(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
	want := "this function is not supported by the connection driver: virDomainAbortJob"

	err = domain.AbortCurrentJob()
	if err.Error() != want {
		t.Errorf("incorrect result\ngot:  %s\nwant: %s", err.Error(), want)
	}
}
func TestGetSchedType(t *testing.T)       {}
func TestGetSchedParams(t *testing.T)     {}
func TestSetSchedParams(t *testing.T)     {}
func TestGetSecurityLabel(t *testing.T)   {}
func TestSaveManagedImage(t *testing.T)   {}
func TestRemoveManagedImage(t *testing.T) {}
func TestHasManagedImage(t *testing.T)    {}
func TestMemoryPeek(t *testing.T)         {}
func TestGetMemoryStats(t *testing.T)     {}
func TestBlockPeek(t *testing.T)          {}
func TestGetBlockStats(t *testing.T)      {}
func TestGetBlockInfo(t *testing.T)       {}
func TestCoreDump(t *testing.T)           {}
func TestGetInterfaceStats(t *testing.T)  {}

func TestHasCurrentSnapshot(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	hasSnapshot, err := domain.HasCurrentSnapshot(0)

	if hasSnapshot {
		t.Errorf("incorrect result\ngot:  %t\nwant: %t", hasSnapshot, false)
	}
}

//TODO
func TestRevertToSnapshot(t *testing.T)           {}
func TestTakeSnapshot(t *testing.T)               {}
func TestGetCurrentSnapshot(t *testing.T)         {}
func TestDeleteSnapshot(t *testing.T)             {}
func TestLookupSnapshotByName(t *testing.T)       {}
func TestGetSnapshots(t *testing.T)               {}
func TestBlockCommit(t *testing.T)                {}
func TestBlockJobAbort(t *testing.T)              {}
func TestBlockJobSetSpeed(t *testing.T)           {}
func TestBlockPull(t *testing.T)                  {}
func TestBlockRebase(t *testing.T)                {}
func TestBlockResize(t *testing.T)                {}
func TestFSTrim(t *testing.T)                     {}
func TestGetBlockIoTune(t *testing.T)             {}
func TestGetBlockJobInfo(t *testing.T)            {}
func TestGetCpuStats(t *testing.T)                {}
func TestGetControlInfo(t *testing.T)             {}
func TestGetDiskErrors(t *testing.T)              {}
func TestGetEmulatorPinInfo(t *testing.T)         {}
func TestGetDomainHostname(t *testing.T)          {}
func TestGetMetadata(t *testing.T)                {}
func TestSetMetadata(t *testing.T)                {}
func TestInjectNmi(t *testing.T)                  {}
func TestManagedSave(t *testing.T)                {}
func TestManagedSaveRemove(t *testing.T)          {}
func TestMigrateGetCompressionCache(t *testing.T) {}
func TestMigrateSetCompressionCache(t *testing.T) {}
func TestMigrateGetMaxSpeed(t *testing.T)         {}
func TestMigrateSetMaxSpeed(t *testing.T)         {}
func TestOpenChannel(t *testing.T)                {}
func TestOpenConsole(t *testing.T)                {}
func TestOpenGraphics(t *testing.T)               {}
func TestPMSuspendForDuration(t *testing.T)       {}
func TestPMWakUp(t *testing.T)                    {}
func TestPinEmulator(t *testing.T)                {}
func TestScreenshot(t *testing.T)                 {}
func TestSendProcessSignal(t *testing.T)          {}
func TestSetNumaParameters(t *testing.T)          {}

func TestUndefine(t *testing.T) {
	h, _ := NewHypervisor("test:///default")
	defer h.CloseConnection()

	domain, err := h.LookupDomainByName("test1")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

        domain.Shutdown()
	err = domain.Undefine(VIR_DOMAIN_UNDEFINE_NOFLAGS)
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}
