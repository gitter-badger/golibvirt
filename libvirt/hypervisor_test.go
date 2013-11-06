package libvirt

import (
	"os"
	"testing"
)

func TestNewHypervisor(t *testing.T) {
	_, err := NewHypervisor("test:///default")
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestCloseConnection(t *testing.T) {
	h, err := NewHypervisor("test:///default")
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

	_, err = h.GetCapabilities()
	if err != nil {
		virtError, _ := err.(*LibvirtError)
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", virtError, nil)
	}
}

func TestGetHostname(t *testing.T) {
	h, err := NewHypervisor("test:///default")

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

	_, err = h.GetSysInfo(0)

	want := "this function is not supported by the connection driver: virConnectGetSysinfo"

	virtError, _ := err.(*LibvirtError)
	if virtError.Message != want {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", virtError.Message, want)
	}
}

func TestGetType(t *testing.T) {
	h, err := NewHypervisor("test:///default")
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
	_, err = h.GetLibVirtVersion()

	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}
}

func TestGetMaxVcpus(t *testing.T) {
	h, err := NewHypervisor("test:///default")
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
	domains, _ := h.ListDomains(0)
	length := len(domains)
	if length != 1 {
		t.Errorf("incorrect result\ngot:  %d\nwant: %d", length, 1)
	}
}
