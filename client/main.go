package main

import (
	"log"

	pb "vyshnavi/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn);

	names := &pb.NamesList{
		Names: []string{"Vyshnavi", "Shravan", "Nikhil", "AI", "ChatGPT"},
	}

	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	// callSayHelloClientStreaming(client, names)
	callSayHelloBidirectionalStream(client, names)
}