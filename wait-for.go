package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkTCPService(conn string) bool {
	timeout := time.Duration(1000 * time.Millisecond)
	_, err := net.DialTimeout("tcp", conn, timeout)
	if err != nil {
		return false
	}
	return true
}

func main() {
	serviceType := os.Args[1]
	connStr := os.Args[2]

	var checkService func(conn string) bool
	if serviceType == "tcp" {
		checkService = checkTCPService
	} else {
		fmt.Println("No valid service name passed.")
		os.Exit(1)
	}

	ticker := time.Tick(1000 * time.Millisecond)
	for {
		if checkService(connStr) {
			fmt.Println(connStr, "is up")
			os.Exit(0)
		}
		fmt.Println(connStr, "is down")
		<-ticker
	}
}
