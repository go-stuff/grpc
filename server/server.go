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
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	// init database client
	client, ctx, err := initMongoClient()
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}
	defer client.Disconnect(ctx)

	// use a default mongo database if the MONGO_DATABASE environment variable is not set
	if os.Getenv("MONGO_DATABASE") == "" {
		os.Setenv("MONGO_DATABASE", "test")
	}

	// init a pointer to the database
	db := client.Database(os.Getenv("MONGO_DATABASE"))

	// use a default api uri if the API_URI environment variable is not set
	if os.Getenv("API_URI") == "" {
		os.Setenv("API_URI", "localhost:6000")
	}

	lis, err := net.Listen("tcp", os.Getenv("API_URI"))
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	// with cert
	creds, err := credentials.NewServerTLSFromFile("./certs/cert.pem", "./certs/key.pem")
	if err != nil {
		log.Fatalf("failed to get creds: %v", err)
	}
	svr := grpc.NewServer(grpc.Creds(creds))

	// without cert
	//svr := grpc.NewServer()

	// Register services with the server
	api.RegisterSessionServiceServer(svr, &service.SessionServiceServer{DB: db})
	api.RegisterPermissionServiceServer(svr, &service.PermissionServiceServer{DB: db})
	api.RegisterUserServiceServer(svr, &service.UserServiceServer{DB: db})
	api.RegisterRoleServiceServer(svr, &service.RoleServiceServer{DB: db})

	// Register reflection service on gRPC server.
	reflection.Register(svr)

	fmt.Printf("listening @ %v...", lis.Addr().String())
	svr.Serve(lis)
}

func initMongoClient() (*mongo.Client, context.Context, error) {
	// a Context carries a deadline, cancelation signal, and request-scoped values
	// across API boundaries. Its methods are safe for simultaneous use by multiple
	// goroutines
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// use a default mongo uri if the MONGO_URI environment variable is not set
	if os.Getenv("MONGO_URI") == "" {
		os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	}

	// connect does not do server discovery, use ping
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
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
