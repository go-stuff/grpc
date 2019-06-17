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

// PermissionCollection is the name of the collection in the database.
const PermissionCollection string = "permissions"

// PermissionServiceServer is a service in the API for managing Permissions.
type PermissionServiceServer struct {
	DB *mongo.Database
}

// Slice returns a slice of Permissions
func (s *PermissionServiceServer) Slice(ctx context.Context, req *api.PermissionSliceReq) (*api.PermissionSliceRes, error) {
	// prepare a Res
	res := new(api.PermissionSliceRes)

	// find all Permissions
	cursor, err := s.DB.Collection(PermissionCollection).Find(ctx,
		bson.M{},
		&options.FindOptions{
			Sort: bson.M{
				"route": 1, // acending
				//"Permissionname": -1, // descending
			},
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// itterate each document returned
	for cursor.Next(ctx) {
		var Permission = new(api.Permission)
		err := cursor.Decode(&Permission)
		if err != nil {
			return nil, err
		}

		// append the current role to the slice
		res.Permissions = append(res.Permissions, Permission)
	}

	// handle any errors with the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// Create inserts a Permission
func (s *PermissionServiceServer) Create(ctx context.Context, req *api.PermissionCreateReq) (*api.PermissionCreateRes, error) {
	// prepare a Res
	res := new(api.PermissionCreateRes)

	Permission := &api.Permission{
		ID:         primitive.NewObjectID().Hex(), // ObjectID's are generated based on time
		RoleID:     req.Permission.RoleID,
		Route:      req.Permission.Route,
		CreatedBy:  req.Permission.CreatedBy,
		CreatedAt:  ptypes.TimestampNow(),
		ModifiedBy: req.Permission.ModifiedBy,
		ModifiedAt: ptypes.TimestampNow(),
	}

	// insert role into mongo
	_, err := s.DB.Collection(PermissionCollection).InsertOne(ctx, Permission)
	if err != nil {
		return nil, err
	}

	// update the Res id
	res.ID = Permission.ID

	return res, nil
}

// Read returns a single Permission
func (s *PermissionServiceServer) Read(ctx context.Context, req *api.PermissionReadReq) (*api.PermissionReadRes, error) {
	// prepare a Res
	res := new(api.PermissionReadRes)

	// initialize Permission
	res.Permission = new(api.Permission)

	// find a Permission
	err := s.DB.Collection(PermissionCollection).FindOne(ctx,
		bson.M{
			"_id": req.ID,
		},
	).Decode(res.Permission)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update modifies a Permission
func (s *PermissionServiceServer) Update(ctx context.Context, req *api.PermissionUpdateReq) (*api.PermissionUpdateRes, error) {
	// prepare a Res
	res := new(api.PermissionUpdateRes)

	// update a Permission
	updateRes, err := s.DB.Collection(PermissionCollection).UpdateOne(ctx,
		bson.M{
			"_id": req.Permission.ID,
		},
		bson.M{
			"$set": bson.M{
				"roleid":     req.Permission.RoleID,
				"route":      req.Permission.Route,
				"modifiedby": req.Permission.ModifiedBy,
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

// Delete removes a Permission
func (s *PermissionServiceServer) Delete(ctx context.Context, req *api.PermissionDeleteReq) (*api.PermissionDeleteRes, error) {
	// prepare a Res
	res := new(api.PermissionDeleteRes)

	// delete a Permission
	deleteRes, err := s.DB.Collection(PermissionCollection).DeleteOne(ctx,
		bson.M{
			"_id": req.ID,
		},
	)
	if err != nil {
		return nil, err
	}

	res.Deleted = deleteRes.DeletedCount

	return res, nil
}
