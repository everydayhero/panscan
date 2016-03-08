#!/bin/sh

CMD=$1
BIN=/go/bin/panscan

case $1 in
  test)
    go test -v app
    ;;
  *)
    exec "$BIN" $@
    ;;
esac
