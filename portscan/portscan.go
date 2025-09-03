package portscan

import (
	"fmt"
	"net"
	"time"
)

func Portscan(host string, ports []int) {
	type result struct {
		port int
		open bool
	}
	results := make(chan result)

	for _, port := range ports {
		go func(p int) {
			address := fmt.Sprintf("%s:%d", host, p)
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err == nil {
				conn.Close()
				results <- result{port: p, open: true}
			} else {
				results <- result{port: p, open: false}
			}
		}(port)
	}

	for range ports {
		r := <-results
		if r.open {
			fmt.Printf("%d: open\n", r.port)
		} else {
			fmt.Printf("%d: closed\n", r.port)
		}
	}
}
