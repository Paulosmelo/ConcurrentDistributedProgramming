#!/bin/bash
#go run ../../TCP/server.go

go run client.go "teste 1 " &
go run client.go "teste 2 " &
go run client.go "teste 3 " &