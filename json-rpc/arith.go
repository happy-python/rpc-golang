package json_rpc

import (
	"log"
	"net/http"
)

type Arith int

type Args struct {
	A, B int
}

func (t *Arith) Multiply(r *http.Request, args *Args, result *int) error {
	log.Printf("Multiply %d with %d\n", args.A, args.B)
	*result = args.A * args.B
	return nil
}
