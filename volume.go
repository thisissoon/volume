package volume

// #cgo LDFLAGS: -lasound -lstdc++
// #include "volume.h"
import "C"

import "errors"

func SetVolume(vol int) {
	C.setVolume(C.int(vol))
}

func GetVolume() (vol int, err error) {
	vol = int(C.getVolume())
	if vol < 0 {
		err = errors.New("Unable to get volume.")
	}
	return
}
