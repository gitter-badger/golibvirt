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

type NetworkFilter struct {
	cptr C.virNWFilterPtr
}

func cleanupNetworkFilter(nfilter *NetworkFilter) {
	C.virNWFilterFree(nfilter.cptr)
}

func newNetworkFilter(cptr C.virNWFilterPtr) *NetworkFilter {
	nfilter := &NetworkFilter{cptr}
	runtime.SetFinalizer(nfilter, cleanupNetworkFilter)
	return nfilter
}

func (n *NetworkFilter) GetName() (string, error) {
	result := C.virNWFilterGetName(n.cptr)
	if result == nil {
		return "", GetLastError()
	}

	name := C.GoString(result)
	return name, nil
}
