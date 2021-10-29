package main

import (
	"log"
	"net"

	"grpcTutorial/chat"

	"google.golang.org/grpc"
)




func main() {
	port := "9000"
	lis, err := net.Listen("tcp",":"+port)

	if err != nil {
		log.Fatal(err);
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatal("Failed to serve gRPC server over port " + port + ": " , err)
	}



}