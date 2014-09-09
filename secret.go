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

type Secret struct {
	cptr C.virSecretPtr
}

func cleanupSecret(secret *Secret) {
	C.virSecretFree(secret.cptr)
}

func newSecret(cptr C.virSecretPtr) *Secret {
	secret := &Secret{cptr}
	runtime.SetFinalizer(secret, cleanupSecret)
	return secret
}
