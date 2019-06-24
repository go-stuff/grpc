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

// AuditCollection is the name of the collection in the database.
const AuditCollection string = "audit"

// AuditServiceServer is a service in the API for managing Audits.
type AuditServiceServer struct {
	DB *mongo.Database
}

// List returns a slice of Audits
func (s *AuditServiceServer) List(ctx context.Context, req *api.AuditListReq) (*api.AuditListRes, error) {
	// prepare a Res
	res := new(api.AuditListRes)

	// find all Audits
	cursor, err := s.DB.Collection(AuditCollection).Find(ctx,
		bson.M{},
		// &options.FindOptions{
		// 	Sort: bson.M{
		// 		"name": 1, // acending
		// 		//"Name": -1, // descending
		// 	},
		// },
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// itterate each document returned
	for cursor.Next(ctx) {
		var Audit = new(api.Audit)
		err := cursor.Decode(&Audit)
		if err != nil {
			return nil, err
		}

		// append the current Audit to the slice
		res.Audits = append(res.Audits, Audit)
	}

	// handle any errors with the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// List100 returns the last 100 audit records
func (s *AuditServiceServer) List100(ctx context.Context, req *api.AuditList100Req) (*api.AuditList100Res, error) {
	// prepare a Res
	res := new(api.AuditList100Res)

	opts := &options.FindOptions{}
	opts.SetLimit(100)
	opts.SetSort(bson.M{"createdat.seconds": -1})

	// find all Audits
	cursor, err := s.DB.Collection(AuditCollection).Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// itterate each document returned
	for cursor.Next(ctx) {
		var Audit = new(api.Audit)
		err := cursor.Decode(&Audit)
		if err != nil {
			return nil, err
		}

		// append the current Audit to the slice
		res.Audits = append(res.Audits, Audit)
	}

	// handle any errors with the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// Create inserts a Audit
func (s *AuditServiceServer) Create(ctx context.Context, req *api.AuditCreateReq) (*api.AuditCreateRes, error) {
	// prepare a Res
	res := new(api.AuditCreateRes)

	Audit := &api.Audit{
		ID:        primitive.NewObjectID().Hex(), // ObjectID's are generated based on time
		Username:  req.Audit.Username,
		Action:    req.Audit.Action,
		Session:   req.Audit.Session,
		CreatedBy: req.Audit.CreatedBy,
		CreatedAt: ptypes.TimestampNow(),
	}

	// insert Audit into mongo
	_, err := s.DB.Collection(AuditCollection).InsertOne(ctx, Audit)
	if err != nil {
		return nil, err
	}

	// update the Res id
	res.ID = Audit.ID

	return res, nil
}

// Read returns a single Audit
func (s *AuditServiceServer) Read(ctx context.Context, req *api.AuditReadReq) (*api.AuditReadRes, error) {
	// prepare a Res
	res := new(api.AuditReadRes)

	// initialize Audit
	res.Audit = new(api.Audit)

	// find a Audit
	err := s.DB.Collection(AuditCollection).FindOne(ctx,
		bson.M{
			"_id": req.ID,
		},
	).Decode(res.Audit)
	if err != nil {
		return nil, err
	}

	return res, nil
}
