package users

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBRepo implements usecase repository
type MongoDBRepo struct {
	db       *mongo.Database
	collName string
}

// NewDBRepo creates the repo
func NewDBRepo(dbURI, dbName, collName string) *MongoDBRepo {
	db, err := dbConnect(dbURI, dbName)
	if err != nil {
		os.Exit(1)
	}
	return &MongoDBRepo{
		db:       db,
		collName: collName,
	}
}

func dbConnect(dbURI, dbName string) (*mongo.Database, error) {
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
func (r *MongoDBRepo) Create(ctx context.Context, user *User) error {
	userColl := r.db.Collection(r.collName)
	_, err := userColl.InsertOne(ctx, user)
	return err
}

// Get a user
func (r *MongoDBRepo) Get(ctx context.Context, id string) (*User, error) {
	user := &User{}
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}

	err := userColl.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAll users
func (r *MongoDBRepo) GetAll(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	user := &User{}
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
func (r *MongoDBRepo) Update(ctx context.Context, id string, user *UpdateUser) error {
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}
	_, err := userColl.ReplaceOne(ctx, filter, user)
	if err != nil {
		return err
	}
	return nil
}

// Delete a user
func (r *MongoDBRepo) Delete(ctx context.Context, id string) error {
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": id}
	_, err := userColl.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
