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
func (s *RouteServiceServer) Slice(ctx context.Context, req *api.RouteSliceReq) (*api.RouteSliceRes, error) {
	// prepare a Res
	res := new(api.RouteSliceRes)

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

// Create inserts a Route
func (s *RouteServiceServer) Create(ctx context.Context, req *api.RouteCreateReq) (*api.RouteCreateRes, error) {
	// prepare a Res
	res := new(api.RouteCreateRes)

	Route := &api.Route{
		ID:         primitive.NewObjectID().Hex(), // ObjectID's are generated based on time
		Name:       req.Route.Name,
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

// Read returns a single Route
func (s *RouteServiceServer) ReadByName(ctx context.Context, req *api.RouteReadByNameReq) (*api.RouteReadByNameRes, error) {
	// prepare a Res
	res := new(api.RouteReadByNameRes)

	// initialize Route
	res.Route = new(api.Route)

	// find a Route
	err := s.DB.Collection(RouteCollection).FindOne(ctx,
		bson.M{
			"name": req.Name,
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
				"name":       req.Route.Name,
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
