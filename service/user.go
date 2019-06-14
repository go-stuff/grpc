package service

import (
	"context"

	"github.com/go-stuff/grpc/api"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserCollection is the name of the collection in the database.
const UserCollection string = "users"

// UserServiceServer is a service in the API for managing users.
type UserServiceServer struct {
	DB *mongo.Database
}

// Slice returns a slice of users
func (s *UserServiceServer) Slice(ctx context.Context, req *api.UserSliceReq) (*api.UserSliceRes, error) {
	// prepare a Res
	res := new(api.UserSliceRes)

	// find all users
	cursor, err := s.DB.Collection(UserCollection).Find(ctx,
		bson.M{},
		&options.FindOptions{
			Sort: bson.M{
				"username": 1, // acending
				//"username": -1, // descending
			},
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// itterate each document returned
	for cursor.Next(ctx) {
		var user = new(api.User)
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		// append the current role to the slice
		res.Users = append(res.Users, user)
	}

	// handle any errors with the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// Create inserts a user
func (s *UserServiceServer) Create(ctx context.Context, req *api.UserCreateReq) (*api.UserCreateRes, error) {
	// prepare a Res
	res := new(api.UserCreateRes)

	user := &api.User{
		Id:         primitive.NewObjectID().Hex(), // ObjectID's are generated based on time
		Username:   req.User.Username,
		Groups:     req.User.Groups,
		CreatedBy:  req.User.CreatedBy,
		CreatedAt:  ptypes.TimestampNow(),
		ModifiedBy: req.User.ModifiedBy,
		ModifiedAt: ptypes.TimestampNow(),
	}

	// insert role into mongo
	_, err := s.DB.Collection(UserCollection).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// update the Res id
	res.Id = user.Id

	return res, nil
}

// Read returns a single user
func (s *UserServiceServer) Read(ctx context.Context, req *api.UserReadReq) (*api.UserReadRes, error) {
	// prepare a Res
	res := new(api.UserReadRes)

	// initialize user
	res.User = new(api.User)

	// find a user
	err := s.DB.Collection(UserCollection).FindOne(ctx,
		bson.M{
			"id": req.Id,
		},
	).Decode(res.User)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Read returns a single user
func (s *UserServiceServer) ByUsername(ctx context.Context, req *api.UserByUsernameReq) (*api.UserByUsernameRes, error) {
	// prepare a Res
	res := new(api.UserByUsernameRes)

	// initialize user
	res.User = new(api.User)

	// find a user
	err := s.DB.Collection(UserCollection).FindOne(ctx,
		bson.M{
			"username": req.Username,
		},
	).Decode(res.User)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update modifies a user
func (s *UserServiceServer) Update(ctx context.Context, req *api.UserUpdateReq) (*api.UserUpdateRes, error) {
	// prepare a Res
	res := new(api.UserUpdateRes)

	// update a user
	updateRes, err := s.DB.Collection(UserCollection).UpdateOne(ctx,
		bson.M{
			"id": req.User.Id,
		},
		bson.M{
			"$set": bson.M{
				"username":   req.User.Username,
				"groups":     req.User.Groups,
				"modifiedby": req.User.ModifiedBy,
				"modifiedat": ptypes.TimestampNow(),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	res.Updated = updateRes.ModifiedCount

	return res, nil
}

// Delete removes a user
func (s *UserServiceServer) Delete(ctx context.Context, req *api.UserDeleteReq) (*api.UserDeleteRes, error) {
	// prepare a Res
	res := new(api.UserDeleteRes)

	// delete a user
	deleteRes, err := s.DB.Collection(UserCollection).DeleteOne(ctx,
		bson.M{
			"id": req.Id,
		},
	)
	if err != nil {
		return nil, err
	}

	res.Deleted = deleteRes.DeletedCount

	return res, nil
}
