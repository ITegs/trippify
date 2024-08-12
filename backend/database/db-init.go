package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initializeDB(db *mongo.Database) {
	fmt.Println("Running DB migration")

	// USER collection
	userJsonSchema := bson.M{
		"bsonType": "object",
		"title":    "user",
		"required": []string{"_id", "username", "name"},
		"properties": bson.M{
			"_id": bson.M{
				"bsonType": "objectId",
			},
			"username": bson.M{
				"bsonType": "string",
				"pattern":  "^[a-z0-9.]+$",
			},
			"name": bson.M{
				"bsonType": "string",
			},
			"profile_pic": bson.M{
				"bsonType": "string",
			},
			"trips": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "objectId",
				},
			},
		},
	}

	userValidator := bson.M{
		"$jsonSchema": userJsonSchema,
	}

	userOpts := options.CreateCollection().SetValidator(userValidator)

	err := db.CreateCollection(context.TODO(), "users", userOpts)
	if err != nil {
		fmt.Println("Failed to create collection users: ", err)
		os.Exit(1)
		return
	}

	// Create a unique index on the username field
	usersCollection := db.Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = usersCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		fmt.Println("Failed to create unique index on username:", err)
		os.Exit(1)
		return
	}

	//	TRIPS collection
	tripJsonSchema := bson.M{
		"bsonType": "object",
		"title":    "trip",
		"required": []string{"_id", "owner", "title", "start_date"},
		"properties": bson.M{
			"_id": bson.M{
				"bsonType": "objectId",
			},
			"owner": bson.M{
				"bsonType": "string",
				"pattern":  "^[a-z0-9.]+$",
			},
			"title": bson.M{
				"bsonType": "string",
			},
			"start_date": bson.M{
				"bsonType": "string",
			},
			"end_date": bson.M{
				"bsonType": "string",
			},
			"spots": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "object",
					"properties": bson.M{
						"spotId": bson.M{
							"bsonType": "objectId",
						},
						"latitude": bson.M{
							"bsonType": "double",
						},
						"longitude": bson.M{
							"bsonType": "double",
						},
					},
				},
			},
		},
	}

	tripValidator := bson.M{
		"$jsonSchema": tripJsonSchema,
	}

	tripOpts := options.CreateCollection().SetValidator(tripValidator)

	err = db.CreateCollection(context.TODO(), "trips", tripOpts)
	if err != nil {
		fmt.Println("Failed to create collection trips: ", err)
		os.Exit(1)
		return
	}

	// SPOTS collection
	spotJsonSchema := bson.M{
		"bsonType": "object",
		"title":    "trip",
		"required": []string{"_id", "roadtripId", "title", "date_from"},
		"properties": bson.M{
			"_id": bson.M{
				"bsonType": "objectId",
			},
			"roadtripId": bson.M{
				"bsonType": "objectId",
			},
			"title": bson.M{
				"bsonType": "string",
			},
			"longitude": bson.M{
				"bsonType": "double",
			},
			"latitude": bson.M{
				"bsonType": "double",
			},
			"date_from": bson.M{
				"bsonType": "string",
			},
			"date_to": bson.M{
				"bsonType": "string",
			},
			"description": bson.M{
				"bsonType": "string",
			},
			"images": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "object",
					"properties": bson.M{
						"source": bson.M{
							"bsonType": "string",
						},
						"timestamp": bson.M{
							"bsonType": "int",
						},
					},
				}},
		},
	}

	spotValidator := bson.M{
		"$jsonSchema": spotJsonSchema,
	}

	spotOpts := options.CreateCollection().SetValidator(spotValidator)

	err = db.CreateCollection(context.TODO(), "spots", spotOpts)
	if err != nil {
		fmt.Println("Failed to create collection spots: ", err)
		os.Exit(1)
		return
	}

	fmt.Println("Migration succeeded")
}
