package database

import (
	"context"
	"fmt"
	"start/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Put(id primitive.ObjectID, updatedData models.SchedulerMsg) models.SchedulerAnswer {
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

	existing := Get(id)

	createdAtRaw, ok := existing["createdat"]
	if !ok {
		panic("createdat not found in document")
	}

	var createdAt time.Time
	switch v := createdAtRaw.(type) {
	case primitive.DateTime:
		createdAt = v.Time()
	case time.Time:
		createdAt = v
	default:
		panic(fmt.Sprintf("unknown type for createdat: %T", v))
	}

	putData := models.SchedulerAnswer{
		Content:   updatedData.Content,
		SendAt:    updatedData.SendAt,
		Channel:   updatedData.Channel,
		Recipient: updatedData.Recipient,
		Tag:       updatedData.Tag,
		Status:    "pending",
		CreatedAt: createdAt,
		UpdatedAt: now,
	}
	_, err = c.ReplaceOne(context.TODO(), bson.M{"_id": id}, putData)
	if err != nil {
		panic(err)
	}
	return putData
}
