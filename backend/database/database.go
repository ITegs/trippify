package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name string `json:"name,omitempty" bson:"name"`
}

type Trip struct {
	Owner     primitive.ObjectID `json:"owner,omitempty" bson:"owner"`
	Title     string             `json:"title,omitempty" bson:"title"`
	StartDate primitive.DateTime `json:"start_date" bson:"start_date"`
	Spots     []Spot             `json:"spots" bson:"spots"`
}

type image struct {
	Source      string `json:"source,omitempty" bson:"source"`
	Description string `json:"description,omitempty" bson:"description"`
}

type Spot struct {
	Title     string  `json:"title,omitempty" bson:"title"`
	Longitude float64 `json:"longitude,omitempty" bson:"longitude"`
	Latitude  float64 `json:"latitude,omitempty" bson:"latitude"`
	Images    []image `json:"images,omitempty" bson:"images"`
}

var mCl *mongo.Client

func init() {
	fmt.Println("Initializing the DB")
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017/")
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed to init the db")
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Failed to init the db")
	}

	fmt.Println("Connection to MongoDB established!")

	mCl = mongoClient
}

type db struct {
	users *mongo.Collection
	trips *mongo.Collection
}

type DB interface {
	NewUser(user *User) error
	NewTrip(trip *Trip) error
	AddSpotToTrip(tripId string, spot Spot) error
}

func NewDB(users *mongo.Collection, trips *mongo.Collection) DB {
	db := &db{
		users: users,
		trips: trips,
	}

	return db
}

func CollectionFactory(databaseName string, collectionName string) (*mongo.Collection, error) {

	database := mCl.Database(databaseName)

	collectionNames, err := database.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	if len(collectionNames) == 0 {
		initializeDB(database)
	}

	return database.Collection(collectionName), nil
}

func (db *db) NewUser(user *User) error {
	result, err := db.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a new user: ", result.InsertedID)
	return nil
}

func (db *db) NewTrip(trip *Trip) error {
	result, err := db.trips.InsertOne(context.TODO(), trip)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a new trip: ", result.InsertedID)
	return nil
}

func (db *db) AddSpotToTrip(tripId string, spot Spot) error {

	objectId, err := primitive.ObjectIDFromHex(tripId)
	if err != nil {
		return fmt.Errorf("invalid trip ID format: %v", err)
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{
		{"$push", bson.D{
			{"spots", spot},
		}},
	}

	result, err := db.trips.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no document matching the filter found")
	}

	return nil
}
