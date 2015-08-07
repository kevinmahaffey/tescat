package main

import (
	"flag"
	"fmt"
	"github.com/kevinmahaffey/tescat"
)

var port = flag.Int("p", 20100, "Destination port for UDP broadcasts")
var fname = flag.String("f", "", "Filename to read pcap from")

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: \n")
		flag.PrintDefaults()
		return
	}
	c, err := tescat.NewCapture()
	if err != nil {
		panic(err)
	} else {
		if *fname != "" {
			err = c.StartFromPCAP(*fname, *port)
			if err != nil {
				panic(err)
			}
		} else {
			err = c.StartFromUDP(*port)
			if err != nil {
				panic(err)
			}
		}
	}
}
