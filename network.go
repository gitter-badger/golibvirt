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

type Network struct {
	cptr C.virNetworkPtr
}

func cleanupNetwork(network *Network) {
	C.virNetworkFree(network.cptr)
}

func newNetwork(cptr C.virNetworkPtr) *Network {
	network := &Network{cptr}
	runtime.SetFinalizer(network, cleanupNetwork)
	return network
}

func (n *Network) GetName() (string, error) {
	result := C.virNetworkGetName(n.cptr)
	if result == nil {
		return "", GetLastError()
	}

	name := C.GoString(result)
	return name, nil
}
