package mongostorage

import (
	"context"
	"errors"
	"fmt"
	"program/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseStorage struct {
	client             *mongo.Client
	collectionVinyls   *mongo.Collection
	collectionUsers    *mongo.Collection
	collectionReleases *mongo.Collection
}

func NewDatabaseStorage(connectURI string) (*DatabaseStorage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURI))
	if err != nil {
		return nil, fmt.Errorf(" error while connecting to mongo: %v", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("pinging mongo: %w", err)
	}
	db := client.Database("VinilkiDatabase")
	_ = db.CreateCollection(ctx, "Vynils")
	_ = db.CreateCollection(ctx, "Users")
	_ = db.CreateCollection(ctx, "Releases")
	ds := &DatabaseStorage{
		client:             client,
		collectionVinyls:   db.Collection("Vynils"),
		collectionUsers:    db.Collection("Users"),
		collectionReleases: db.Collection("Releases"),
	}
	return ds, nil
}
func (ds *DatabaseStorage) CreateUser(user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := ds.collectionUsers.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("this username already exists")
	}

	_, insertErr := ds.collectionUsers.InsertOne(ctx, user)

	if insertErr != nil {
		return insertErr
	}
	return nil

}
func (ds *DatabaseStorage) AssignID() (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var foundUser model.User

	opts := options.FindOneOptions{}
	opts.SetSort(bson.D{{Key: "user_id", Value: -1}})

	err := ds.collectionUsers.FindOne(ctx, bson.D{}, &opts).Decode(&foundUser)
	defer cancel()

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			foundUser.User_id = 0
			return foundUser, nil
		}
		return model.User{}, fmt.Errorf("failed to execute query,error:%w", err)
	}
	return foundUser, nil
}
