package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get(id primitive.ObjectID) bson.M {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:password@mongo:27017"))
	if err != nil {
		panic(err)
	}
	c := client.Database("user").Collection("scheduler")
	filter := bson.M{"_id": id}
	cursor, err := c.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())
	var item bson.M
	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&item)
		if err != nil {
			panic(err)
		}
	}
	return item
}

func GetAll() ([]bson.M, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:password@mongo:27017"))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	c := client.Database("user").Collection("scheduler")
	cursor, err := c.Find(context.TODO(), bson.M{}) // Фильтр пустой - выбираются все записи
	if err != nil {
		return nil, err
	}
	var schedulers []bson.M
	if err = cursor.All(context.TODO(), &schedulers); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return schedulers, nil
}
