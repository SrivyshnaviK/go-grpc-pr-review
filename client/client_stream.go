package main

import (
	"context"
	"time"
	"log"

	pb "vyshnavi/grpc-go/proto"

)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Client Streaming Started");

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %v", err);
	}

	for _,name:= range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err:= stream.Send(req); err != nil {
			log.Fatalf("Error occured while streaming %v", err);
		}
		log.Printf("Request sent with name %s", name);
		time.Sleep(2 * time.Second);
	}
	
	log.Println("Streaming Done");

	res, err := stream.CloseAndRecv()

	if err!=nil {
		log.Fatalf("Error occured while receiving %v", err);
	}

	log.Printf("Response %v", res.Messages)
}