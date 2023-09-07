package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"qthang02/hello-rpc/pb/proto/hello"
)

type HelloRequest struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.POST("/hello", func(c *gin.Context) {
		var data HelloRequest

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		response := ClientRPC(data)

		c.JSON(200, gin.H{"message": response})
	})

	r.Run(":8080")
}

func ClientRPC(data HelloRequest) string {
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &hello.HelloRequest{Name: data.Name})
	if err != nil {
		log.Fatal(err)
	}

	return reply.Message
}
