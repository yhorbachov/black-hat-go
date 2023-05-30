package main

import (
	"fmt"

	portscanner "github.com/yhorbachov/black-hat-go/pkg/port_scanner"
)

func main() {
	openports := portscanner.Scan("scanme.nmap.org", 1, 1024)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
