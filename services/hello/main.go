package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"qthang02/hello-rpc/pb/proto/hello"
)

type Server struct {
	hello.UnimplementedGreeterServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	hello.RegisterGreeterServer(s, &Server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8001")
	log.Fatal(s.Serve(lis))
}
