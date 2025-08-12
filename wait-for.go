package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

func checkTCPService(conn string) bool {
	timeout := time.Second
	_, err := net.DialTimeout("tcp", conn, timeout)
	return err == nil
}

func checkPGService(conn string) bool {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println("Error opening database connection:", err)
		return false
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Println("Error pinging database:", err)
		return false
	}

	log.Println("Postgres is up")
	return true
}

func run(ctx context.Context, checkService func(conn string) bool, connStr string) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if checkService(connStr) {
				log.Println(connStr, "is up")
				return nil
			}
			log.Println(connStr, "is down")
		}
	}
}

func main() {
	app := &cli.App{
		Name:  "wait-for",
		Usage: "Given a service type and connection parameter it will wait till service is running",
		Commands: []*cli.Command{
			{
				Name:    "tcp",
				Aliases: []string{"t"},
				Usage:   "check a tcp connection",
				Action: func(c *cli.Context) error {
					connStr := c.Args().First()
					log.Println("checking tcp connection:", connStr)
					return run(c.Context, checkTCPService, connStr)
				},
			},
			{
				Name:    "postgres",
				Aliases: []string{"p"},
				Usage:   "check a postgres database",
				Action: func(c *cli.Context) error {
					connStr := c.Args().First()
					log.Println("checking postgres server:", connStr)
					return run(c.Context, checkPGService, connStr)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
