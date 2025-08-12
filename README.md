# wait-for

A small command-line tool that waits for a service to be ready before exiting. This is useful in environments like `docker-compose` where one container needs to wait for another to be ready before starting.

## Why it is useful?

When working with microservices or containerized applications, it's a common scenario to have services that depend on each other. For instance, an application container might need to wait for a database container to be fully initialized before it can start. `wait-for` provides a simple and efficient way to handle this by pausing execution until the dependent service is available.

## Installation

You can download the latest pre-compiled binaries for your system from the [GitHub Releases page](https://github.com/fabiolb/wait-for/releases).

For example, to download the `v1.0.0` release for Linux amd64, you could use `curl`:

```bash
curl -L -o wait-for https://github.com/fabiolb/wait-for/releases/download/v1.0.0/wait-for_linux_amd64
chmod +x ./wait-for
```

## Usage

`wait-for` supports waiting for TCP and PostgreSQL services.

### TCP

The `tcp` command waits for a TCP port to be open on a given host.

```bash
wait-for tcp <host>:<port>
```

**Example:**

```bash
$ wait-for tcp localhost:8080
checking tcp connection: localhost:8080
localhost:8080 is down
localhost:8080 is down
localhost:8080 is up
```

### PostgreSQL

The `postgres` command waits for a PostgreSQL server to be ready to accept connections.

```bash
wait-for postgres "<connection-string>"
```

**Example:**

```bash
$ wait-for postgres "postgres://user:password@host:port/dbname?sslmode=disable"
checking postgres server: postgres://user:password@host:port/dbname?sslmode=disable
Error pinging database: dial tcp 127.0.0.1:5432: connect: connection refused
postgres://user:password@host:port/dbname?sslmode=disable is down
Error pinging database: dial tcp 127.0.0.1:5432: connect: connection refused
postgres://user:password@host:port/dbname?sslmode=disable is down
Postgres is up
postgres://user:password@host:port/dbname?sslmode=disable is up
```

## Building from source

To build `wait-for` from source, you'll need to have Go installed (version 1.22 or later).

1.  Clone the repository:
    ```bash
    git clone https://github.com/fabiolb/wait-for.git
    cd wait-for
    ```

2.  Build the binary:
    ```bash
    go build -o wait-for .
    ```

3.  Run the tool:
    ```bash
    ./wait-for --help
    ```

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).
