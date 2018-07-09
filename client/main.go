package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {

	address := flag.String("address", ":50001", "The address to listen for grpc")

	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{
		Name: "marcel",
	}

	resp, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("got a reply: " + resp.Message)
}
