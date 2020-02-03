package mongo_hist

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelingOutputValue struct {
	project_id                   string
	code                         int
	name                         string
	profitableGeologicalReserves float32
}

func Get_data_for_hist(client *mongo.Client, project_id string) []*ModelingOutputValue {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := client.Database("vega").Collection("business.ModelingOutputValues")
	filter := bson.D{{"project_id", project_id}}
	var find_one_res ModelingOutputValue
	err := collection.FindOne(ctx, filter).Decode(&find_one_res)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// fmt.Println(find_one_res)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Found a document: %+v\n", cur)
	var results []*ModelingOutputValue

	for cur.Next(ctx) {
		var elem ModelingOutputValue
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(elem.code)
		results = append(results, &elem)
	}
	return results
}
