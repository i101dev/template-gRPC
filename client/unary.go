package main

import (
	"context"
	"log"
	"time"

	pb "github.com/i101dev/template-gRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("\n*** >>> @client.unary - [client.SayHello] - %+v", err)
	}

	log.Printf("\n*** >>> @client.unary - [client.SayHelloe] - res - %s", res.Message)
}
