package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatalln("ListenTCP error: ", err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalln("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}
