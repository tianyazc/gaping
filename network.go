package main

import (
	"fmt"
	"github.com/labstack/gommon/color"
	"net"
	"time"
)

type Network struct {
	Host      string
	Port      int
	Protocols ProtocolStruct
}

type ProtocolStruct struct {
	protocol string
}

type ModParam func(ps *ProtocolStruct)

func SetProtocol(p string) ModParam {
	return func(ps *ProtocolStruct) {
		ps.protocol = p
	}
}

func NewNetwork(host string, port int, param ...ModParam) *Network {
	option := ProtocolStruct{
		protocol: "tcp",
	}
	for _, fn := range param {
		fn(&option)
	}
	return &Network{
		Host:      host,
		Port:      port,
		Protocols: option,
	}
}

func (n *Network) Connect() (error, time.Duration) {
	address := fmt.Sprintf("%s:%d", n.Host, n.Port)
	nets := new(net.Dialer)
	nets.Timeout = time.Millisecond * 3000
	startTime := time.Now()
	conn, err := nets.Dial(n.Protocols.protocol, address)
	if err != nil {
		return err, 1000000
	}
	conn.Write([]byte("hello I'm test!"))
	times := time.Since(startTime).Round(time.Microsecond)
	defer conn.Close()
	return nil, times
}

func (n *Network) PingPrint() {
	date := time.Now().Format("2006-01-02 15:04:05")
	err, t := n.Connect()
	if err != nil {
		fmt.Printf("%s %s %s", date, color.Blue(n.Host), color.Red("Connection timed out\n"))
	} else {
		fmt.Printf("%s %s Connected to  %s:%s: time=%s protocol=%s port=%s\n",
			date, color.Blue(n.Host), color.Green(n.Host), color.Green(n.Port), color.Green(t), color.Green(n.Protocols.protocol), color.Green(n.Port))
	}
	time.Sleep(1000 * time.Millisecond)
}
