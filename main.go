package main

import (
    pb "github.com/JuanKPaniagua/goTest/booksapp"
    "google.golang.org/grpc"
    "log"
    "net"
)

func main() {
    port := "8080"
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterBookInfoServer(s, &server{})

    log.Printf("Starting gRPC listener on port " + port)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}