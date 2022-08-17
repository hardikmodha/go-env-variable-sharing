#!/bin/bash

# Build child
cd child
go build -o child

cd ../

# Run the interceptor
go run interceptor/main.go > interceptor_output.txt 2>&1 &

interceptor_pid=$!

# Run the parent
go run parent/main.go


kill $interceptor_pid