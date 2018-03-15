package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func checkTCPService(conn string) bool {
	timeout := time.Duration(1000 * time.Millisecond)
	_, err := net.DialTimeout("tcp", conn, timeout)
	if err != nil {
		return false
	}
	return true
}

func checkPGService(conn string) bool {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println(err)
		return false
	}
	rows, queryErr := db.Query("select version()")
	if queryErr != nil {
		fmt.Println(queryErr)
		return false
	}
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			fmt.Println(err)
			return false
		}
		fmt.Printf("version is %s\n", version)
	}
	return true
}

func main() {
	serviceType := os.Args[1]
	connStr := os.Args[2]

	var checkService func(conn string) bool
	if serviceType == "tcp" {
		checkService = checkTCPService
	} else if serviceType == "postgres" {
		checkService = checkPGService
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
