package libvirt

import (
	"fmt"
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

  <memory>654321</memory>
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

// func TestGetName(t *testing.T) {
// 	h, _ := NewHypervisor("test:///default")
// 	defer h.CloseConnection()

// 	fmt.Printf("%s\n", "blah")
// }
