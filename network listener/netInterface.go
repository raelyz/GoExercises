package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	StartTCPServer()
}

func inface() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range interfaces {
		fmt.Printf("Interfaces:%v\n", i.Name)
		fmt.Printf("MACaddress:%v\n", i.HardwareAddr)

		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("InterfaceAddress#%v:%v\n", k, v.String())
		}
		fmt.Println()
	}
}
