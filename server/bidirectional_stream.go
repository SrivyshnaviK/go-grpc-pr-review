package main 

import(
	pb "vyshnavi/grpc-go/proto"
	"log"
	"io"
)

func (s *helloServer)SayBidirectionalStreaming(stream pb.GreetService_SayBidirectionalStreamingServer) error{
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}

}