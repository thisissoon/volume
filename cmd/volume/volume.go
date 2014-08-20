package main

import (
	"flag"
	"fmt"
	"github.com/bklimt/volume"
)

func main() {
	vol := flag.Int("volume", -1, "The volume to set.")
	flag.Parse()

	if *vol != -1 {
		volume.SetVolume(*vol)
	}

	if v, err := volume.GetVolume(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Volume: %d%%\n", v)
	}
}
