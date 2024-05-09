package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Welcome to GoToDo")

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error while loading .env files - ", err)
		}
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// TO OPTIMISE THE CODE -
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDb Atlas")
	// -------------------------------------x------------------------------------- //
	// Connection to Db :

	collection = client.Database("goToDo").Collection("todos")

	app := fiber.New()

	// DO NOT NEED IN DEVELOPMENT MODE -
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:5173",
	// 	AllowHeaders: "Origin,Content-Type,Accept",
	// }))

	app.Get("/api/v1/todos", getToDos)
	app.Post("/api/v1/todos", createToDos)
	app.Patch("/api/v1/todos/:id", updateToDos)
	app.Delete("/api/v1/todos/:id", deleteToDos)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
	}

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getToDos(c *fiber.Ctx) error {
	var todos []Todo

	// 1) Cursor = When we execute a query in MongoDb, it returns a cursor. It is a pointer to the result set which is used to iterate over the result and display it to user.
	// 2) bson.M{} = No filters. Display every todo from the db.
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	// To optimize the code -
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var singleTodo Todo
		if err := cursor.Decode(&singleTodo); err != nil {
			return err
		}
		todos = append(todos, singleTodo)
	}
	return c.JSON(todos)
}

func createToDos(c *fiber.Ctx) error {
	newTodo := new(Todo)

	if err := c.BodyParser(newTodo); err != nil {
		return err
	}

	if newTodo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"Error": "Todo body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), newTodo)
	if err != nil {
		return err
	}

	newTodo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(newTodo)
}

func updateToDos(c *fiber.Ctx) error {
	// This is of String type
	id := c.Params("id")

	// ---> Convert the string to ObjectID -
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "Invalid Todo ID"})
	}

	// collection.UpdateOne(context.Background(), filter (Which one we will update), update(Inside that obj, which field we are going to update))
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"Success": true})
}

func deleteToDos(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"Success": true})
}
