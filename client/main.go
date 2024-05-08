package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/i101dev/template-gRPC/proto"
)

const (
	port = ":5000"
)

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to diala the server %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	clientNames := &pb.NamesList{
		Names: []string{"Sampson", "Frank", "Rick"},
	}

	call_SayHello(client)
	call_SayHello_ServerStr(client, clientNames)
	call_SayHello_ClientStr(client, clientNames)
	call_SayHello_BidirStr(client, clientNames)
}
