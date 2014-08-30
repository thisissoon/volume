// Package volume lets you control the volume of ALSA sound outputs.
// This package specifically targets the Raspberry Pi.
package volume

// #cgo LDFLAGS: -lasound -lstdc++
// #include "volume.h"
import "C"

import "errors"

// SetVolume sets the volume of all sound outputs.
func SetVolume(vol int) {
	C.setVolume(C.int(vol))
}

// GetVolume returns the volume of the first sounds output found.
func GetVolume() (vol int, err error) {
	vol = int(C.getVolume())
	if vol < 0 {
		err = errors.New("Unable to get volume.")
	}
	return
}
