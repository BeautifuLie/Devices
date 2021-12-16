package mongostorage

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"program/model"

	"strconv"
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
	_ = db.CreateCollection(ctx, "Events")
	ms := &MongoStorage{
		client:     client,
		collection: db.Collection("Events"),
	}

	model := []mongo.IndexModel{
		// {
		// 	Keys: bson.D{
		// 		{Key: "endDate", Value: 1},
		// 		{Key: "startDate", Value: 1},
		// 	}},

		{
			Keys: bson.D{
				{Key: "endDate", Value: 1}},
		},
	}
	_, err = ms.collection.Indexes().CreateMany(context.TODO(), model)
	if err != nil {
		panic(err)
	}

	return ms, nil
}
func (ms *MongoStorage) Insert() {

	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	t := time.Date(2021, time.December, 6, 8, 0, 0, 0, time.Local)
	rand.Seed(time.Now().Unix())
	limit := 300000

	ids := "ABCDEFGHIJ"

	for i := 0; i < 200000; i++ {

		randomStart := rand.Intn(limit)
		startTime := t.Add(time.Millisecond * time.Duration(randomStart))

		startToString := startTime.String()
		startToInt, _ := strconv.Atoi(startToString)

		randomEndTime := rand.Intn(limit-startToInt) + randomStart
		endTime := t.Add(time.Millisecond * time.Duration(randomEndTime))

		randomID := rand.Intn(len(ids))
		id := ids[randomID]

		doc := model.Event{
			ID:        primitive.NewObjectID(),
			DeviceID:  string(id),
			StartDate: startTime,
			EndDate:   endTime,
		}

		ms.collection.InsertOne(ctx, doc)

	}

}

func (ms *MongoStorage) LastStartime() ([]model.Event, []string) {
	var e []model.Event
	//.SetProjection(bson.D{{"endDate", 0}, {"startDate", 0}, {"_id", 0}})
	opts := options.Find()
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

func (ms *MongoStorage) EventsTime(t1, t2 time.Time) []model.Event {
	var e []model.Event

	r1 := primitive.NewDateTimeFromTime(t1)
	r2 := primitive.NewDateTimeFromTime(t2)

	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"endDate", bson.D{{"$gte", r1}, {"$lte", r2}}}},
			bson.D{{"startDate", bson.D{{"$lte", r2}}}, {"endDate", bson.D{{"$exists", false}}}},
		}},
	}

	cursor, err := ms.collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &e); err != nil {
		panic(err)
	}
	fmt.Println(len(e))
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

// {"$or":[
// 	{endDate:{"$gte":ISODate("2021-12-06T06:01:00+00:00"),"$lte":ISODate("2021-12-06T06:02:00+00:00") }},
// 	{startDate:{"$gte":ISODate("2021-12-06T06:01:00+00:00"),"$lte":ISODate("2021-12-06T06:02:00+00:00") }},
// 	{"$and":[
// 		{startDate:{"$lte":ISODate("2021-12-06T06:01:00+00:00")},endDate:{"$gte":ISODate("2021-12-06T06:02:00+00:00")}}
// 	]},
// ]}

// {"$or":[
// 	{endDate:{"$gte":ISODate("2021-12-06T06:01:00+00:00"),"$lte":ISODate("2021-12-06T06:02:00+00:00") }},
// {startDate:{"$lte":ISODate("2021-12-06T06:01:00+00:00")},endDate:{"$gte":ISODate("2021-12-06T06:02:00+00:00")}},

// ]}
// {"$or":[
// 	{endDate:{"$gte":ISODate("2021-12-06T06:01:00+00:00"),"$lte":ISODate("2021-12-06T06:02:00+00:00") }},
// 	{startDate:{"$lte":ISODate("2021-12-06T06:02:00+00:00")},endDate:{"$exists":false}},
// ]}
