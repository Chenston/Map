package graph

import (
	"context"
	"backend/db"
	"backend/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Resolver 结构体
type Resolver struct{}

// 修正：把 `queryResolver` 改成 `Resolver`
func (r *Resolver) SearchStates(ctx context.Context, query string) ([]*model.State, error) {
	collection := db.MongoClient.Database("us_states").Collection("states")

	// MongoDB 查询
	filter := bson.M{"name": bson.M{"$regex": query, "$options": "i"}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 解析查询结果
	var states []*model.State
	for cursor.Next(ctx) {
		var state model.State
		if err := cursor.Decode(&state); err != nil {
			return nil, err
		}
		states = append(states, &state)
	}

	return states, nil
}
