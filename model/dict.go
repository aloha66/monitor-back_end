package model

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"time"
)

type DictModel struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Value      string    `json:"value" binding:"required"`
	CreateTime time.Time `json:"createTime,omitempty" bson:"create_time"`
	UpdateTime time.Time `json:"updateTime,omitempty" bson:"update_time"`
	//BaseModel
}

func (dict *DictModel) CreateDict() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := getDBTable(viper.GetString("db.name"), "dict")
	insertResult, err := collection.InsertOne(ctx, &dict)
	if err != nil {
		log.Println("err", err)
		return err
	}
	log.Println("insertResult", insertResult)
	return nil
}
