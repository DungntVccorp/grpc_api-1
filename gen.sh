#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"
protoc -I=. --go_out=. api.proto
protoc -I=. --go_out=. connect.api.proto
protoc -I=. --go-grpc_out=. grpc.api.proto
