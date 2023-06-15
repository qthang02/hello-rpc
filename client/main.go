package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatalln("Dialing: ", err)
	}

	var reply string

	err = client.Call("HelloService.Hello", "World", &reply)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(reply)
}
