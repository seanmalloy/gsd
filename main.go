package main

import (
	"fmt"
	"net"
)

func get_dns_names(names []string) {
	for _, n := range names {
		addr, _ := net.LookupHost(n)
		fmt.Println(n, addr[0]) // TODO what to do with multiple results?
	}
}

func main() {
	names := []string{"www.google.com", "www.openbsd.org"}
	get_dns_names(names)
}
