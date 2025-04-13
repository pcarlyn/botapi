package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProcessById(id primitive.ObjectID) bson.M {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:password@mongo:27017"))
	if err != nil {
		panic(err)
	}
	c := client.Database("user").Collection("process")
	filter := bson.M{"_id": id}
	cursor, err := c.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())
	var process bson.M
	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&process)
		if err != nil {
			panic(err)
		}
	}
	return process
}

func GetAllProcesses() ([]bson.M, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:password@mongo:27017"))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	c := client.Database("user").Collection("process")
	cursor, err := c.Find(context.TODO(), bson.M{}) // Фильтр пустой - выбираются все записи
	if err != nil {
		return nil, err
	}
	var process []bson.M
	if err = cursor.All(context.TODO(), &process); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return process, nil
}
