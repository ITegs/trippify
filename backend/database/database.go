package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username   string               `json:"username,omitempty" bson:"username"`
	Password   string               `json:"password,omitempty" bson:"password"`
	Name       string               `json:"name,omitempty" bson:"name"`
	ProfilePic string               `json:"profile_pic" bson:"profile_pic"`
	Trips      []primitive.ObjectID `json:"trips" bson:"trips"`
}

type Trip struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Owner         *User              `json:"owner,omitempty" bson:"owner"`
	OwnerUsername string             `json:"owner_username,omitempty" bson:"owner_username"`
	Title         string             `json:"title,omitempty" bson:"title"`
	StartDate     string             `json:"start_date" bson:"start_date"`
	EndDate       string             `json:"end_date" bson:"end_date"`
	Spots         []TripSpot         `json:"spots" bson:"spots"`
}

type TripSpot struct {
	SpotId    primitive.ObjectID `json:"spotId,omitempty" bson:"spotId"`
	Latitude  float64            `json:"latitude,omitempty" bson:"latitude"`
	Longitude float64            `json:"longitude,omitempty" bson:"longitude"`
}

type Spot struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	RoadtripId  primitive.ObjectID `json:"roadtripId,omitempty" bson:"roadtripId"`
	Title       string             `json:"title,omitempty" bson:"title"`
	Latitude    float64            `json:"latitude,omitempty" bson:"latitude"`
	Longitude   float64            `json:"longitude,omitempty" bson:"longitude"`
	Images      []image            `json:"images,omitempty" bson:"images"`
	Description string             `json:"description" bson:"description"`
	DateFrom    string             `json:"date_from" bson:"date_from"`
	DateTo      string             `json:"date_to" bson:"date_to"`
}

type image struct {
	Source    string `json:"source,omitempty" bson:"source"`
	Timestamp int32  `json:"timestamp" bson:"timestamp"`
}

var mCl *mongo.Client

func init() {
	fmt.Println("Initializing the DB")

	dbUri, ok := os.LookupEnv("DATABASE_URI")
	if !ok {
		dbUri = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(dbUri)
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
	spots *mongo.Collection
}

type DB interface {
	GetUserByUsername(username string) (*User, error)
	NewUser(user *User) (*User, error)
	GetFirstTrip() (*Trip, error)
	GetTrip(tripId string) (*Trip, error)
	NewTrip(trip *Trip) error
	AddSpot(spot Spot) error
	GetSpot(spotId string) (*Spot, error)
}

func NewDB(users *mongo.Collection, trips *mongo.Collection, spots *mongo.Collection) DB {
	db := &db{
		users: users,
		trips: trips,
		spots: spots,
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

func RunMigrations(databaseName string) {
	database := mCl.Database(databaseName)
	initializeDB(database)
}

func (db *db) GetUserByUsername(username string) (*User, error) {
	result := db.users.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}})

	var user User

	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *db) NewUser(user *User) (*User, error) {
	user.Trips = []primitive.ObjectID{}
	result, err := db.users.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserted a new user: ", result.InsertedID)
	fullUser, err := db.GetUserByUsername(user.Username)
	if err != nil {
		// TODO: didnt work as planned
		return nil, err
	}

	return fullUser, nil
}

func (db *db) GetFirstTrip() (*Trip, error) {
	result := db.trips.FindOne(context.TODO(), bson.D{})

	var trip Trip
	err := result.Decode(&trip)
	if err != nil {
		return nil, err
	}

	return &trip, nil
}

func (db *db) GetTrip(tripId string) (*Trip, error) {
	objectId, err := primitive.ObjectIDFromHex(tripId)
	if err != nil {
		return nil, err
	}

	result := db.trips.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}})

	var trip Trip

	err = result.Decode(&trip)
	if err != nil {
		return nil, err
	}

	return &trip, nil
}

func (db *db) NewTrip(trip *Trip) error {

	tripResult, err := db.trips.InsertOne(context.TODO(), trip)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a new trip: ", tripResult.InsertedID)

	filter := bson.D{{Key: "username", Value: trip.OwnerUsername}}
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "trips", Value: tripResult.InsertedID},
		}},
	}

	_, err = db.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (db *db) AddSpot(spot Spot) error {
	spot.ID = primitive.NewObjectID()
	spotResult, err := db.spots.InsertOne(context.TODO(), spot)
	if err != nil {
		return err
	}
	fmt.Println("Inserted a new spot: ", spotResult.InsertedID)

	// set the last spots date_to to now
	filter := bson.D{{Key: "date_to", Value: ""}}
	update := bson.D{{"$set", bson.D{{"date_to", spot.DateFrom}}}}
	dateToResult, err := db.spots.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	fmt.Println(dateToResult.ModifiedCount)

	tripSpot := TripSpot{
		SpotId:    spotResult.InsertedID.(primitive.ObjectID),
		Longitude: spot.Longitude,
		Latitude:  spot.Latitude,
	}

	filter = bson.D{{Key: "_id", Value: spot.RoadtripId}}
	update = bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "spots", Value: tripSpot},
		}},
	}

	tripsResult, err := db.trips.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("Added the spot to the trip: ", tripsResult.UpsertedID)

	return nil
}

func (db *db) GetSpot(spotId string) (*Spot, error) {
	objectId, err := primitive.ObjectIDFromHex(spotId)
	if err != nil {
		return nil, err
	}

	result := db.spots.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}})

	var spot Spot
	err = result.Decode(&spot)
	if err != nil {
		return nil, err
	}

	return &spot, nil
}
