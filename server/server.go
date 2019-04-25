package main

import (
	"io"
	"log"
	"net"

	"github.com/djacobs24/grpc-stream-example/model"
	"google.golang.org/grpc"
)

// The port number being used
const port = ":12345"

type server struct{}

func (s server) Max(srv model.Math_MaxServer) error {

	log.Println("Server: Starting new server")

	// The service's contexts
	svcCtx := srv.Context()

	// Holds the max value found
	var max int32

	for {
		// Exit if context is done
		select {
		// Receive from service context channel
		case <-svcCtx.Done():
			return svcCtx.Err()
		default:
		}

		// Receive data from the client
		req, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				// Return will close the stream from the server side
				log.Println("Server: Exit")
				return nil
			}
			log.Printf("Server: Receiving error: %v", err)
			continue
		}

		// Continue if number received from stream is less than max
		if req.Number <= max {
			continue
		}

		// Update the max
		max = req.Number

		// Send it to the stream
		resp := model.NumberResponse{Result: max}
		if err := srv.Send(&resp); err != nil {
			log.Printf("Server: Sending error: %v", err)
		}

		log.Printf("Server: Sending new max: %d", max)
	}
}

func main() {
	// Create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Server: Failed to create listener: %v", err)
	}

	// Create GRPC server
	s := grpc.NewServer()
	model.RegisterMathServer(s, server{})

	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server: Failed to serve: %v", err)
	}
}
