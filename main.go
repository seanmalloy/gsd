package main

import (
	"fmt"
	"log"
	"net"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func getDNSNames(names []string) {
	for _, n := range names {
		addr, _ := net.LookupHost(n)
		fmt.Println(n, addr[0]) // TODO what to do with multiple results?
	}
}

func testTCPPort(host string, port string) {
	var status string
	timeout := time.Duration(3000000000) // 3 seconds
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		log.Println("Connection error:", err)
		status = "Unreachable"
	} else {
		status = "Online"
		defer conn.Close()
	}
	log.Println(status)
}

var (
	dns  = kingpin.Command("dns", "Get information about hosts from DNS")
	ping = kingpin.Command("ping", "Get information about hosts from DNS")
	tcp  = kingpin.Command("tcp", "Get information about hosts from DNS")
)

func main() {
	switch kingpin.Parse() {
	case "dns":
		println("TODO: do DNS look up")
		names := []string{"www.google.com", "www.openbsd.org"}
		getDNSNames(names)
	case "ping":
		println("TODO: do ping")
	case "tcp":
		println("TODO: do tcp connection")
		testTCPPort("8.8.8.8", "53")
	}
}
