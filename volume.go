// Package volume lets you control the volume of ALSA sound outputs.
// This package specifically targets the Raspberry Pi.
package volume

// #cgo LDFLAGS: -lasound -lstdc++
// #include <stdlib.h>
// #include "volume.h"
import "C"

import (
	"errors"
	"unsafe"
)

// SetVolume sets the volume of all sound outputs.
func SetVolume(card string, vol int) {
	s := C.CString(card)
	C.setVolume(s, C.int(vol))
	C.free(unsafe.Pointer(s))
}

// GetVolume returns the volume of the first sounds output found.
func GetVolume(card string) (vol int, err error) {
	s := C.CString(card)
	vol = int(C.getVolume(s))
	C.free(unsafe.Pointer(s))
	if vol < 0 {
		err = errors.New("Unable to get volume.")
	}
	return
}
