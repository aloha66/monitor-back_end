package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type BaseModel struct {
	CreateTime time.Time `json:"createTime,omitempty" bson:"create_time"`
	UpdateTime time.Time `json:"updateTime,omitempty" bson:"update_time"`
}

func GetIncreaseId(dbName string, tableName string) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	count, err := getDBTable(dbName, tableName).CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0
	}
	return count
}
