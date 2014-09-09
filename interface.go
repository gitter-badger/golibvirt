// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
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

type Interface struct {
	cptr C.virInterfacePtr
}

func cleanupInterface(vinterface *Interface) {
	C.virInterfaceFree(vinterface.cptr)
}

func newInterface(cptr C.virInterfacePtr) *Interface {
	vinterface := &Interface{cptr}
	runtime.SetFinalizer(vinterface, cleanupInterface)
	return vinterface
}

func (i *Interface) GetName() (string, error) {
	result := C.virInterfaceGetName(i.cptr)
	if result == nil {
		return "", GetLastError()
	}

	name := C.GoString(result)
	return name, nil
}
