package main

import (
	"context"
	"io"
	"log"

	pb "vyshnavi/grpc-go/proto"

)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Server Started");

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Could not send names: %v", err);
	}
	
	for {
		message, err := stream.Recv();
		
		if err == io.EOF {
			break;
		}
		if err != nil {
			log.Fatalf("Error occured while streaming %v", err);
		}
		log.Println(message);
	}
	log.Println("Streaming Done");
}