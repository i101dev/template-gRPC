package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/i101dev/template-gRPC/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("*** >>> [bidirectional] Server got request with name: %v", req.Name)

		res := &pb.HelloResponse{
			Message: "- What's happening" + req.Name,
		}

		if err := stream.Send(res); err != nil {
			return fmt.Errorf("\n*** >>> [bidirectional] error sending response stream: %+v", err)
		}
	}
}
