package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing", err)
	}
	args := [5]Args{
		{7, 8},
		{2, 3},
		{4, 5},
		{9, 10},
		{1, 6},
	}
	for _,arg := range args {
		var reply int
		err = client.Call("Arith.Multiply", arg, &reply)
		if err != nil {
			log.Fatal("arith error", err)
		}
		fmt.Printf("Arith : %d * %d = %d \n", arg.A, arg.B, reply)
	}
}
