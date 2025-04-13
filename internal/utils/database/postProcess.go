package database

import (
	"context"
	"start/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostProcess(proc models.Process, pid uint32) primitive.ObjectID {
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
	result := models.ProcessAnswer{
		Command: models.Process{
			Path: proc.Path,
			Args: proc.Args,
		},
		Pid:   pid,
		RunAt: now,
	}
	c := client.Database("user").Collection("process")
	insertResult, err := c.InsertOne(context.TODO(), result)
	if err != nil {
		panic(err)
	}

	return insertResult.InsertedID.(primitive.ObjectID)

}
