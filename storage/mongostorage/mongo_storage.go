package mongostorage

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"program/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStorage(connectURI string) (*MongoStorage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURI))
	if err != nil {
		return nil, fmt.Errorf(" error while connecting to mongo: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("pinging mongo: %w", err)
	}

	db := client.Database("Devices")

	ms := &MongoStorage{
		client:     client,
		collection: db.Collection("Events"),
	}

	return ms, nil
}

func (ms *MongoStorage) Insert() {

	var c int64
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	c = 300
	rand.Seed(time.Now().Unix())

	ids := "ABCDEFGHIJ"

	for i := 0; i < 50; i++ {
		startTime := rand.Int63n(c)
		endTime := rand.Int63n(c-startTime) + startTime

		random := rand.Intn(len(ids))
		id := ids[random]

		doc := model.Event{
			ID:        primitive.NewObjectID(),
			DeviceID:  string(id),
			StartDate: int(startTime),
			EndDate:   int(endTime),
		}

		ms.collection.InsertOne(ctx, doc)
	}

}

func (ms *MongoStorage) LastStartime() ([]model.Event, []string) {
	var e []model.Event

	opts := options.Find().SetProjection(bson.D{{"endDate", 0}, {"startDate", 0}, {"_id", 0}})
	opts.SetSort(bson.D{{Key: "startDate", Value: -1}})
	opts.SetLimit(10)
	cursor, err := ms.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		panic(err)
	}
	var results []string
	if err = cursor.All(context.TODO(), &e); err != nil {
		panic(err)
	}
	for _, result := range e {
		results = append(results, result.DeviceID)
	}
	return e, results

}

func (ms *MongoStorage) EventsTime() []model.Event {
	var e []model.Event

	filter := bson.D{
		{"$or", bson.A{
			// bson.D{{"startDate", bson.D{{"$gt", 230}, {"$lt", 250}}}, {"endDate", bson.D{{"$gt", 230}, {"$lt", 250}}}},
			bson.D{{"endDate", bson.D{{"$gte", 230}, {"$lte", 250}}}},
			bson.D{{"startDate", bson.D{{"$gte", 230}, {"$lte", 250}}}},
		}},
	}

	cursor, err := ms.collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &e); err != nil {
		panic(err)
	}

	return e
}

func (ms *MongoStorage) CloseClientDB() {

	if ms.client == nil {
		return
	}

	err := ms.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
