#!/bin/bash
set -eo pipefail

# Kill the ports we are going to use, in case we need it
# lsof -ti :3000 | xargs kill -9
# lsof -ti :3001 | xargs kill -9

protoc -I ./pb --go_out=plugins=grpc:./pb ./pb/*.proto

go build -o api/api api/main.go
go build -o gcd/gcd gcd/main.go

./gcd/gcd &
./api/api -target="localhost" -port="3001" &

# curl localhost:3001/gcd/294/44
