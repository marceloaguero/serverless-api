package mongo

import (
	"context"
	"os"
	"time"

	"github.com/marceloaguero/serverless-api/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository implements users repository
type repository struct {
	db       *mongo.Database
	collName string
}

// NewRepository creates the repo
func NewRepository(dbURI, dbName, collName string) *users.Repository {
	db, err := mongoConnect(dbURI, dbName)
	if err != nil {
		os.Exit(1)
	}
	return &repository{
		db:       db,
		collName: collName,
	}
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
func (r *repository) Create(ctx context.Context, user *users.User) error {
	userColl := r.db.Collection(r.collName)
	_, err := userColl.InsertOne(ctx, user)
	return err
}

// Get a user
func (r *repository) Get(ctx context.Context, id string) (*users.User, error) {
	user := &users.User{}
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}

	err := userColl.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAll users
func (r *repository) GetAll(ctx context.Context) ([]*users.User, error) {
	users := make([]*users.User, 0)
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
		users = append(users, user)
	}
	return users, nil
}

// Update a user
func (r *repository) Update(ctx context.Context, id string, user *users.UpdateUser) error {
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}
	_, err := userColl.ReplaceOne(ctx, filter, user)
	if err != nil {
		return err
	}
	return nil
}

// Delete a user
func (r *repository) Delete(ctx context.Context, id string) error {
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}
	_, err := userColl.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
