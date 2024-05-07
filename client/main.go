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

	// serverNames := &pb.NamesList{
	// 	Names: []string{"Ricardo", "Jimmy", "Money"},
	// }
	clientNames := &pb.NamesList{
		Names: []string{"Sampson", "Frank", "Rick"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client, serverNames)
	// callSayHelloClientStream(client, clientNames)
	callSayHelloBidirectionalStream(client, clientNames)
}
