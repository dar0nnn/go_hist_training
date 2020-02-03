package mongo_hist

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create_client() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

func Create_conn(client *mongo.Client) error {
	err := client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func Check_conn(client *mongo.Client) bool {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func Create_client_with_conn() (*mongo.Client, error) {
	client, err := Create_client()
	if err != nil {
		return nil, err
	}
	err = Create_conn(client)
	return client, err
}
