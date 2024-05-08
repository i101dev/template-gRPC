package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/i101dev/template-gRPC/proto"
)

func call_SayHello(client pb.GreetServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("\n*** >>> @client.unary - [client.SayHello] - %+v", err)
	}

	log.Printf("\n*** >>> @client.unary - [client.SayHelloe] - res - %s", res.Message)
}

func call_SayHello_ClientStr(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Client streaming started")

	stream, err := client.SayHello_ClientStr(context.Background())

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
func call_SayHello_ServerStr(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Server streaming started")

	stream, err := client.SayHello_ServerStr(context.Background(), names)

	if err != nil {
		log.Fatalf("could not send names: %+v", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while Server streaming: %+v", err)
		}

		log.Println("*** >>> [@client] stream.Recv - ", message)
	}

	log.Printf("Server streaming finished")
}
func call_SayHello_BidirStr(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Bidirectional streaming has started")

	stream, err := client.SayHello_BidirStr(context.Background())

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
