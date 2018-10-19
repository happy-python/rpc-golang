package main

import (
	"bytes"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	. "rpc-golang/json-rpc"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func main() {
	url := "http://localhost:1234/rpc"
	args := Args{
		A: 2,
		B: 3,
	}

	message, err := json.EncodeClientRequest("Arith.Multiply", args)
	checkError(err)

	resp, err := http.Post(url, "application/json", bytes.NewReader(message))
	defer resp.Body.Close()

	checkError(err)

	reply := new(int)
	err = json.DecodeClientResponse(resp.Body, reply)
	checkError(err)

	log.Printf("%d * %d = %d\n", args.A, args.B, *reply)
}
