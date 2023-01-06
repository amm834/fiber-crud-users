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
			return c.Status(500).JSON(err.Error())
		}
		return c.JSON(users)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		collection := connection.Db.Collection("users")
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		result, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			return err
		}
		return c.Status(201).JSON(result)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		collection := connection.Db.Collection("users")
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return err
		}
		filter := bson.M{"_id": id}
		var user User
		err = collection.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			return err
		}
		return c.JSON(user)
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		collection := connection.Db.Collection("users")

		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return err
		}
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		filter := bson.M{"_id": id}
		update := bson.M{"$set": bson.M{"name": user.Username, "age": user.Age}}
		result, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return err
		}
		return c.Status(200).JSON(result)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		collection := connection.Db.Collection("users")
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return err
		}
		filter := bson.M{"_id": id}
		result, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return err
		}

		return c.Status(204).JSON(result)
	})

	log.Fatal(app.Listen(":8000"))
}
