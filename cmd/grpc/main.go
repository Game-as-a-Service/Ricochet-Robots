package main

import (
	"fmt"
	"ricochetrobots/server/grpc"
)

func main() {
	process := make(chan struct{})
	fmt.Println("starting gRPC server...")
	server := grpc.New()
	server.Run("50051")

	<-process
}
