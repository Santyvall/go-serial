package main

import (
	"log"

	"github.com/mikepb/go-serial"
)

func main() {
	options := serial.RawOptions
	options.BitRate = 115200
	p, err := options.Open("/dev/ttyUSB0")
	if err != nil {
		log.Panic(err)
	}

	defer p.Close()

	buf := make([]byte, 5)
	if _, err := p.Read(buf); err != nil {
		log.Panic(err)
	} else {
		log.Println(buf)
	}
}
