package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBRepo implements usecase repository
type MongoDBRepo struct {
	db       *mongo.Database
	collName string
}

// NewMongoDBRepo creates the repo
func NewMongoDBRepo(db *mongo.Database, collName string) *MongoDBRepo {
	return &MongoDBRepo{
		db:       db,
		collName: collName,
	}
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
	return users, nil
}

// Update a user
func (r *MongoDBRepo) Update(ctx context.Context, id string, user *User) error {
	userColl := r.db.Collection(r.collName)
	filter := bson.M{"id": user.ID}
	_, err := userColl.ReplaceOne(ctx, filter, user)
	if err != nil {
		return err
	}
	return nil
}

// Delete a user
func (r *MongoDBRepo) Delete(ctx context.Context, id string) error {
	// userColl := r.db.Collection(r.collName)
	return nil
}
