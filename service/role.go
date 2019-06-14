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

// RoleCollection is the name of the collection in the database.
const RoleCollection string = "roles"

// RoleServiceServer is a service in the API for managing Roles.
type RoleServiceServer struct {
	DB *mongo.Database
}

// Slice returns a slice of Roles
func (s *RoleServiceServer) Slice(ctx context.Context, req *api.RoleSliceReq) (*api.RoleSliceRes, error) {
	// prepare a Res
	res := new(api.RoleSliceRes)

	// find all Roles
	cursor, err := s.DB.Collection(RoleCollection).Find(ctx,
		bson.M{},
		&options.FindOptions{
			Sort: bson.M{
				"Name": 1, // acending
				//"Name": -1, // descending
			},
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// itterate each document returned
	for cursor.Next(ctx) {
		var Role = new(api.Role)
		err := cursor.Decode(&Role)
		if err != nil {
			return nil, err
		}

		// append the current role to the slice
		res.Roles = append(res.Roles, Role)
	}

	// handle any errors with the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// Create inserts a Role
func (s *RoleServiceServer) Create(ctx context.Context, req *api.RoleCreateReq) (*api.RoleCreateRes, error) {
	// prepare a Res
	res := new(api.RoleCreateRes)

	Role := &api.Role{
		ID:          primitive.NewObjectID().Hex(), // ObjectID's are generated based on time
		Name:        req.Role.Name,
		Description: req.Role.Description,
		CreatedBy:   req.Role.CreatedBy,
		CreatedAt:   ptypes.TimestampNow(),
		ModifiedBy:  req.Role.ModifiedBy,
		ModifiedAt:  ptypes.TimestampNow(),
	}

	// insert role into mongo
	_, err := s.DB.Collection(RoleCollection).InsertOne(ctx, Role)
	if err != nil {
		return nil, err
	}

	// update the Res id
	res.Id = Role.ID

	return res, nil
}

// Read returns a single Role
func (s *RoleServiceServer) Read(ctx context.Context, req *api.RoleReadReq) (*api.RoleReadRes, error) {
	// prepare a Res
	res := new(api.RoleReadRes)

	// initialize Role
	res.Role = new(api.Role)

	// find a Role
	err := s.DB.Collection(RoleCollection).FindOne(ctx,
		bson.M{
			"ID": req.Id,
		},
	).Decode(res.Role)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update modifies a Role
func (s *RoleServiceServer) Update(ctx context.Context, req *api.RoleUpdateReq) (*api.RoleUpdateRes, error) {
	// prepare a Res
	res := new(api.RoleUpdateRes)

	// update a Role
	updateRes, err := s.DB.Collection(RoleCollection).UpdateOne(ctx,
		bson.M{
			"ID": req.Role.ID,
		},
		bson.M{
			"$set": bson.M{
				"Name":        req.Role.Name,
				"Description": req.Role.Description,
				"Modifiedby":  req.Role.ModifiedBy,
				"Modifiedat":  ptypes.TimestampNow(),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	res.Updated = updateRes.ModifiedCount

	return res, nil
}

// Delete removes a Role
func (s *RoleServiceServer) Delete(ctx context.Context, req *api.RoleDeleteReq) (*api.RoleDeleteRes, error) {
	// prepare a Res
	res := new(api.RoleDeleteRes)

	// delete a Role
	deleteRes, err := s.DB.Collection(RoleCollection).DeleteOne(ctx,
		bson.M{
			"ID": req.Id,
		},
	)
	if err != nil {
		return nil, err
	}

	res.Deleted = deleteRes.DeletedCount

	return res, nil
}
