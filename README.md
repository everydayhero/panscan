# PANscan

Scans through databases for PANs. Pretty simple.

## Usage

You can use as a standalone binary:

```
panscan mysql root@tcp(localhost:3306)/
```

Or via a Docker image:

```
docker run --rm -it panscan mysql root@tcp(localhost:3306)/
```

### Parameters

* `-d name[,name]` - Ignore databases
* `-t name[,name]` - Ignore tables

## Building

Make sure you have Go 1.5 or higher installed and make sure the source is in your GOPATH.

```
go get -d -v panscan
go install -v panscan
go build -o bin/panscan panscan
```

You can also build the Docker container using `docker-compose`:

```
docker-compose build
```


## Testing
The tests are ran in containers against a specific database engine to ensure compatibility.

To run tests locally ensure an instance of MySQL is running and has `fixtures/db.sql` loaded.

Be sure to have the following environment variables setup:
* `DB_DRIVER` - Specifies the database driver to use.
* `DB_CONN` - The connection options to connect to the database.

```
DB_DRIVER=mysql DB_CONN="root@tcp(localhost:3306)/" go test -v panscan
```

You can also run tests using Docker Compose which is configured to run tests against a specific database engine.

```
docker-compose run test_mysql
```

## Todo

* Add support for Postgres
* Support database URIs
