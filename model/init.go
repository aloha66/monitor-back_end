package model

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//对多个数据库进行初始化和连接管理

type Database struct {
	Self   *mongo.Client
	Docker *mongo.Client
}

var DB *Database

var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
		//Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	defer func() {
		if err := DB.Self.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func openDB(username, password, addr, name string) *mongo.Client {

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("db.addr")))

	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	return client
}

// used for cli
func InitSelfDB() *mongo.Client {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *mongo.Client {
	return InitSelfDB()
}

func getDBTable(databaseName string, table string) *mongo.Collection {
	return DB.Self.Database(databaseName).Collection(table)
}
