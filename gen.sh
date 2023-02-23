#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"
protoc -I=. --go_out=. *.proto
protoc -I=. --go-grpc_out=. grpc.api.proto
