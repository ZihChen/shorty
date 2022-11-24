package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var connectPool *mongo.Database

type dbIns struct {
	ctx context.Context
}

type Interface interface {
	GetConnection() (*mongo.Database, error)
	GetContext() *context.Context
}

func NewInstance() Interface {
	return &dbIns{
		ctx: context.Background(),
	}
}

func (db *dbIns) GetContext() *context.Context {
	return &db.ctx
}

func (db *dbIns) GetConnection() (*mongo.Database, error) {
	ctxTimeout, cancel := context.WithTimeout(db.ctx, 15*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017").
		SetMaxPoolSize(2000).                // 連線池最大連接數
		SetMaxConnIdleTime(10 * time.Second) // 閒置連接最大保持毫秒數 10s

	client, err := mongo.Connect(ctxTimeout, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(db.ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	connectPool = client.Database("test")

	fmt.Println("Connected to MongoDB!")

	return connectPool, nil
}
