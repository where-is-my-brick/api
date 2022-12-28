package main

import (
	"github.com/where-is-my-brick/api/grpc/category_service"
	"github.com/where-is-my-brick/api/grpc/image_service"
	"github.com/where-is-my-brick/api/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

const serverPort = ":50051"
const pathPrefix = "/tmp/images"

func main() {
	// initialize the server
	var grpcServer = grpc.NewServer()

	// register the services
	image_service.RegisterImageServiceServer(grpcServer, &services.ImageServiceServer{PathPrefix: pathPrefix})
	category_service.RegisterCategoryServiceServer(grpcServer, &services.CategoryServiceServer{PathPrefix: pathPrefix})

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// start the server
	grpcServer.Serve(lis)
}
