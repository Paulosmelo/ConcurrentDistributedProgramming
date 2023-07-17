#!/bin/bash
go run ../../UDP/server.go

cust_func(){
 go run ../../UDP/client.go & 
}

cust_func

wait 
echo "All done"