package main

import (
	"fmt"
	"log"
	"net"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func getDNSNames(dnsNames *[]string) {
	for _, name := range *dnsNames {
		addr, _ := net.LookupHost(name)

		// TODO: do not use fmt here, instead return a data structure
		fmt.Println(name, addr[0])
	}

}

// TODO: create function to print output

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
	dns          = kingpin.Command("dns", "Get information about hosts from DNS")
	dnsHostNames = dns.Arg("hostname", "hostname to lookup in DNS").Required().Strings()

	ping = kingpin.Command("ping", "Get information about hosts from ICMP(ping)")

	tcp = kingpin.Command("tcp", "Get information about a TCP port for hosts")
)

func main() {
	switch kingpin.Parse() {
	case "dns":
		getDNSNames(dnsHostNames)
	case "ping":
		println("TODO: do ping")
	case "tcp":
		println("TODO: do tcp connection")
		testTCPPort("8.8.8.8", "53")
	}
}
