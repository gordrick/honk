package main

import (
	"github.com/gordrick/honk/pkg/pb/honk"
	"github.com/gordrick/honk/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	url := os.Getenv("URL")

	lis, err := net.Listen("tcp", url)

	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	s := services.HonkService{}

	grpcServer := grpc.NewServer()

	honk.RegisterHonkServiceServer(grpcServer, &s)

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
