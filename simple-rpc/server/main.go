package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	
	*reply = args.A * args.B
	fmt.Printf("ding dong got the request %d * %d = %d \n",args.A,args.B,*reply)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	listener, err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal(err)
	}
	rpc.Accept(listener)
}