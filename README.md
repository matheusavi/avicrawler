# How to run it?
`go run .\cmd\api\main.go`


## How to create migrations?
You'll need [golang-migrate](https://github.com/golang-migrate/) installed:

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Then you can execute the following command to create a migration:

```
migrate create -ext sql -dir cmd/migrate/migrations -seq my_migration_name  
```


## How to debug the app?
Execute:
```
go install -v github.com/go-delve/delve/cmd/dlv@latest
```
Then hit F5 at visual studio code.


# How to build the docker image?
```
docker build -t avicrawler
```

## How to run the app through docker?

```
docker run --publish 3457:3457 --name avicrawler avicrawler
```