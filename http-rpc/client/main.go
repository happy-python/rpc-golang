package main

import (
	"log"
	"net/rpc"
	. "rpc-golang/http-rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	args := Args{
		A: 2,
		B: 3,
	}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Arith error:", err)
	}
	log.Printf("%d * %d = %d\n", args.A, args.B, reply)
}
