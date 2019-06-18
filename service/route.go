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

// RouteCollection is the name of the collection in the database.
const RouteCollection string = "routes"

// RouteServiceServer is a service in the API for managing Routes.
type RouteServiceServer struct {
	DB *mongo.Database
}

// Slice returns a slice of Routes
func (s *RouteServiceServer) List(ctx context.Context, req *api.RouteListReq) (*api.RouteListRes, error) {
	// prepare a Res
	res := new(api.RouteListRes)

	// find all Routes
	cursor, err := s.DB.Collection(RouteCollection).Find(ctx,
		bson.M{},
		&options.FindOptions{
			Sort: bson.M{
				"name": 1, // acending
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
		var Route = new(api.Route)
		err := cursor.Decode(&Route)
		if err != nil {
			return nil, err
		}

		// append the current Route to the slice
		res.Routes = append(res.Routes, Route)
	}

	// handle any errors with the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// ListByRoleID returns a slice of Routes by RoleID
func (s *RouteServiceServer) ListByRoleID(ctx context.Context, req *api.RouteListByRoleIDReq) (*api.RouteListByRoleIDRes, error) {
	// prepare a Res
	res := new(api.RouteListByRoleIDRes)

	// find all Routes
	cursor, err := s.DB.Collection(RouteCollection).Find(ctx,
		bson.M{
			"roleid": req.RoleID,
		},
		&options.FindOptions{
			Sort: bson.M{
				"path": 1, // acending
				//"path": -1, // descending
			},
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// itterate each document returned
	for cursor.Next(ctx) {
		var Route = new(api.Route)
		err := cursor.Decode(&Route)
		if err != nil {
			return nil, err
		}

		// append the current Route to the slice
		res.Routes = append(res.Routes, Route)
	}

	// handle any errors with the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// Create inserts a Route
func (s *RouteServiceServer) Create(ctx context.Context, req *api.RouteCreateReq) (*api.RouteCreateRes, error) {
	// prepare a Res
	res := new(api.RouteCreateRes)

	Route := &api.Route{
		ID:         primitive.NewObjectID().Hex(), // ObjectID's are generated based on time
		Path:       req.Route.Path,
		CreatedBy:  req.Route.CreatedBy,
		CreatedAt:  ptypes.TimestampNow(),
		ModifiedBy: req.Route.ModifiedBy,
		ModifiedAt: ptypes.TimestampNow(),
	}

	// insert Route into mongo
	_, err := s.DB.Collection(RouteCollection).InsertOne(ctx, Route)
	if err != nil {
		return nil, err
	}

	// update the Res id
	res.ID = Route.ID

	return res, nil
}

// Read returns a single Route
func (s *RouteServiceServer) Read(ctx context.Context, req *api.RouteReadReq) (*api.RouteReadRes, error) {
	// prepare a Res
	res := new(api.RouteReadRes)

	// initialize Route
	res.Route = new(api.Route)

	// find a Route
	err := s.DB.Collection(RouteCollection).FindOne(ctx,
		bson.M{
			"_id": req.ID,
		},
	).Decode(res.Route)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ReadByRoleIDAndPath returns a single Route
func (s *RouteServiceServer) ReadByRoleIDAndPath(ctx context.Context, req *api.RouteReadByRoleIDAndPathReq) (*api.RouteReadByRoleIDAndPathRes, error) {
	// prepare a Res
	res := new(api.RouteReadByRoleIDAndPathRes)

	// initialize Route
	res.Route = new(api.Route)

	// find a Route
	err := s.DB.Collection(RouteCollection).FindOne(ctx,
		bson.M{
			"roleid": req.Route.RoleID,
			"path":   req.Route.Path,
		},
	).Decode(res.Route)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update modifies a Route
func (s *RouteServiceServer) Update(ctx context.Context, req *api.RouteUpdateReq) (*api.RouteUpdateRes, error) {
	// prepare a Res
	res := new(api.RouteUpdateRes)

	// update a Route
	updateRes, err := s.DB.Collection(RouteCollection).UpdateOne(ctx,
		bson.M{
			"_id": req.Route.ID,
		},
		bson.M{
			"$set": bson.M{
				"path":       req.Route.Path,
				"modifiedby": req.Route.ModifiedBy,
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

// UpdateByRoleIDAndPath modifies a Route
func (s *RouteServiceServer) UpdateByRoleIDAndPath(ctx context.Context, req *api.RouteUpdateByRoleIDAndPathReq) (*api.RouteUpdateByRoleIDAndPathRes, error) {
	// prepare a Res
	res := new(api.RouteUpdateByRoleIDAndPathRes)

	// update a Route
	updateRes, err := s.DB.Collection(RouteCollection).UpdateOne(ctx,
		bson.M{
			"roleid": req.Route.RoleID,
			"path":   req.Route.Path,
		},
		bson.M{
			"$set": bson.M{
				// "roleid":     req.Route.RoleID,
				// "path":       req.Route.Path,
				"permission": req.Route.Permission,
				"modifiedby": req.Route.ModifiedBy,
				"modifiedat": ptypes.TimestampNow(),
			},
		},
		//options.Update().SetUpsert(true),
	)
	if err != nil {
		return nil, err
	}

	res.Updated = updateRes.ModifiedCount

	if res.Updated == 0 {

		Route := &api.Route{
			ID:         primitive.NewObjectID().Hex(),
			RoleID:     req.Route.RoleID,
			Path:       req.Route.Path,
			Permission: req.Route.Permission,
			CreatedBy:  "System",
			CreatedAt:  ptypes.TimestampNow(),
			ModifiedBy: "System",
			ModifiedAt: ptypes.TimestampNow(),
		}

		// insert Route into mongo
		_, err := s.DB.Collection(RouteCollection).InsertOne(ctx, Route)
		if err != nil {
			return nil, err
		}

		// update the Res id
		res.Updated = 1
		//res.ID = Route.ID
	}

	return res, nil
}

// Delete removes a Route
func (s *RouteServiceServer) Delete(ctx context.Context, req *api.RouteDeleteReq) (*api.RouteDeleteRes, error) {
	// prepare a Res
	res := new(api.RouteDeleteRes)

	// delete a Route
	deleteRes, err := s.DB.Collection(RouteCollection).DeleteOne(ctx,
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
