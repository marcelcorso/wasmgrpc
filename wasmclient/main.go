package main

import (
	"context"
	"log"
	"syscall/js"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	done   = make(chan struct{})
	client pb.GreeterClient
)

func main() {
	address := "localhost:50001"
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Println("error: did not connect: %v", err)
	}

	defer conn.Close()
	client = pb.NewGreeterClient(conn)

	callback := js.NewCallback(sayHello)
	defer callback.Close() // To defer the callback releasing is a good practice (on master its `callback.Release()`)
	setSayHello := js.Global().Get("setSayHello")
	setSayHello.Invoke(callback)

	// run forever
	<-done
}

func sayHello(args []js.Value) {
	name := args[0].String()

	req := &pb.HelloRequest{
		Name: name,
	}

	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	js.Global().Get("sayHelloReply").Invoke(resp.Message)
}
