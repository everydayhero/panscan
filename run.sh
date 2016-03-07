#!/bin/sh

go run $(find . -name '*.go' -not -name '*_test.go') $@
