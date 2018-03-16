# wait-for-it
Small utility that waits for a service to be ready and then exits

# Commands

## tcp
```bash
wait-for tcp localhost:80
```
## postgres
```bash
wait-for postgres "postgres://postgres:postgres@172.18.0.3:5432/postgres?sslmode=disable&connect_timeout=1"
```
