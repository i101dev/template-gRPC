package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/i101dev/template-gRPC/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Bidirectional streaming has started")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("\n*** >>> @client - [bidirectional] - error creating stream: %+v", err)
	}

	wait_ch := make(chan struct{})

	go func() {

		for {

			message, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while streaming: %+v", err)
			}

			log.Println("*** >>> [0x1] - ", message)
		}

		close(wait_ch)
	}()

	for _, name := range names.Names {

		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("\n*** >>> @client - [bidirectional] - error sending request: %+v", err)
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()

	<-wait_ch

	log.Printf("Bidirectional streaming has ended")
}
