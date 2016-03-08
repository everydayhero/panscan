FROM golang:1.5-alpine

RUN apk add --no-cache git
COPY docker-entrypoint.sh /entrypoint.sh
COPY . /go/src/app
RUN go get -d -v app
RUN go install -v app
RUN go build -o bin/panscan app

ENTRYPOINT ["/entrypoint.sh"]
CMD [""]
