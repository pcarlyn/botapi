package database

import (
	"context"
	"start/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Save(msg models.SchedulerMsg) primitive.ObjectID {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:password@mongo:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	now := time.Now()

	scheduler := models.SchedulerAnswer{
		Content:   msg.Content,
		SendAt:    msg.SendAt,
		Channel:   msg.Channel,
		Recipient: msg.Recipient,
		Tag:       msg.Tag,
		Status:    "pending",
		CreatedAt: now,
		UpdatedAt: now,
	}
	c := client.Database("user").Collection("scheduler")
	insertResult, err := c.InsertOne(context.TODO(), scheduler)
	if err != nil {
		panic(err)
	}
	return insertResult.InsertedID.(primitive.ObjectID)
}
