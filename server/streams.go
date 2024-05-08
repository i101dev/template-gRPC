package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/i101dev/template-gRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	res := &pb.HelloResponse{
		Message: "@server.unary - SayHello - success",
	}

	return res, nil
}

func (s *helloServer) SayHello_ClientStr(stream pb.GreetService_SayHello_ClientStrServer) error {

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

func (s *helloServer) SayHello_ServerStr(req *pb.NamesList, stream pb.GreetService_SayHello_ServerStrServer) error {

	log.Printf("Got request with names: %+v", req.Names)

	for _, name := range req.Names {

		res := &pb.HelloResponse{
			Message: name + " says Word",
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}

func (s *helloServer) SayHello_BidirStr(stream pb.GreetService_SayHello_BidirStrServer) error {

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