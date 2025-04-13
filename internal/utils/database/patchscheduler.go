package database

import (
	"context"
	"start/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Patch(id primitive.ObjectID, updateMsg models.UpdateMsg) models.SchedulerAnswer {
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
	now := time.Now()

	update := bson.M{
		"$set": bson.M{
			"status":    updateMsg.Msg,
			"updatedat": now,
		},
	}
	var result models.SchedulerAnswer

	err = c.FindOneAndUpdate(
		context.TODO(),
		bson.M{"_id": id},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&result)

	if err != nil {
		panic(err)
	}
	return result
}
