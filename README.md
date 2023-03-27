- go to `http://localhost:8989`
- login with
    - server: `mysql-master`
    - username: `root`
    - password: `12345678`
- create database `ks-rewards`

```shell
go run ./cmd/app migration -up 0
go run ./cmd/app api
```

When you need to run workers
```shell
go run ./cmd/app worker --master
go run ./cmd/app worker
```

Whenever you need to migrate db
```shell
migrate create -ext sql -dir migrations/mysql <name_of_migration>
```
