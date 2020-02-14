// RPCClient
package main

import (
	"log"
	"net/rpc"
)

// Args is arguments of rpc call
type Args struct {
	Id string
}

func main() {
	var reply int64
	args := Args{"1234"}

	client, err := rpc.DialHTTP("tcp", "localhost:1234/rpc")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	err = client.Call("JSONServer.GiveBookDetail", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("%v", reply)
}

// How to fixed it with codec?
