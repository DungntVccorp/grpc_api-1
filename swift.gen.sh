#!/bin/bash
protoc --swift_out=./swift/ --plugin=/opt/homebrew/bin/protoc-gen-grpc-swift --grpc-swift_out=./swift/ *.proto