package main

import (
	"context"
	"log"
	"time"

	pb "github.com/i101dev/template-gRPC/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for _, name := range names.Names {

		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}

		log.Printf("Successfully sent the request with name: %s", name)

		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error @ [stream.CloseAndRecv]")
	}

	log.Printf("Final result [res.messages] - %+v", res.Messages)

	log.Printf("Client streaming ended")
}
