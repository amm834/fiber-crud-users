package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var connection MongoDBInstance

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	databaseName := os.Getenv("MONGODB_DB_NAME")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	db := client.Database(databaseName)

	connection = MongoDBInstance{
		Client: client,
		Db:     db,
	}
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"name"`
	Age      int                `bson:"age"`
}

func main() {

	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		collection := connection.Db.Collection("users")

		filter := bson.M{}

		var users []User
		cursor, err := collection.Find(context.Background(), filter)
		if err != nil {
			return err
		}
		if err = cursor.All(context.Background(), &users); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(users)
	})

	log.Fatal(app.Listen(":8000"))
}
