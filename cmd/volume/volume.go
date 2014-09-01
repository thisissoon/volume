// The volume command lets you get or set the volume of ALSA sound outputs.
// This command is especially targeted for use on a Raspberry Pi.
package main

import (
	"flag"
	"fmt"
	"github.com/bklimt/volume"
)

func main() {
	vol := flag.Int("volume", -1, "The volume to set.")
	card := flag.String("card", "default", "The card to set the volume for.")
	flag.Parse()

	if *vol != -1 {
		volume.SetVolume(*card, *vol)
	}

	if v, err := volume.GetVolume(*card); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Volume: %d%%\n", v)
	}
}
