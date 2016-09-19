package main

import (
	"fmt"
	"net"
)

func main() {
	var name string = "www.google.com"
	fmt.Println(net.LookupHost(name))
}
