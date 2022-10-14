package orm

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectOnce sync.Once
var client *mongo.Client
var err error

func connect(uri string) *mongo.Client {
	connectOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(1000)
		client, err = mongo.Connect(context.TODO(), clientOptions)
	})

	if err != nil {
		log.Fatal(err)
	}

	return client
}

type MongoQuery struct {
	MongoURI   string
	Database   string
	Collection string
	Key        string
	Value      interface{}
	Page       int64
	PerPage    int64
}

/*
   FindOne returns a single document from the collection.

   It takes a struct as a parameter, and returns a bson.M

   The Key / Value pair is not used with this type of Query.
*/

func (m *MongoQuery) FindAll() ([]bson.M, error) {
	client := connect(m.MongoURI)
	ctx := context.Background()
	var results []bson.M
	collection := client.Database(m.Database).Collection(m.Collection)

	filter := bson.D{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

/*
FindOne returns a single document from the collection.

It takes a struct as a parameter, and returns a bson.M
*/
func (m *MongoQuery) FindOne() (bson.M, error) {
	client := connect(m.MongoURI)
	ctx := context.Background()
	var result bson.M
	collection := client.Database(m.Database).Collection(m.Collection)
	filter := bson.D{{Key: m.Key, Value: m.Value}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

/*
Find returns multple documents from the collection.

It takes a struct as a parameter, and returns a []bson.M
*/
func (m *MongoQuery) Find() ([]bson.M, error) {
	client := connect(m.MongoURI)
	ctx := context.Background()
	var results []bson.M
	collection := client.Database(m.Database).Collection(m.Collection)

	filter := bson.D{{Key: m.Key, Value: m.Value}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

type MongoInsert struct {
	MongoURI   string
	Database   string
	Collection string
	Value      bson.D
}

/*
InsertOne inserts one document to the collection.

It takes a struct as a parameter, and returns an error if
something happened.
*/
func (m *MongoInsert) InsertOne() error {
	client := connect(m.MongoURI)
	ctx := context.Background()
	collection := client.Database(m.Database).Collection(m.Collection)
	_, err := collection.InsertOne(ctx, m.Value)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type MongoUpdate struct {
	MongoURI    string
	Database    string
	Collection  string
	FilterKey   string
	FilterValue interface{}
	Value       bson.D
}

/*
UpdateOne updates one document to the collection.

It takes a struct as a parameter, and returns an error if
something happened.
*/
func (m *MongoUpdate) UpdateOne() error {
	client := connect(m.MongoURI)
	ctx := context.Background()
	collection := client.Database(m.Database).Collection(m.Collection)
	filter := bson.D{{Key: m.FilterKey, Value: m.FilterValue}}
	update := bson.D{{Key: "$set", Value: m.Value}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

/*
   Raw Mongo Query
*/

type MongoRawQuery struct {
	MongoURI   string
	Database   string
	Collection string
	Query      bson.D
}

/*
FindRaw returns multple documents from the collection.

It takes a struct as a parameter, and returns a []bson.M

This function exposes filtering for advanced queries.

Please refer to the mongo documentation for further information.
*/
func (m *MongoRawQuery) FindRaw() ([]bson.M, error) {
	client := connect(m.MongoURI)
	ctx := context.Background()
	collection := client.Database(m.Database).Collection(m.Collection)
	var results []bson.M
	filter := m.Query
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

type MongoRawUpdate struct {
	MongoURI   string
	Database   string
	Collection string
	Filter     bson.D
	Update     bson.D
}

/*
UpdateRaw returns an error if something happened.

This function exposes filtering and updating custom queries for advanced use.

Please refer to the mongo documentation for further information.
*/
func (m *MongoRawUpdate) UpdateRaw() error {
	client := connect(m.MongoURI)
	ctx := context.Background()
	collection := client.Database(m.Database).Collection(m.Collection)
	_, err := collection.UpdateOne(ctx, m.Filter, m.Update)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type MongoDelete struct {
	MongoURI    string
	Database    string
	Collection  string
	FilterKey   string
	FilterValue interface{}
}

func (m *MongoDelete) Delete() error {
	client := connect(m.MongoURI)
	ctx := context.Background()
	collection := client.Database(m.Database).Collection(m.Collection)
	filter := bson.D{{Key: m.FilterKey, Value: m.FilterValue}}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
