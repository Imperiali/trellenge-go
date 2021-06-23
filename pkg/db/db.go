package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Database interface {
	GetById(collection string, id primitive.ObjectID, contract interface{}) (interface{}, error)
	Save(collection string, content interface{}) error
	//Delete(string) (interface{}, error)
}

type database struct {
	client *mongo.Client
}

func (d database) GetById(collection string, id primitive.ObjectID, contract interface{}) (interface{}, error) {
	c := d.client.Database("trellenge").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//filter := bson.D{primitive.E{Key: id}}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	if err := c.FindOne(ctx, filter).Decode(&contract); err != nil {
		fmt.Println("Não achou...")
		log.Fatal(err)
		return nil, err
	}
	return contract, nil
}

func (d database) Save(collection string, content interface{}) error {
	c := d.client.Database("trellenge").Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if _, err := c.InsertOne(ctx, content); err != nil {
		fmt.Println("Não achou...")
		log.Fatal(err)
		return err
	}

	return nil
}

//func (d database) Delete (string) (interface{}, error) {}

func New() Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo.hud:27017/trellenge"))
	if err != nil {
		fmt.Println("Error on create new client")
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting")
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error on list databases")
		log.Fatal(err)
	}

	fmt.Println("Funcionando, esses são os databases")
	fmt.Println(databases)

	return &database{
		client,
	}
}
