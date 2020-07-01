package mongo

import (
	"context"
	"time"

	"github.com/marceloaguero/serverless-api/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository implements users repository
type mongoRepo struct {
	db       *mongo.Database
	collName string
}

// NewMongoRepo creates the repo
func NewMongoRepo(dbURI, dbName, collName string) (users.Repository, error) {
	db, err := mongoConnect(dbURI, dbName)
	if err != nil {
		return nil, err
	}

	return &mongoRepo{
		db:       db,
		collName: collName,
	}, nil
}

func mongoConnect(dbURI, dbName string) (*mongo.Database, error) {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = dbClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = dbClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return dbClient.Database(dbName), nil
}

// Create a user
func (r *mongoRepo) Create(ctx context.Context, user *users.User) error {
	userColl := r.db.Collection(r.collName)
	_, err := userColl.InsertOne(ctx, user)
	return err
}

// Get a user
func (r *mongoRepo) Get(ctx context.Context, id string) (*users.User, error) {
	result := users.User{}
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}

	err := userColl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAll users
func (r *mongoRepo) GetAll(ctx context.Context) ([]*users.User, error) {
	result := make([]*users.User, 0)
	user := &users.User{}
	userColl := r.db.Collection(r.collName)

	cur, err := userColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

// Update a user
func (r *mongoRepo) Update(ctx context.Context, id string, user *users.UpdateUser) error {
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}
	_, err := userColl.ReplaceOne(ctx, filter, user)
	if err != nil {
		return err
	}
	return nil
}

// Delete a user
func (r *mongoRepo) Delete(ctx context.Context, id string) error {
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}
	_, err := userColl.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
