package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"k8s.io/klog/v2"
	"log"
	"os"
)

func connectToMongo() *mongo.Client {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")

	url := fmt.Sprintf("mongodb://%s:%s@localhost:27017/admin?directConnection=true&serverSelectionTimeoutMS=2000&authSource=admin", username, password)
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
	return client
}

func main() {
	client := connectToMongo()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	//db := client.Database("test")
	//test(db, "kubedb")

	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	dbs, err := client.ListDatabases(context.TODO(), bson.D{})
	if err != nil {
		klog.Errorf("err = %s \n", err.Error())
		return
	}
	for _, d := range dbs.Databases {
		if d.Name == "config" || d.Name == "admin" || d.Name == "local" {
			continue
		}
		db := client.Database(d.Name)
		names, err := db.ListCollectionNames(context.Background(), bson.D{})
		if err != nil {
			return
		}
		for _, name := range names {
			calcSize(db, name)
		}
	}
}

const dir = "log"

func writeFile(fileName string, data []byte) {
	fileName = dir + "/" + fileName + ".json"
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func calcSize(db *mongo.Database, coll string) {
	cmd := bson.D{{"collStats", coll}}
	var result bson.M
	err := db.RunCommand(context.TODO(), cmd).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	//var b []byte
	//buf := bytes.NewBuffer(b)
	//enc := json.NewEncoder(buf)
	//err = enc.Encode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//writeFile(fmt.Sprintf("%s.%s", db.Name(), coll), buf.Bytes())

	indentedData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	writeFile(fmt.Sprintf("%s.%s", db.Name(), coll), indentedData)
}
