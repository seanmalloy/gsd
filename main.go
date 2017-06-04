package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"time"
	"log"
)

func getDnsNames(names []string) {
	for _, n := range names {
		addr, _ := net.LookupHost(n)
		fmt.Println(n, addr[0]) // TODO what to do with multiple results?
	}
}

// TODO requires root
func pingHosts(name string) {
	p := fastping.NewPinger()
	//p.Network("udp")
	ra, err := net.ResolveIPAddr("ip4:icmp", name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func testTcpPort(host string, port string) {
	var status string
	timeout := time.Duration(3000000000) // 3 seconds
	conn, err := net.DialTimeout("tcp", host + ":" + port, timeout)
	if err != nil {
		log.Println("Connection error:", err)
		status = "Unreachable"
	} else {
		status = "Online"
		defer conn.Close()
	}
	log.Println(status)
}

func main() {
	names := []string{"www.google.com", "www.openbsd.org"}
	getDnsNames(names)

	pingHosts("8.8.8.8")

	testTcpPort("8.8.8.8", "53")
	testTcpPort("8.8.8.8", "22")
}
