package repository

import (
	"context"
	"fmt"
	. "go-rest-mongodb/config"
	"go-rest-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UsersRepository struct {}

var config Config
var collection = new(mongo.Collection)
const UsersCollection = "Users"

func init() {
	config.Read()

	// Connect to DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection = client.Database(config.Database.DatabaseName).Collection(UsersCollection)
}

// Get all Users
func (p *UsersRepository) FindAll() ([]models.User, error) {
	var users []models.User

	findOptions := options.Find()
	findOptions.SetLimit(100)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)

	// Iterate through the cursor
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.User
		err := cur.Decode(&result)

		users = append(users, result)
	}
	return users, err
}

// Create a new User
func (p *UsersRepository) Insert(user models.User) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, &user)
	fmt.Println("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err
}

// Delete an existing User
func (p *UsersRepository) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.DeleteOne(ctx, filter)
	fmt.Println("Deleted a single document: ", result.DeletedCount)
	return err
}
