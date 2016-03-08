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

## Todo

* Add support for Postgres
* Support database URIs
