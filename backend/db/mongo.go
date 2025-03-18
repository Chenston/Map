package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB 连接客户端
var MongoClient *mongo.Client

// 连接 MongoDB
func ConnectDB() {
	uri := "mongodb://admin:admin123@localhost:27017/"
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB 连接失败: %v", err)
	}

	// 检查连接状态
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("无法连接 MongoDB: %v", err)
	}

	MongoClient = client
	fmt.Println("MongoDB 连接成功！")
}
