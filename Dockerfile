FROM golang:1.5-alpine

RUN apk add --no-cache git
COPY . /go/src/app
RUN go get -d -v app
RUN go install -v app

CMD ["/go/src/app/run.sh"]
