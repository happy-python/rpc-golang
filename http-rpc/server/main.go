package main

import (
	"log"
	"net/http"
	"net/rpc"
	. "rpc-golang/http-rpc"
)

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	// 把该服务注册到 HTTP 协议上
	rpc.HandleHTTP()
	log.Println("HTTP RPC service listen and serving on port 1234")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
