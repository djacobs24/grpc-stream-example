package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/djacobs24/grpc-stream-example/model"
	"google.golang.org/grpc"
)

const port = ":12345"

func main() {
	rand.Seed(time.Now().Unix())

	// Dial the server
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client: Failed to connect to the server: %v", err)
	}

	// Create the stream
	client := model.NewMathClient(conn)
	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("Client: Failed to create stream: %v", err)
	}

	// The stream's context
	strCtx := stream.Context()

	// Holds the max value found
	var max int32

	// Make a new channel that can send and receive a boolean value
	doneChan := make(chan bool)

	// Send a random number to stream
	// Closes after 10 iterations
	go func() {
		for i := 1; i <= 10; i++ {
			// Generate a random number
			num := int32(rand.Intn(i))

			// Build the request
			req := model.NumberRequest{Number: num}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("Client: Failed to send: %v", err)
			}
			log.Printf("Client: %d sent", req.Number)
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// Receives data from stream
	// Saves result in max variable
	// Closes channel if stream is finished
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(doneChan)
				return
			}
			if err != nil {
				log.Fatalf("Client: Failed to receive: %v", err)
			}
			max = resp.Number
			log.Printf("Client: New max %d received", max)
		}
	}()

	// Closes the done channel if context is done
	go func() {
		// Receive from stream context channel
		<-strCtx.Done()
		if err := strCtx.Err(); err != nil {
			log.Printf("Client: %v", err)
		}

		// Close the channel
		close(doneChan)
	}()

	// Receive from done channel
	<-doneChan

	// Print the final max
	log.Printf("Client: Finished with max of %d!", max)

}
