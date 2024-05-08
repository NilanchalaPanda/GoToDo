package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Creating Structure of the ToDo Application -
type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	// fmt.Println("Hello, World! Hey Fiber!")
	app := fiber.New()

	todos := []Todo{}

	// ____ TEST ROUTES ____ //
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{"message": "Hello, World! Nilanchala here"})
	// })
	// app.Get("/test", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{"message": "Hello, this is a test message!"})
	// })

	// ____ ACTUAL ROUTES ____ //
	// -----------> POST ROUTE : C
	app.Post("/api/v1/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"Error": "Todo body is required"})
		}

		// Go ONLY have for loops -
		for _, existingTodo := range todos {
			if existingTodo.Body == todo.Body {
				return c.Status(400).JSON(fiber.Map{"Error": "Todo already exists"})
			}
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(200).JSON(todo)
	})

	// -----------> UPDATE ROUTE : R
	app.Get("/api/v1/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// -----------> UPDATE ROUTE : U
	app.Patch("/api/v1/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for _, existingTodo := range todos {
			// L.H.S WILL BE OF TYPE "STRING", convert it to INTEGER
			if fmt.Sprint(existingTodo.ID) == id {
				existingTodo.Completed = !existingTodo.Completed
				return c.Status(200).JSON(existingTodo)
			}
		}
		return c.Status(400).JSON(fiber.Map{"Error": "Todo not found"})
	})

	// -----------> DELETE ROUTE : D
	// app.Delete("/api/v1/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i, todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos = append(todos[:i], todos[i+1:]...)
	// 			return c.Status(200).JSON(fiber.Map{"Success": "Todo deleted succesfully"})
	// 		}
	// 	}
	// 	return c.Status(400).JSON(fiber.Map{"Error": "Todo not found"})
	// })

	app.Delete("/api/v1/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		fmt.Println(todos)

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				// Todo found and deleted, returning success response
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		// Todo not found, returning error response
		return c.Status(400).JSON(fiber.Map{"Error": "Todo not found"})
	})

	log.Fatal(app.Listen(":4000"))
}
