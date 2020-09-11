package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host = "localhost"
	port = "27017"

	database   = "dbName"
	collection = "myCollection"

	user = "aaa"
	pwd  = "bbb"
)

type UserWalletPreferences struct {
	ID            string `json:"_id"`
	InWalletToken string `json:"inWalletToken"`
	StoreID       int64  `json:"storeId"`
	UserID        int64  `json:"userId"`
}

func main() {
	fmt.Println("Started")

	uri := fmt.Sprintf("mongodb://%v:%v", host, port)
	credential := options.Credential{
		Username: user,
		Password: pwd,
	}
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")

	collectionDB := client.Database(database).Collection(collection)
	//Find all and filter
	cursor, err := collectionDB.Find(context.TODO(), bson.M{"storeId": 11713})
	if err != nil {
		log.Fatal(err)
	}
	var data []UserWalletPreferences
	if err = cursor.All(context.TODO(), &data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(data))
	fmt.Println(data)

	//Find One
	// var result UserWalletPreferences
	// errFindOne := collectionDB.FindOne(context.TODO(), bson.M{"inWalletToken": "ixZ7wq97C6lzcWdxArOU"}).Decode(&result)
	// if errFindOne != nil {
	// 	log.Fatal(errFindOne)
	// }
	// fmt.Printf("found document %v", result)
	// fmt.Println(result.ID, result.UserID)
}
