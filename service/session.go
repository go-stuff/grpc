package service

import (
	"context"

	"github.com/go-stuff/grpc/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SessionCollection is the name of the collection in the database.
const SessionCollection string = "sessions"

// SessionServiceServer is a service in the API for managing Sessions.
type SessionServiceServer struct {
	DB *mongo.Database
}

// List returns a slice of Sessions
func (s *SessionServiceServer) List(ctx context.Context, req *api.SessionListReq) (*api.SessionListRes, error) {
	// prepare a Res
	res := new(api.SessionListRes)

	// find all Sessions
	cursor, err := s.DB.Collection(SessionCollection).Find(ctx,
		bson.M{},
		&options.FindOptions{
			Sort: bson.M{
				"username": 1, // acending
				//"username": -1, // descending
			},
		},
	)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// itterate each document returned
	for cursor.Next(ctx) {
		var session = new(api.Session)
		err := cursor.Decode(&session)
		if err != nil {
			return nil, err
		}

		// append the current role to the slice
		res.Sessions = append(res.Sessions, session)
	}

	// handle any errors with the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
