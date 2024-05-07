package main

import (
	"io"
	"log"

	pb "github.com/i101dev/template-gRPC/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {

	var messages []string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{
				Messages: messages,
			})
		}

		if err != nil {
			return err
		}

		log.Printf("*** >>> [@server] stream.Recv - name - %v", req.Name)

		messages = append(messages, req.Name)
	}
}
