package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/urfave/cli"
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

func run(checkService func(conn string) bool, connStr string) {
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

func main() {

	app := cli.NewApp()
	app.Usage = "Given a service type and connection parameter it will wait till service is running"
	app.Version = "0.1.0-beta"
	app.Commands = []cli.Command{
		{
			Name:    "tcp",
			Aliases: []string{"t"},
			Usage:   "check a tcp connection",
			Action: func(c *cli.Context) error {
				fmt.Println("checking tcp connection: ", c.Args().First())
				run(checkTCPService, c.Args().First())
				return nil
			},
		},
		{
			Name:    "postgres",
			Aliases: []string{"p"},
			Usage:   "check a postgres database",
			Action: func(c *cli.Context) error {
				fmt.Println("checking postgres server: ", c.Args().First())
				run(checkPGService, c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
