package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	help  bool
	v, V  bool
	h     string
	p     int
	types string
	c     int
)

func init() {
	flag.BoolVar(&help, "help", false, "Show gaping help")
	flag.BoolVar(&v, "v", false, "Show Version and exit")
	flag.BoolVar(&V, "V", false, "Show Version and exit")
	flag.StringVar(&h, "h", "", "Dest host ipaddress")
	flag.IntVar(&p, "p", 0, "Dest host port")
	flag.StringVar(&types, "types", "tcp", "Set network protocol")
	flag.IntVar(&c, "c", 100000000, "Tcp test conuts")
}

// version
func Version() {
	if v || V {
		fmt.Println("gaping  v0.01")
		os.Exit(0)
	}
}

// help
func Help() {
	fmt.Println(`Gaping  v0.02- Copyright (c) 2019 Mike chulinx
Example: pping -h 127.0.0.1,localhost -p 80 -type tcp`)
	flag.PrintDefaults()
	os.Exit(0)
}

// usage
func Usages() {
	if help {
		Help()
	}
}

func main() {
	flag.Parse()
	Version()
	Usages()
	if h == "" || p == 0 {
		Help()
	}
	hosts := strings.Split(h, ",")
	protocol := SetProtocol(types)
	if len(hosts) > 1 {

		for i := 0; i < c; i++ {
			for _, i := range hosts {
				network := NewNetwork(i, p, protocol)
				network.PingPrint()
			}
		}

	} else {
		network := NewNetwork(h, p, protocol)
		for i := 0; i < c; i++ {
			network.PingPrint()
		}
	}
}
