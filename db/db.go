package db

import (
	"goodsman/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongodb struct {
	Records *mongo.Collection
	Goods   *mongo.Collection
}

var (
	MongoDB Mongodb
)

func Init() {
	logrus.Info("connecting databases...")
	MongoClient, err := initMongo(config.Mongo)
	if err != nil {
		logrus.Fatal("failed to connect MongoDB & ", err.Error())
	}
	MongoDB.Goods = MongoClient.Collection("goods")
	MongoDB.Records = MongoClient.Collection("records")

	logrus.Info("all databases connected")
}
