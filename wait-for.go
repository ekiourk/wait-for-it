package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkService(conn string) bool {
	timeout := time.Duration(1000 * time.Millisecond)
	_, err := net.DialTimeout("tcp", conn, timeout)
	if err != nil {
		return false
	}
	return true
}

func main() {
	connStr := os.Args[1]

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

