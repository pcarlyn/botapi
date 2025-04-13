package database

import (
	"context"
	"start/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Remove(id primitive.ObjectID) models.SchedulerAnswer {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:password@mongo:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	c := client.Database("user").Collection("scheduler")
	var deleted models.SchedulerAnswer
	err = c.FindOneAndDelete(context.TODO(), bson.M{"_id": id}).Decode(&deleted)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.SchedulerAnswer{}
		}
		panic(err)
	}
	return deleted
}
