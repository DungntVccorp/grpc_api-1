#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"
protoc -I=. --go_out=. api.proto
protoc -I=. --go_out=. authen.api.proto
protoc -I=. --go_out=. connect.api.proto
protoc -I=. --go_out=. internal.api.proto
protoc -I=. --go-grpc_out=. grpc.api.proto
