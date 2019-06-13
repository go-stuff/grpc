package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/go-stuff/grpc/api"
	"github.com/go-stuff/grpc/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// init database client
	client, ctx, err := initMongoClient()
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}
	defer client.Disconnect(ctx)

	// init a pointer to the database
	db := client.Database("test")

	lis, err := net.Listen("tcp", "localhost:6000")
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()

	// Register services with the server
	api.RegisterUserServiceServer(server, &service.UserServiceServer{DB: db})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	fmt.Printf("listening @ %v...", lis.Addr().String())
	server.Serve(lis)
}

func initMongoClient() (*mongo.Client, context.Context, error) {
	// a Context carries a deadline, cancelation signal, and request-scoped values
	// across API boundaries. Its methods are safe for simultaneous use by multiple
	// goroutines
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// use a default mongo uri if the MONGOURL environment variable is not set
	if os.Getenv("MONGOURL") == "" {
		os.Setenv("MONGOURL", "mongodb://localhost:27017")
	}

	// connect does not do server discovery, use ping
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		return nil, nil, err
	}

	// ping for server discovery
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	log.Println("main.go > INFO > Connected to MongoDB @", os.Getenv("MONGOURL"))
	return client, ctx, nil
}
