package main

import (
	"flag"
	"fmt"

	portscanner "github.com/yhorbachov/black-hat-go/pkg/port_scanner"
)

func main() {
	var from int
	var to int
	var host string

	flag.StringVar(&host, "host", "", "target host")
	flag.IntVar(&from, "from", 1, "beggining of port scanning range (>0)")
	flag.IntVar(&to, "to", 65535, "end of port scanning range (<65535)")

	flag.Parse()

	if host == "" {
		fmt.Println("Invalid host")
		return
	}

	if from < 1 {
		fmt.Println("'from' argument has to be bigger than zero")
		return
	}

	if to > 65535 {
		fmt.Println("'to' argument has to be smaller rhan 65535")
		return
	}

	openports := portscanner.Scan(host, from, to)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
