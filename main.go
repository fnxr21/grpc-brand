package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fnxr21/brand/config"
	"github.com/fnxr21/brand/internal/handler"
	"github.com/fnxr21/brand/internal/repository"
	"github.com/fnxr21/brand/protobuf/golang_protobuf_brand"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// gateway server going to reverse-proxy all incoming requests to gRPC sibling server (AppConfig.RPCPort is equal “6001”)
// which is run as separate go routine (see my previous chapter for details).

//create a traditional HTTP server, passing custom Mux. This server listens to port AppConfig.Port,
//which is equal “6000” and previously was used by gorilla mux, and going to be decommissioned.

// start new HTTP server, which will be transcoded and reverse-proxy to gRPC server.

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	config.InitDB()

	// push RPC server as goroutine
	go StartRPCServer()

	// push gRPC-Gateway generated server as goroutine
	go StartRPCGatewayServer()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	log.Println("Termination signal received. Exiting...")

}

func StartRPCGatewayServer() {
	gwmux := runtime.NewServeMux()
	IPPORT := os.Getenv("IP_PORT")
	GRPCPORT := os.Getenv("GRPC_PORT")

	err := golang_protobuf_brand.RegisterCrudHandlerFromEndpoint(context.Background(), gwmux, ":"+GRPCPORT, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	gwServer := &http.Server{
		Addr:    ":" + IPPORT,
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://localhost:%s\n", IPPORT)
	log.Fatalln(gwServer.ListenAndServe())
}

func StartRPCServer() {
	GRPCPORT := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPCPORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userRepository := repository.Repositorybrand(config.DB)

	h := handler.HandlerBrand(userRepository)
	
	golang_protobuf_brand.RegisterCrudServer(s, h)

	log.Printf("gRPC server listening on port %v\n", GRPCPORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
