Volume
======

A go library for controlling the volume of output on a Raspberry Pi.

# Dependencies

To install the dependencies in raspbian, install these apt packages:

    sudo apt-get install libasound2-dev libasound2

# Usage

    go install github.com/bklimt/volume/...
    $GOPATH/bin/volume 75
