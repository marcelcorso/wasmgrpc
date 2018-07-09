package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	httpListenAddress = flag.String("http-listen-internal-address", ":5001", "The private address to listen for http")
	grpcListenAddress = flag.String("grpc-listen-address", ":50001", "The address to listen for grpc")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "yo " + req.Name,
	}, nil
}

func main() {
	flag.Parse()

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	go setupHTTP()

	lis, err := net.Listen("tcp", *grpcListenAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening grpc at " + *grpcListenAddress)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func setupHTTP() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	fs := http.FileServer(http.Dir("static"))
	// http.Handle("/", http.StripPrefix("/static", fs))
	http.Handle("/", fs)

	log.Println("listening http at " + *httpListenAddress)

	log.Fatal(http.ListenAndServe(*httpListenAddress, nil))
}
