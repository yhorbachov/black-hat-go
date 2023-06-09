package portscanner

import (
	"fmt"
	"net"
	"sort"
)

func Scan(host string, from, to int) []int {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := from; i < cap(ports); i++ {
		go worker(host, ports, results)
	}

	go func() {
		for i := from; i <= to; i++ {
			ports <- i
		}
	}()

	for i := from; i <= to; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)

	return openports
}

func worker(host string, ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- port
	}
}
