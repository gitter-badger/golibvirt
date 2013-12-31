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
