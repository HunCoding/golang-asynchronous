package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  int64              `bson:"age"`
}

type ChanUser struct {
	Err  error
	User *User
}

func InsertUserAsyncMongoDB(
	ctx context.Context,
	user User,
	chanUser chan ChanUser,
) {
	collection := Database.Collection("user_performance_test")
	insertedResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		chanUser <- ChanUser{
			Err:  err,
			User: nil,
		}
		return
	}

	user.ID = insertedResult.InsertedID.(primitive.ObjectID)
	chanUser <- ChanUser{
		Err:  nil,
		User: &user,
	}
	close(chanUser)
}

func InsertUserSyncMongoDB(
	ctx context.Context,
	user User,
) (*User, error) {
	collection := Database.Collection("user_performance_test")
	insertedResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &user, nil
}
